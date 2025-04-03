package model

import (
	"github.com/google/uuid"
)

type Supplier struct {
	ID          uuid.UUID
	Name        string
	AddressID   uuid.UUID
	PhoneNumber string
}
