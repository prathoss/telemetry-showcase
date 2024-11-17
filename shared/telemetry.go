package shared

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/host"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

// SetupTelemetry initializes tracing and metrics. Returns trace and meter shutdown
// functions which should be deferred.
//
// Automatically sets instrumentation:
//   - runtime
//   - host
func SetupTelemetry() (func(), func(), error) {
	ctx := context.Background()
	slog := slog.Default().With(slog.String("component", "telemetry"))

	rs, err := resource.New(
		ctx,
		resource.WithFromEnv(),
		resource.WithContainer(),
		resource.WithTelemetrySDK(),
	)

	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		slog.Warn("wrong configuration of telemetry resources", Err(err))
	} else if err != nil {
		return nil, nil, fmt.Errorf("could not initialize telemetry resource: %w", err)
	}

	// prepare tracing with grpc otlp exporter
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create trace exporter: %w", err)
	}
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(rs),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// prepare metrics with grpc otlp exporter
	metricExporter, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create metric exporter: %w", err)
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(rs),
	)
	otel.SetMeterProvider(meterProvider)

	// setup default metrics
	if err := runtime.Start(); err != nil {
		slog.Warn("could not start runtime metric exporter", Err(err))
	}
	if err := host.Start(); err != nil {
		slog.Warn("could not start host metric exporter", Err(err))
	}

	// prepare shutdown functions
	traceShutdownFunc := func() {
		ctx, cFunc := context.WithTimeout(context.Background(), 10*time.Second)
		if err := tracerProvider.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "could not shutdown tracer", Err(err))
		}
		cFunc()
	}
	meterShutdownFunc := func() {
		ctx, cFunc := context.WithTimeout(context.Background(), 10*time.Second)
		if err := meterProvider.Shutdown(ctx); err != nil {
			slog.ErrorContext(ctx, "could not shutdown meter", Err(err))
		}
		cFunc()
	}
	return traceShutdownFunc, meterShutdownFunc, nil
}
