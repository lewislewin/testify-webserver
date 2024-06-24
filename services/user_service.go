package services

import (
	"database/sql"
	"log"
	"testify-webserver/database"
	"testify-webserver/models"

	"github.com/google/uuid"
)

func CreateUser(user *models.User) error {
	// Ensure the ID is a valid UUID
	user.ID = uuid.New()

	query := "INSERT INTO users (id, email, first_name, surname, password, company_id) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, user.ID, user.Email, user.FirstName, user.Surname, user.Password, user.CompanyID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	query := "SELECT id, email, first_name, surname, password, company_id FROM users WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.Surname, &user.Password, &user.CompanyID)
	if err == sql.ErrNoRows {
		log.Println("User not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id string, user *models.User) error {
	query := "UPDATE users SET email = ?, first_name = ?, surname = ?, password = ?, company_id = ? WHERE id = ?"
	_, err := database.DB.Exec(query, user.Email, user.FirstName, user.Surname, user.Password, user.CompanyID, id)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}

func ListUsers() ([]models.User, error) {
	rows, err := database.DB.Query("SELECT id, email, first_name, surname, password, company_id FROM users")
	if err != nil {
		log.Printf("Error listing users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.Surname, &user.Password, &user.CompanyID); err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
