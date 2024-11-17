package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/prathoss/telemetry_showcase/gateway/config"
	"github.com/prathoss/telemetry_showcase/gateway/middleware"
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

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("couldn't setup config", shared.Err(err))
		os.Exit(1)
		return
	}

	res, err := resolver.New(cfg)
	if err != nil {
		slog.Error("couldn't setup resolver", shared.Err(err))
		os.Exit(1)
		return
	}

	srv := handler.NewDefaultServer(server.NewExecutableSchema(server.Config{Resolvers: res}))
	srv.Use(otelgqlgen.Middleware(otelgqlgen.WithCreateSpanFromFields(func(ctx *graphql.FieldContext) bool {
		return false
	})))
	srv.Use(&shared.GraphqlLogger{})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	s := http.Server{
		Addr:              cfg.Address,
		Handler:           middleware.UserMiddleware(mux),
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
