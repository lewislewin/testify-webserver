package models

import "github.com/google/uuid"

type Permission struct {
	InternalID int
	ID         uuid.UUID
	Name       string
}
