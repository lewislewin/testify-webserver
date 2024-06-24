package models

import "github.com/google/uuid"

type TestSuite struct {
	InternalID       int
	ID               uuid.UUID `json:"id"`
	EndpointID       uuid.UUID `json:"endpoint_id" binding:"required"`
	Name             string    `json:"name" binding:"required"`
	MostRecentResult uuid.UUID `json:"most_recent_result"`
}
