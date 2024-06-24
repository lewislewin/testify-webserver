package models

import "github.com/google/uuid"

type TestResult struct {
	InternalID int
	ID         uuid.UUID
	TestPlanID uuid.UUID
	TestResult string
}
