package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/bikes"
)

func (s *Service) SetBikeReserved(ctx context.Context, r *bikes.SetBikeReservedRequest) (*bikes.SetBikeReservedReply, error) {
	bike, err := s.repo.GetBikeByID(ctx, uuid.MustParse(r.Id))
	if err != nil {
		return nil, err
	}
	if err := s.repo.SetBikeReserved(ctx, bike); err != nil {
		return nil, err
	}
	return &bikes.SetBikeReservedReply{}, nil
}
