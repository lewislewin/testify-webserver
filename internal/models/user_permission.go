package models

import "github.com/google/uuid"

type UserPermission struct {
	InternalID   int
	ID           uuid.UUID
	UserID       uuid.UUID
	PermissionID uuid.UUID
}
