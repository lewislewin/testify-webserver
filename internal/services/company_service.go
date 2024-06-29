package services

import (
	"database/sql"
	"log"

	"github.com/lewislewin/testify-webserver/internal/database"
	"github.com/lewislewin/testify-webserver/internal/models"

	"github.com/google/uuid"
)

func CreateCompany(company *models.Company) error {
	// Ensure the ID is a valid UUID
	company.ID = uuid.New()

	query := "INSERT INTO companies (id, name, website, email, owning_user, plan_tier) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, company.ID, company.Name, company.Website, company.Email, company.OwningUser, company.PlanTier)
	if err != nil {
		log.Printf("Error creating company: %v", err)
		return err
	}

	return nil
}

func GetCompany(id string) (*models.Company, error) {
	var company models.Company
	query := "SELECT id, name, website, email, owning_user, plan_tier FROM companies WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&company.ID, &company.Name, &company.Website, &company.Email, &company.OwningUser, &company.PlanTier)
	if err == sql.ErrNoRows {
		log.Println("Company not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving company: %v", err)
		return nil, err
	}

	return &company, nil
}

func UpdateCompany(id string, company *models.Company) error {
	query := "UPDATE companies SET name = ?, website = ?, email = ?, owning_user = ?, plan_tier = ? WHERE id = ?"
	_, err := database.DB.Exec(query, company.Name, company.Website, company.Email, company.OwningUser, company.PlanTier, id)
	if err != nil {
		log.Printf("Error updating company: %v", err)
		return err
	}

	return nil
}

func DeleteCompany(id string) error {
	query := "DELETE FROM companies WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting company: %v", err)
		return err
	}

	return nil
}

func ListCompanies() ([]models.Company, error) {
	rows, err := database.DB.Query("SELECT id, name, website, email, owning_user, plan_tier FROM companies")
	if err != nil {
		log.Printf("Error listing companies: %v", err)
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company
	for rows.Next() {
		var company models.Company
		if err := rows.Scan(&company.ID, &company.Name, &company.Website, &company.Email, &company.OwningUser, &company.PlanTier); err != nil {
			log.Printf("Error scanning company: %v", err)
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}
