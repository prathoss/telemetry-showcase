package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/prathoss/telemetry_showcase/bikes/config"
	"github.com/prathoss/telemetry_showcase/bikes/dao/repository"
	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"github.com/prathoss/telemetry_showcase/shared"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisotel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

var _ bikes.BikesServer = (*Service)(nil)

type Service struct {
	bikes.UnimplementedBikesServer
	cfg  config.Config
	repo repository.Repository
}

func New(cfg config.Config) (*Service, error) {
	// connect to DB
	db, err := gorm.Open(postgres.Open(cfg.DbConnStr), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "showcase."}})
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	// instrument tracing for ORM
	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, fmt.Errorf("could not attach tracing to gorm: %w", err)
	}

	// add redis instrumentation
	rdb, err := rueidisotel.NewClient(rueidis.ClientOption{
		InitAddress: cfg.RedisAddrs,
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to redis: %w", err)
	}
	r := repository.NewRepository(db, rdb)

	return &Service{
		cfg:  cfg,
		repo: r,
	}, nil
}

func (s *Service) Run() error {
	grpcServer := grpc.NewServer(
		// instrument grpc server
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(shared.GrpcLoggingInterceptor),
	)
	bikes.RegisterBikesServer(grpcServer, s)

	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// wait for context to finish and stop grpc server
	go func(ctx context.Context) {
		<-ctx.Done()
		slog.Info("Stopping gRPC server...")
		grpcServer.GracefulStop()
		slog.Info("gRPC server stopped gracefully")
	}(ctx)

	slog.Info("starting gRPC server", slog.String("address", s.cfg.Address))
	if err := grpcServer.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}
