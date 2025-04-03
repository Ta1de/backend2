package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID             uuid.UUID
	Name           string
	Category       string
	Price          float64
	AvailableStock int
	LastUpdateDate time.Time
	SupplierID     uuid.UUID
	ImageID        *uuid.UUID
}
