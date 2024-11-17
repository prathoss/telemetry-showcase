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

	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"github.com/prathoss/telemetry_showcase/rides/config"
	"github.com/prathoss/telemetry_showcase/rides/dao/repository"
	"github.com/prathoss/telemetry_showcase/rides/messages/kafka"
	"github.com/prathoss/telemetry_showcase/shared"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

var _ rides.RidesServer = (*Service)(nil)

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

	r := repository.NewGormRepository(db)

	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not create kafka producer: %w", err)
	}

	bikesConn, err := grpc.NewClient(
		cfg.BikesAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	bikesClient := bikes.NewBikesClient(bikesConn)

	return &Service{
		cfg:         cfg,
		repo:        r,
		producer:    producer,
		bikesClient: bikesClient,
	}, nil
}

type Service struct {
	rides.UnimplementedRidesServer
	cfg         config.Config
	repo        repository.Repository
	bikesClient bikes.BikesClient
	// usersClient users.UserClient
	producer *kafka.Producer
}

func (s *Service) mustEmbedUnimplementedRidesServer() {
	// TODO implement me
	panic("implement me")
}

func (s *Service) Run() error {
	grpcServer := grpc.NewServer(
		// instrument grpc server
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(shared.GrpcLoggingInterceptor),
	)
	rides.RegisterRidesServer(grpcServer, s)

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
