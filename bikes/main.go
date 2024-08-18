package main

import (
	"log/slog"

	"github.com/prathoss/telemetry_showcase/bikes/config"
	"github.com/prathoss/telemetry_showcase/bikes/service"
	"github.com/prathoss/telemetry_showcase/shared"
)

func main() {
	shared.SetupLogging()
	traceShutdown, metricsShutdown, err := shared.SetupTelemetry()
	if err != nil {
		slog.Error("couldn't setup telemetry", shared.Err(err))
	} else {
		defer traceShutdown()
		defer metricsShutdown()
	}

	cfg := config.NewConfig()
	s := service.New(cfg)
	if err := s.Run(); err != nil {
		slog.Error("could not run service", shared.Err(err))
	}
}
