package models

import "github.com/google/uuid"

type EndpointType struct {
	InternalID int
	ID         uuid.UUID
	Name       string
}
