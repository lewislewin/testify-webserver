package models

import "github.com/google/uuid"

type Company struct {
	InternalID int
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Website    string    `json:"website"`
	Email      string    `json:"email" binding:"required"`
	OwningUser uuid.UUID `json:"owning_user"`
	PlanTier   int       `json:"tier"`
}
