// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type BikeResponse struct {
	ID       uuid.UUID `json:"id"`
	ImageURL *string   `json:"imageUrl,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type RideResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	BikeID    uuid.UUID `json:"bikeId"`
	StartTime string    `json:"startTime"`
	EndTime   *string   `json:"endTime,omitempty"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
}