package model

import (
	"github.com/google/uuid"
)

type Address struct {
	ID      uuid.UUID
	Country string
	City    string
	Street  string
}
