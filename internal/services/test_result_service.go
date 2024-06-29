package services

import (
	"database/sql"
	"log"

	"github.com/lewislewin/testify-webserver/internal/database"

	"github.com/lewislewin/testify-webserver/internal/models"

	"github.com/google/uuid"
)

func CreateTestResult(testResult *models.TestResult) error {
	// Ensure the ID is a valid UUID
	testResult.ID = uuid.New()

	query := "INSERT INTO test_results (id, test_suite_id, test_result) VALUES (?, ?, ?)"
	_, err := database.DB.Exec(query, testResult.ID, testResult.TestSuiteID, testResult.TestResult)
	if err != nil {
		log.Printf("Error creating test result: %v", err)
		return err
	}

	return nil
}

func GetTestResult(id string) (*models.TestResult, error) {
	var testResult models.TestResult
	query := "SELECT id, test_suite_id, test_result FROM test_results WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&testResult.ID, &testResult.TestSuiteID, &testResult.TestResult)
	if err == sql.ErrNoRows {
		log.Println("Test result not found")
		return nil, err
	} else if err != nil {
		log.Printf("Error retrieving test result: %v", err)
		return nil, err
	}

	return &testResult, nil
}

func UpdateTestResult(id string, testResult *models.TestResult) error {
	query := "UPDATE test_results SET test_suite_id = ?, test_result = ? WHERE id = ?"
	_, err := database.DB.Exec(query, testResult.TestSuiteID, testResult.TestResult, id)
	if err != nil {
		log.Printf("Error updating test result: %v", err)
		return err
	}

	return nil
}

func DeleteTestResult(id string) error {
	query := "DELETE FROM test_results WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting test result: %v", err)
		return err
	}

	return nil
}

func ListTestResults() ([]models.TestResult, error) {
	rows, err := database.DB.Query("SELECT id, test_suite_id, test_result FROM test_results")
	if err != nil {
		log.Printf("Error listing test results: %v", err)
		return nil, err
	}
	defer rows.Close()

	var testResults []models.TestResult
	for rows.Next() {
		var testResult models.TestResult
		if err := rows.Scan(&testResult.ID, &testResult.TestSuiteID, &testResult.TestResult); err != nil {
			log.Printf("Error scanning test result: %v", err)
			return nil, err
		}
		testResults = append(testResults, testResult)
	}

	return testResults, nil
}
