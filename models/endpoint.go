package models

import "github.com/google/uuid"

type Endpoint struct {
	InternalID   int
	ID           uuid.UUID
	Name         string
	CredentialID uuid.UUID
	CompanyID    uuid.UUID
	EndpointType uuid.UUID
}
