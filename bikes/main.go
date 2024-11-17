package main

import (
	"log/slog"
	"os"

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

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("couldn't setup config", shared.Err(err))
		os.Exit(1)
		return
	}

	s, err := service.New(cfg)
	if err != nil {
		slog.Error("couldn't setup service", shared.Err(err))
		os.Exit(1)
		return
	}
	if err := s.Run(); err != nil {
		slog.Error("could not run service", shared.Err(err))
	}
}
