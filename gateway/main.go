package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/prathoss/telemetry_showcase/gateway/config"
	"github.com/prathoss/telemetry_showcase/gateway/server"
	"github.com/prathoss/telemetry_showcase/gateway/server/resolver"
	"github.com/prathoss/telemetry_showcase/shared"
	"github.com/ravilushqa/otelgqlgen"
)

func main() {
	shared.SetupLogging()
	tracingShutdown, metricsShutdown, err := shared.SetupTelemetry()
	if err != nil {
		slog.Error("couldn't setup telemetry", shared.Err(err))
	} else {
		defer tracingShutdown()
		defer metricsShutdown()
	}

	cfg := config.NewConfig()

	srv := handler.NewDefaultServer(server.NewExecutableSchema(server.Config{Resolvers: &resolver.Resolver{}}))
	srv.Use(otelgqlgen.Middleware())

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	s := http.Server{
		Addr:              cfg.Address,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
		ErrorLog:          slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func(ctx context.Context) {
		<-ctx.Done()
		shutdownCtx, shutdownCFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCFunc()
		if err := s.Shutdown(shutdownCtx); err != nil {
			slog.Error("couldn't shutdown gracefully", shared.Err(err))
		}
	}(ctx)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("couldn't start http server", shared.Err(err))
	}
}
