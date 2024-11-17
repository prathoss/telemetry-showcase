package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/prathoss/telemetry_showcase/invoices/config"
	"github.com/prathoss/telemetry_showcase/invoices/messages/kafka"
	"github.com/prathoss/telemetry_showcase/shared"
)

func main() {
	shared.SetupLogging()
	traceShutdown, metricsShutdown, err := shared.SetupTelemetry()
	if err != nil {
		slog.Error("could not setup telemetry", shared.Err(err))
	} else {
		defer traceShutdown()
		defer metricsShutdown()
	}
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("could not create config", shared.Err(err))
		os.Exit(1)
		return
	}

	ctx, cFunc := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cFunc()

	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		slog.Error("could not create consumer", shared.Err(err))
		os.Exit(1)
		return
	}
	consumer.ConsumeRideEnd(ctx)
}
