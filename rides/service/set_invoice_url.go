package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/rides"
)

func (s *Service) SetInvoiceUlr(ctx context.Context, request *rides.SetInvoiceUrlRequest) (*rides.RideReply, error) {
	rd, err := s.repo.Fail(ctx, uuid.MustParse(request.RideId))
	if err != nil {
		return nil, err
	}
	return &rides.RideReply{Id: rd.ID.String()}, nil
}
