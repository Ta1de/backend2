package model

import "github.com/google/uuid"

type Image struct {
	ID    uuid.UUID
	Image []byte
}
