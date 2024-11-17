package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/shared"
	"github.com/prathoss/telemetry_showcase/users/dao"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(ctx context.Context, ID uuid.UUID) (dao.User, error)
	GetUserByEmail(ctx context.Context, email string) (dao.User, error)
}

var _ Repository = (*GormRepository)(nil)

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

type GormRepository struct {
	db *gorm.DB
}

func (g *GormRepository) GetUserByID(ctx context.Context, ID uuid.UUID) (dao.User, error) {
	var user dao.User
	result := g.db.WithContext(ctx).Where("id = ?", ID).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, shared.NewErrNotFound("could not find user by id")
	}
	if result.Error != nil {
		return dao.User{}, result.Error
	}
	return user, nil
}

func (g *GormRepository) GetUserByEmail(ctx context.Context, email string) (dao.User, error) {
	var user dao.User
	result := g.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, shared.NewErrNotFound("could not find user by email")
	}
	if result.Error != nil {
		return dao.User{}, result.Error
	}
	return user, nil
}
