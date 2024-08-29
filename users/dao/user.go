package dao

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
}
