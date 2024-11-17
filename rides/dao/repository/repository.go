package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/rides/dao"
	"github.com/prathoss/telemetry_showcase/shared"
	"gorm.io/gorm"
)

type Repository interface {
	StartRide(ctx context.Context, bikeID, userID uuid.UUID) (dao.Ride, error)
	EndRide(ctx context.Context, rideID uuid.UUID) (dao.Ride, error)
	GetRide(ctx context.Context, rideID uuid.UUID) (dao.Ride, error)
	Fail(ctx context.Context, rideID uuid.UUID) (dao.Ride, error)
}

var _ Repository = (*GormRepository)(nil)

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

type GormRepository struct {
	db *gorm.DB
}

func (g *GormRepository) Fail(ctx context.Context, rideID uuid.UUID) (dao.Ride, error) {
	type Fail struct {
		ID uuid.UUID
	}

	if err := g.db.WithContext(ctx).Create(&Fail{rideID}).Error; err != nil {
		slog.ErrorContext(ctx, "failed to insert", shared.Err(err))
		return dao.Ride{}, err
	}
	return dao.Ride{}, nil
}

func (g *GormRepository) StartRide(ctx context.Context, bikeID, userID uuid.UUID) (dao.Ride, error) {
	r := dao.Ride{
		ID:        uuid.New(),
		BikeID:    bikeID,
		UserID:    userID,
		StartDate: time.Now().UTC(),
	}
	if err := g.db.WithContext(ctx).Create(&r).Error; err != nil {
		return dao.Ride{}, err
	}
	return r, nil
}

func (g *GormRepository) EndRide(ctx context.Context, rideID uuid.UUID) (dao.Ride, error) {
	r, err := g.GetRide(ctx, rideID)
	if err != nil {
		return dao.Ride{}, err
	}
	r.EndDate = time.Now().UTC()
	if err := g.db.WithContext(ctx).Save(&r).Error; err != nil {
		return dao.Ride{}, err
	}
	r, err = g.GetRide(ctx, rideID)
	if err != nil {
		return dao.Ride{}, err
	}
	return r, nil
}

func (g *GormRepository) GetRide(ctx context.Context, rideID uuid.UUID) (dao.Ride, error) {
	var r dao.Ride
	if err := g.db.WithContext(ctx).Where("id = ?", rideID).First(&r).Error; err != nil {
		return dao.Ride{}, err
	}
	return r, nil
}
