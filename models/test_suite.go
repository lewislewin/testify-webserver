package models

import "github.com/google/uuid"

type TestSuite struct {
	InternalID       int
	ID               uuid.UUID
	EndpointID       uuid.UUID
	Name             string
	MostRecentResult uuid.UUID
}
