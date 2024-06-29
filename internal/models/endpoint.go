package models

import "github.com/google/uuid"

type Endpoint struct {
	InternalID   int
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name" binding:"required"`
	CredentialID uuid.UUID `json:"credential_id"`
	CompanyID    uuid.UUID `json:"company_id" binding:"required"`
	EndpointType uuid.UUID `json:"endpoint_type" binding:"required"`
}
