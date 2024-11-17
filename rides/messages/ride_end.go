package messages

import "github.com/google/uuid"

type EndRide struct {
	RideID uuid.UUID `json:"ride_id"`
}
