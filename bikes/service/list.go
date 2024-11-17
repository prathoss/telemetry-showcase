package service

import (
	"context"

	"github.com/prathoss/telemetry_showcase/proto/bikes"
	"google.golang.org/genproto/googleapis/type/latlng"
)

func (s *Service) ListBikes(ctx context.Context, location *latlng.LatLng) (*bikes.ListBikesReply, error) {
	bks, err := s.repo.ListBikes(ctx, location.GetLatitude(), location.GetLongitude())
	if err != nil {
		return nil, err
	}
	b := make([]*bikes.BikeReply, 0, len(bks))
	for _, bk := range bks {
		b = append(b, &bikes.BikeReply{
			Id: bk.ID.String(),
			Location: &latlng.LatLng{
				Latitude:  bk.Lat,
				Longitude: bk.Lon,
			},
			ImageUrl: bk.ImageUrl,
		})
	}
	return &bikes.ListBikesReply{Bikes: b}, nil
}
