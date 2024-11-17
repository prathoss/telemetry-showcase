package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) EndRide(ctx context.Context, request *rides.EndRideRequest) (*rides.RideReply, error) {
	r, err := s.repo.EndRide(ctx, uuid.MustParse(request.GetRideId()))
	if err != nil {
		return nil, err
	}

	if err := s.producer.SendRideEnd(ctx, r.ID); err != nil {
		return nil, err
	}

	if _, err := s.bikesClient.SetBikeAvailable(ctx, &bikes.SetBikeAvailableRequest{Id: r.BikeID.String()}); err != nil {
		return nil, err
	}

	return &rides.RideReply{
		Id:         r.ID.String(),
		UserId:     r.UserID.String(),
		BikeId:     r.BikeID.String(),
		StartTime:  timestamppb.New(r.StartDate),
		EndTime:    timestamppb.New(r.EndDate),
		InvoiceUrl: nil,
	}, nil
}
