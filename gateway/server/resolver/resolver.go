package resolver

import (
	"github.com/prathoss/telemetry_showcase/gateway/config"
	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"github.com/prathoss/telemetry_showcase/proto/users"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func New(cfg config.Config) (*Resolver, error) {
	usersConn, err := grpc.NewClient(
		cfg.UsersAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	usersClient := users.NewUserClient(usersConn)

	bikesConn, err := grpc.NewClient(
		cfg.BikesAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	bikesClient := bikes.NewBikesClient(bikesConn)

	ridesConn, err := grpc.NewClient(
		cfg.RidesAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	ridesClient := rides.NewRidesClient(ridesConn)

	return &Resolver{
		usersClient: usersClient,
		bikesClient: bikesClient,
		ridesClient: ridesClient,
	}, nil
}

type Resolver struct {
	usersClient users.UserClient
	bikesClient bikes.BikesClient
	ridesClient rides.RidesClient
}
