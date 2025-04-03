package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID
	ClientName       string
	ClientSurname    string
	Birthday         time.Time
	Gender           string
	RegistrationDate time.Time
	AddressID        uuid.UUID
}
