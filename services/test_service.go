package services

import (
	"database/sql"
	"log"
	"testify-webserver/database"
	"testify-webserver/models"

	"github.com/google/uuid"
)

func CreateTest(test *models.Test) error {
	// Ensure the ID is a valid UUID
	test.ID = uuid.New()

	query := "INSERT INTO tests (id, test_suite_id, test_data) VALUES (?, ?, ?)"
	_, err := database.DB.Exec(query, test.ID, test.TestSuiteID, test.TestData)
	if err != nil {
		log.Printf("Error creating test: %v", err)
		return err
	}

	return nil
}

func GetTest(id string) (*models.Test, error) {
	var test models.Test
	query := "SELECT id, test_suite_id, test_data FROM tests WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&test.ID, &test.TestSuiteID, &test.TestData)
	if err == sql.ErrNoRows {
		log.Println("Test not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving test: %v", err)
		return nil, err
	}

	return &test, nil
}

func UpdateTest(id string, test *models.Test) error {
	query := "UPDATE tests SET test_suite_id = ?, test_data = ? WHERE id = ?"
	_, err := database.DB.Exec(query, test.TestSuiteID, test.TestData, id)
	if err != nil {
		log.Printf("Error updating test: %v", err)
		return err
	}

	return nil
}

func DeleteTest(id string) error {
	query := "DELETE FROM tests WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting test: %v", err)
		return err
	}

	return nil
}

func ListTests() ([]models.Test, error) {
	rows, err := database.DB.Query("SELECT id, test_suite_id, test_data FROM tests")
	if err != nil {
		log.Printf("Error listing tests: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tests []models.Test
	for rows.Next() {
		var test models.Test
		if err := rows.Scan(&test.ID, &test.TestSuiteID, &test.TestData); err != nil {
			log.Printf("Error scanning test: %v", err)
			return nil, err
		}
		tests = append(tests, test)
	}

	return tests, nil
}
