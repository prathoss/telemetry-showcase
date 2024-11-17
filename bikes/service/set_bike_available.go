package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/bikes"
)

// bike, err := s.repo.GetBikeByID(ctx, uuid.MustParse(r.Id))
//	if err != nil {
//		return nil, err
//	}
//	if err := s.repo.SetBikeReserved(ctx, bike); err != nil {
//		return nil, err
//	}
//	return &bikes.SetBikeReservedReply{}, nil

func (s *Service) SetBikeAvailable(ctx context.Context, r *bikes.SetBikeAvailableRequest) (*bikes.SetBikeAvailableReply, error) {
	bike, err := s.repo.GetBikeByID(ctx, uuid.MustParse(r.Id))
	if err != nil {
		return nil, err
	}
	if err := s.repo.SetBikeAvailable(ctx, bike); err != nil {
		return nil, err
	}
	return &bikes.SetBikeAvailableReply{}, nil
}
