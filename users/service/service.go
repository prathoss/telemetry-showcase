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

	"github.com/prathoss/telemetry_showcase/proto/users"
	"github.com/prathoss/telemetry_showcase/users/config"
	"google.golang.org/grpc"
)

var _ users.UserServer = (*Service)(nil)

type Service struct {
	users.UnimplementedUserServer
	cfg config.Config
}

func New(cfg config.Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) GetUserById(ctx context.Context, request *users.GetUserByIdRequest) (*users.UserReply, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) GetUserByEmail(ctx context.Context, request *users.GetUserByEmailRequest) (*users.UserReply, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Service) Run() error {
	grpcServer := grpc.NewServer()
	users.RegisterUserServer(grpcServer, s)

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
