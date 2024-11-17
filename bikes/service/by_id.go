package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"github.com/prathoss/telemetry_showcase/shared"
	"google.golang.org/genproto/googleapis/type/latlng"
)

func (s *Service) GetBikeById(ctx context.Context, request *bikes.GetBikeByIdRequest) (*bikes.BikeReply, error) {
	bike, err := s.repo.GetBikeByID(ctx, uuid.MustParse(request.GetId()))
	if err != nil {
		slog.ErrorContext(ctx, "could not load bike", shared.Err(err))
	}

	return &bikes.BikeReply{
		Id: bike.ID.String(),
		Location: &latlng.LatLng{
			Latitude:  bike.Lat,
			Longitude: bike.Lon,
		},
		ImageUrl: bike.ImageUrl,
	}, nil
}
