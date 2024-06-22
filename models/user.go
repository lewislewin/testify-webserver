package models

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CompanyID int    `json:"company_id" binding:"required"`
}
