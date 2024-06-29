package models

import "github.com/google/uuid"

type TestResult struct {
	InternalID  int
	ID          uuid.UUID `json:"id"`
	TestSuiteID uuid.UUID `json:"test_suite_id"`
	TestResult  string    `json:"test_result"`
}
