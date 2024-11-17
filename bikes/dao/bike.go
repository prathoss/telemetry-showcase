package dao

import "github.com/google/uuid"

type Bike struct {
	ID        uuid.UUID
	Lat       float64
	Lon       float64
	ImageUrl  string
	Available bool
}
