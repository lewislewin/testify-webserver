package models

import (
	"github.com/google/uuid"
)

type Test struct {
	InternalID  int
	ID          uuid.UUID `json:"id"`
	TestSuiteID uuid.UUID `json:"test_suite_id" binding:"required"`
	TestData    string    `json:"test_data" binding:"required"`
}
