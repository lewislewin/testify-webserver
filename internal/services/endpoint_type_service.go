package services

import (
	"database/sql"
	"log"

	"github.com/lewislewin/testify-webserver/internal/database"

	"github.com/lewislewin/testify-webserver/internal/models"

	"github.com/google/uuid"
)

func CreateEndpointType(endpointType *models.EndpointType) error {
	// Ensure the ID is a valid UUID
	endpointType.ID = uuid.New()

	query := "INSERT INTO endpoint_types (id, name) VALUES (?, ?)"
	_, err := database.DB.Exec(query, endpointType.ID.String(), endpointType.Name)
	if err != nil {
		log.Printf("Error creating endpoint type: %v", err)
		return err
	}

	return nil
}

func GetEndpointType(id string) (*models.EndpointType, error) {
	var endpointType models.EndpointType
	query := "SELECT id, name FROM endpoint_types WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&endpointType.ID, &endpointType.Name)
	if err == sql.ErrNoRows {
		log.Println("Endpoint type not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving endpoint type: %v", err)
		return nil, err
	}

	return &endpointType, nil
}

func UpdateEndpointType(id string, endpointType *models.EndpointType) error {
	query := "UPDATE endpoint_types SET name = ? WHERE id = ?"
	_, err := database.DB.Exec(query, endpointType.Name, id)
	if err != nil {
		log.Printf("Error updating endpoint type: %v", err)
		return err
	}

	return nil
}

func DeleteEndpointType(id string) error {
	query := "DELETE FROM endpoint_types WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting endpoint type: %v", err)
		return err
	}

	return nil
}

func ListEndpointTypes() ([]models.EndpointType, error) {
	rows, err := database.DB.Query("SELECT id, name FROM endpoint_types")
	if err != nil {
		log.Printf("Error listing endpoint types: %v", err)
		return nil, err
	}
	defer rows.Close()

	var endpointTypes []models.EndpointType
	for rows.Next() {
		var endpointType models.EndpointType
		if err := rows.Scan(&endpointType.ID, &endpointType.Name); err != nil {
			log.Printf("Error scanning endpoint type: %v", err)
			return nil, err
		}
		endpointTypes = append(endpointTypes, endpointType)
	}

	return endpointTypes, nil
}
