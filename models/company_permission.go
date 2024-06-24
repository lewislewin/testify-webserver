package models

import "github.com/google/uuid"

type CompanyPermission struct {
	InternalID   int
	ID           uuid.UUID
	CompanyID    uuid.UUID
	PermissionID uuid.UUID
}
