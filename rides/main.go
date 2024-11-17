package main

import (
	"log/slog"
	"os"

	"github.com/prathoss/telemetry_showcase/rides/config"
	"github.com/prathoss/telemetry_showcase/rides/service"
	"github.com/prathoss/telemetry_showcase/shared"
)

func main() {
	shared.SetupLogging()
	stopTracer, stopMeter, err := shared.SetupTelemetry()
	if err != nil {
		slog.Error("could not setup telemetry ", shared.Err(err))
	} else {
		defer stopTracer()
		defer stopMeter()
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("could not load config ", shared.Err(err))
		os.Exit(1)
		return
	}

	s, err := service.New(cfg)
	if err != nil {
		slog.Error("could not create service ", shared.Err(err))
		os.Exit(1)
		return
	}
	if err := s.Run(); err != nil {
		slog.Error("could not run service ", shared.Err(err))
		os.Exit(1)
		return
	}
}
