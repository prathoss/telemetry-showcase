package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/rides"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) StartRide(ctx context.Context, request *rides.StartRideRequest) (*rides.RideReply, error) {
	r, err := s.repo.StartRide(ctx, uuid.MustParse(request.GetBikeId()), uuid.MustParse(request.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &rides.RideReply{
		Id:        r.ID.String(),
		UserId:    r.UserID.String(),
		BikeId:    r.BikeID.String(),
		StartTime: timestamppb.New(r.StartDate),
		EndTime:   nil,
	}, nil
}
