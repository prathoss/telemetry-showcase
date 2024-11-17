package shared

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

// SetupLogging configures defaults slog logger for uniform logging
func SetupLogging() {
	logger := slog.New(&SlogHandlerWrapper{
		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		// setup logging middleware, when writing logs, check data from context to enrich logs
		Extractors: []Extractor{
			otelLogExtractor,
		},
	})
	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	if serviceName != "" {
		logger = logger.With(slog.String("service", serviceName))
	}
	slog.SetDefault(logger)
}

type Extractor func(ctx context.Context) []slog.Attr

type SlogHandlerWrapper struct {
	slog.Handler
	Extractors []Extractor
}

func (s *SlogHandlerWrapper) Handle(ctx context.Context, record slog.Record) error {
	for _, extractor := range s.Extractors {
		record.AddAttrs(extractor(ctx)...)
	}

	return s.Handler.Handle(ctx, record)
}

func (s *SlogHandlerWrapper) WithAttrs(attrs []slog.Attr) slog.Handler {
	w := s.Handler.WithAttrs(attrs)
	return &SlogHandlerWrapper{
		Handler:    w,
		Extractors: s.Extractors,
	}
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
