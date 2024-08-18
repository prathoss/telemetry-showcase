package shared

import (
	"context"
	"log/slog"
	"os"

	"github.com/prathoss/logenricher"
	"go.opentelemetry.io/otel/trace"
)

// SetupLogging configures defaults slog logger for uniform logging
func SetupLogging() {
	logger := slog.New(&logenricher.SlogHandlerWrapper{
		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		// setup logging middleware, when writing logs, check data from context to enrich logs
		Extractors: []logenricher.Extractor{
			otelLogExtractor,
		},
	})
	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName != "" {
		logger = logger.With(slog.String("service", serviceName))
	}
	slog.SetDefault(logger)
}

// otelLogExtractor extracts trace id and span id from context, if context does not have span returns empty slice
func otelLogExtractor(ctx context.Context) []slog.Attr {
	if span := trace.SpanFromContext(ctx); span != nil {
		return []slog.Attr{
			slog.String("trace_id", span.SpanContext().TraceID().String()),
			slog.String("span_id", span.SpanContext().SpanID().String()),
		}
	}
	return []slog.Attr{}
}

func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}
