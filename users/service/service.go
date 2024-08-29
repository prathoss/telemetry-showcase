package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/users/config"
	"github.com/prathoss/telemetry_showcase/users/dao"
	"github.com/prathoss/telemetry_showcase/users/dao/repository"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisaside"
	"github.com/redis/rueidis/rueidisotel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
	"proto/users"
	"shared"
)

var _ users.UserServer = (*Service)(nil)

type Service struct {
	users.UnimplementedUserServer
	cfg        config.Config
	cache      rueidisaside.CacheAsideClient
	repository repository.Repository
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
	r := repository.NewGormRepository(db)

	// create redis cache client
	cache, err := rueidisaside.NewClient(
		rueidisaside.ClientOption{
			ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
				// add redis instrumentation
				c, err := rueidisotel.NewClient(option)
				return c, err
			},
			ClientOption: rueidis.ClientOption{
				InitAddress: cfg.RedisAddrs,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect to redis: %w", err)
	}

	return &Service{
		cfg:        cfg,
		repository: r,
		cache:      cache,
	}, nil
}

func (s *Service) GetUserById(ctx context.Context, request *users.GetUserByIdRequest) (*users.UserReply, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	val, err := s.cache.Get(
		ctx,
		5*time.Minute,
		fmt.Sprintf("user:%s", id.String()),
		func(ctx context.Context, key string) (string, error) {
			u, err := s.repository.GetUserByID(ctx, id)
			if err != nil {
				return "", err
			}
			val, err := json.Marshal(u)
			if err != nil {
				return "", err
			}
			return string(val), nil
		},
	)
	var errNotFound *shared.ErrNotFound
	if errors.As(err, &errNotFound) {
		return nil, status.Error(codes.NotFound, errNotFound.Error())
	}
	if err != nil {
		return nil, err
	}

	var user dao.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	return &users.UserReply{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, request *users.GetUserByEmailRequest) (*users.UserReply, error) {
	val, err := s.cache.Get(
		ctx,
		5*time.Minute,
		fmt.Sprintf("user:%s", request.Email),
		func(ctx context.Context, key string) (string, error) {
			u, err := s.repository.GetUserByEmail(ctx, request.Email)
			if err != nil {
				return "", err
			}
			val, err := json.Marshal(u)
			if err != nil {
				return "", err
			}
			return string(val), nil
		},
	)
	var errNotFound *shared.ErrNotFound
	if errors.As(err, &errNotFound) {
		return nil, status.Error(codes.NotFound, errNotFound.Error())
	}
	if err != nil {
		return nil, err
	}

	var user dao.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	return &users.UserReply{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (s *Service) Run() error {
	grpcServer := grpc.NewServer(
		// instrument grpc server
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
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
