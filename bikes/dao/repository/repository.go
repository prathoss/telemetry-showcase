package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/bikes/dao"
	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

const geosearchKeyBikes = "bikes"

type Repository interface {
	GetBikeByID(ctx context.Context, id uuid.UUID) (dao.Bike, error)
	ListBikes(ctx context.Context, lat float64, lon float64) ([]dao.Bike, error)
	SetBikeReserved(ctx context.Context, bike dao.Bike) error
	SetBikeAvailable(ctx context.Context, bike dao.Bike) error
}

func NewRepository(db *gorm.DB, rdb rueidis.Client) *GormRepository {
	return &GormRepository{db: db, rdb: rdb}
}

var _ Repository = (*GormRepository)(nil)

type GormRepository struct {
	db  *gorm.DB
	rdb rueidis.Client
}

func (g *GormRepository) ListBikes(ctx context.Context, lat float64, lon float64) ([]dao.Bike, error) {
	bikeIDs, err := g.getBikesNearby(ctx, lat, lon)
	if err != nil {
		return nil, err
	}
	if len(bikeIDs) == 0 {
		bikes, err := g.getAvailableBikes(ctx)
		if err != nil {
			return nil, err
		}
		storeCmd := g.rdb.B().Geoadd().Key(geosearchKeyBikes).LongitudeLatitudeMember()
		for _, b := range bikes {
			storeCmd = storeCmd.LongitudeLatitudeMember(b.Lon, b.Lat, b.ID.String())
		}
		if err := g.rdb.Do(ctx, storeCmd.Build()).Error(); err != nil {
			return nil, err
		}
		bikeIDs, err = g.getBikesNearby(ctx, lat, lon)
		if err != nil {
			return nil, err
		}
		if len(bikeIDs) == 0 {
			return []dao.Bike{}, nil
		}
	}

	bikes := make([]dao.Bike, 0, len(bikeIDs))
	if err := g.db.WithContext(ctx).Find(&bikes, bikeIDs).Error; err != nil {
		return nil, err
	}
	return bikes, nil
}

func (g *GormRepository) SetBikeAvailable(ctx context.Context, bike dao.Bike) error {
	command := g.rdb.B().Geoadd().
		Key(geosearchKeyBikes).
		LongitudeLatitudeMember().
		LongitudeLatitudeMember(bike.Lon, bike.Lat, bike.ID.String()).
		Build()
	if err := g.rdb.Do(ctx, command).Error(); err != nil {
		return err
	}
	if err := g.db.WithContext(ctx).
		Model(&dao.Bike{}).
		Where("id = ?", bike.ID).
		Update("available", true).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) SetBikeReserved(ctx context.Context, bike dao.Bike) error {
	command := g.rdb.B().Zrem().Key(geosearchKeyBikes).Member(bike.ID.String()).Build()
	if err := g.rdb.Do(ctx, command).Error(); err != nil {
		return err
	}
	if err := g.db.WithContext(ctx).
		Model(&dao.Bike{}).
		Where("id = ?", bike.ID).
		Update("available", false).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormRepository) GetBikeByID(ctx context.Context, id uuid.UUID) (dao.Bike, error) {
	var bike dao.Bike
	err := g.db.WithContext(ctx).First(&bike, "id = ?", id).Error
	return bike, err
}

func (g *GormRepository) getAvailableBikes(ctx context.Context) ([]dao.Bike, error) {
	bikes := make([]dao.Bike, 0, 10)
	if err := g.db.WithContext(ctx).Where("available = ?", true).Find(&bikes).Error; err != nil {
		return nil, err
	}
	return bikes, nil
}

func (g *GormRepository) getBikesNearby(ctx context.Context, lat, lon float64) ([]uuid.UUID, error) {
	command := g.rdb.B().Geosearch().Key(geosearchKeyBikes).Fromlonlat(lon, lat).Bybox(12).Height(12).Km().Build()
	result := g.rdb.Do(ctx, command)
	if err := result.Error(); err != nil {
		return nil, err
	}
	locations, err := result.AsGeosearch()
	if err != nil {
		return nil, err
	}
	bikeIDs := make([]uuid.UUID, 0, len(locations))
	for _, location := range locations {
		bikeID, err := uuid.Parse(location.Name)
		if err != nil {
			return nil, err
		}
		bikeIDs = append(bikeIDs, bikeID)
	}
	return bikeIDs, nil
}
