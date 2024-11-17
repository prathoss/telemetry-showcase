package dao

import (
	"time"

	"github.com/google/uuid"
)

type Ride struct {
	ID        uuid.UUID
	StartDate time.Time
	EndDate   time.Time
	UserID    uuid.UUID
	BikeID    uuid.UUID
}
