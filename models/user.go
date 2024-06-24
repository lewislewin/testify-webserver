package models

import "github.com/google/uuid"

type User struct {
	InternalID int
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email" binding:"required,email"`
	FirstName  string    `json:"first_name" binding:"required"`
	Surname    string    `json:"surname"`
	Password   string    `json:"password" binding:"required"`
	CompanyID  uuid.UUID `json:"company_id" binding:"required"`
}
