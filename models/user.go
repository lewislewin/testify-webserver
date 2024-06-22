package models

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CompanyID int    `json:"company_id" binding:"required"`
}
