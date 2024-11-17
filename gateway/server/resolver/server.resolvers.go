package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"time"

	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"github.com/prathoss/telemetry_showcase/proto/users"
	"github.com/prathoss/telemetry_showcase/shared"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/gateway/server"
	"github.com/prathoss/telemetry_showcase/gateway/server/model"
	"google.golang.org/genproto/googleapis/type/latlng"
)

// StartRide is the resolver for the startRide field.
func (r *mutationResolver) StartRide(ctx context.Context, bikeID uuid.UUID) (*model.RideResponse, error) {
	userID := shared.UserIdFromContext(ctx)
	resp, err := r.ridesClient.StartRide(ctx, &rides.StartRideRequest{BikeId: bikeID.String(), UserId: userID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to start ride: %w", err)
	}
	var endTime *string = nil
	if resp.GetEndTime() != nil {
		et := resp.GetEndTime().AsTime().Format(time.RFC3339)
		endTime = &et
	}
	return &model.RideResponse{
		ID:        uuid.MustParse(resp.GetId()),
		UserID:    uuid.MustParse(resp.GetUserId()),
		BikeID:    uuid.MustParse(resp.GetBikeId()),
		StartTime: resp.StartTime.AsTime().Format(time.RFC3339),
		EndTime:   endTime,
	}, nil
}

// EndRide is the resolver for the endRide field.
func (r *mutationResolver) EndRide(ctx context.Context, rideID uuid.UUID) (*model.RideResponse, error) {
	resp, err := r.ridesClient.EndRide(ctx, &rides.EndRideRequest{RideId: rideID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to end ride: %w", err)
	}
	var endTime *string = nil
	if resp.GetEndTime() != nil {
		et := resp.GetEndTime().AsTime().Format(time.RFC3339)
		endTime = &et
	}
	return &model.RideResponse{
		ID:        uuid.MustParse(resp.GetId()),
		UserID:    uuid.MustParse(resp.GetUserId()),
		BikeID:    uuid.MustParse(resp.GetBikeId()),
		StartTime: resp.StartTime.AsTime().Format(time.RFC3339),
		EndTime:   endTime,
	}, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*model.UserResponse, error) {
	userID := shared.UserIdFromContext(ctx)

	u, err := r.usersClient.GetUserById(ctx, &users.GetUserByIdRequest{Id: userID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &model.UserResponse{
		ID:        uuid.MustParse(u.GetId()),
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Email:     u.GetEmail(),
	}, nil
}

// GetBike is the resolver for the getBike field.
func (r *queryResolver) GetBike(ctx context.Context, id uuid.UUID) (*model.BikeResponse, error) {
	bike, err := r.bikesClient.GetBikeById(ctx, &bikes.GetBikeByIdRequest{Id: id.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to get bike: %w", err)
	}

	return &model.BikeResponse{
		ID:       uuid.MustParse(bike.Id),
		Lat:      bike.GetLocation().GetLatitude(),
		Lon:      bike.GetLocation().GetLongitude(),
		ImageURL: bike.GetImageUrl(),
	}, nil
}

// ListBikes is the resolver for the listBikes field.
func (r *queryResolver) ListBikes(ctx context.Context, req model.ListBikesRequest) ([]*model.BikeResponse, error) {
	resp, err := r.bikesClient.ListBikes(ctx, &latlng.LatLng{
		Latitude:  req.Lat,
		Longitude: req.Lon,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get bikes: %w", err)
	}
	b := make([]*model.BikeResponse, 0, len(resp.Bikes))
	for _, bike := range resp.Bikes {
		b = append(b, &model.BikeResponse{
			ID:       uuid.MustParse(bike.Id),
			Lat:      bike.Location.Latitude,
			Lon:      bike.GetLocation().GetLongitude(),
			ImageURL: bike.GetImageUrl(),
		})
	}
	return b, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
