package models

type Permission struct {
	ID          int
	UserID      int
	CompanyID   int
	AccessLevel int
}
