package models

import (
	"github.com/google/uuid"
)

type Test struct {
	InternalID int
	ID         uuid.UUID
	TestPlanID uuid.UUID
	TestData   string
}
