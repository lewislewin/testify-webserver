package services

import (
	"database/sql"
	"log"
	"testify-webserver/database"
	"testify-webserver/models"

	"github.com/google/uuid"
)

func CreateEndpoint(endpoint *models.Endpoint) error {
	endpoint.ID = uuid.New()

	query := "INSERT INTO endpoints (id, name, credential_id, company_id, endpoint_type) VALUES (?, ?, ?, ?, ?)"
	_, err := database.DB.Exec(query, endpoint.ID, endpoint.Name, endpoint.CredentialID, endpoint.CompanyID, endpoint.EndpointType)
	if err != nil {
		log.Printf("Error creating endpoint: %v", err)
		return err
	}

	return nil
}

func GetEndpoint(id string) (*models.Endpoint, error) {
	var endpoint models.Endpoint
	query := "SELECT id, name, credential_id, company_id, endpoint_type FROM endpoints WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&endpoint.ID, &endpoint.Name, &endpoint.CredentialID, &endpoint.CompanyID, &endpoint.EndpointType)
	if err == sql.ErrNoRows {
		log.Println("Endpoint not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving endpoint: %v", err)
		return nil, err
	}

	return &endpoint, nil
}

func UpdateEndpoint(id string, endpoint *models.Endpoint) error {
	query := "UPDATE endpoints SET name = ?, credential_id = ?, company_id = ?, endpoint_type = ? WHERE id = ?"
	_, err := database.DB.Exec(query, endpoint.Name, endpoint.CredentialID, endpoint.CompanyID, endpoint.EndpointType, id)
	if err != nil {
		log.Printf("Error updating endpoint: %v", err)
		return err
	}

	return nil
}

func DeleteEndpoint(id string) error {
	query := "DELETE FROM endpoints WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting endpoint: %v", err)
		return err
	}

	return nil
}

func ListEndpoints() ([]models.Endpoint, error) {
	rows, err := database.DB.Query("SELECT id, name, credential_id, company_id, endpoint_type FROM endpoints")
	if err != nil {
		log.Printf("Error listing endpoints: %v", err)
		return nil, err
	}
	defer rows.Close()

	var endpoints []models.Endpoint
	for rows.Next() {
		var endpoint models.Endpoint
		if err := rows.Scan(&endpoint.ID, &endpoint.Name, &endpoint.CredentialID, &endpoint.CompanyID, &endpoint.EndpointType); err != nil {
			log.Printf("Error scanning endpoint: %v", err)
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}
