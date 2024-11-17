package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetRide(ctx context.Context, r *rides.GetRideRequest) (*rides.RideReply, error) {
	ride, err := s.repo.GetRide(ctx, uuid.MustParse(r.GetRideId()))
	if err != nil {
		return nil, err
	}
	return &rides.RideReply{
		Id:         ride.ID.String(),
		UserId:     ride.UserID.String(),
		BikeId:     ride.BikeID.String(),
		StartTime:  timestamppb.New(ride.StartDate),
		EndTime:    timestamppb.New(ride.EndDate),
		InvoiceUrl: nil,
	}, nil
}
