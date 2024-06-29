package services

import (
	"database/sql"
	"log"

	"github.com/lewislewin/testify-webserver/internal/database"

	"github.com/lewislewin/testify-webserver/internal/models"

	"github.com/google/uuid"
)

func CreateTestSuite(testSuite *models.TestSuite) error {
	// Ensure the ID is a valid UUID
	testSuite.ID = uuid.New()

	query := "INSERT INTO test_suites (id, endpoint_id, name, most_recent_result) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(query, testSuite.ID, testSuite.EndpointID, testSuite.Name, testSuite.MostRecentResult)
	if err != nil {
		log.Printf("Error creating test suite: %v", err)
		return err
	}

	return nil
}

func GetTestSuite(id string) (*models.TestSuite, error) {
	var testSuite models.TestSuite
	query := "SELECT id, endpoint_id, name, most_recent_result FROM test_suites WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&testSuite.ID, &testSuite.EndpointID, &testSuite.Name, &testSuite.MostRecentResult)
	if err == sql.ErrNoRows {
		log.Println("Test suite not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving test suite: %v", err)
		return nil, err
	}

	return &testSuite, nil
}

func UpdateTestSuite(id string, testSuite *models.TestSuite) error {
	query := "UPDATE test_suites SET endpoint_id = ?, name = ?, most_recent_result = ? WHERE id = ?"
	_, err := database.DB.Exec(query, testSuite.EndpointID, testSuite.Name, testSuite.MostRecentResult, id)
	if err != nil {
		log.Printf("Error updating test suite: %v", err)
		return err
	}

	return nil
}

func DeleteTestSuite(id string) error {
	query := "DELETE FROM test_suites WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting test suite: %v", err)
		return err
	}

	return nil
}

func ListTestSuites() ([]models.TestSuite, error) {
	rows, err := database.DB.Query("SELECT id, endpoint_id, name, most_recent_result FROM test_suites")
	if err != nil {
		log.Printf("Error listing test suites: %v", err)
		return nil, err
	}
	defer rows.Close()

	var testSuites []models.TestSuite
	for rows.Next() {
		var testSuite models.TestSuite
		if err := rows.Scan(&testSuite.ID, &testSuite.EndpointID, &testSuite.Name, &testSuite.MostRecentResult); err != nil {
			log.Printf("Error scanning test suite: %v", err)
			return nil, err
		}
		testSuites = append(testSuites, testSuite)
	}

	return testSuites, nil
}
