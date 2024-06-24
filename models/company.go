package models

import "github.com/google/uuid"

type Company struct {
	InternalID int
	ID         uuid.UUID
	Name       string
	Website    string
	Email      string
	OwningUser uuid.UUID
	PlanTier   int
}
