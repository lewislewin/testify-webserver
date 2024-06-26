package services

import (
	"database/sql"
	"fmt"
	"log"
	"testify-webserver/database"
	"testify-webserver/models"
	"testify-webserver/platforms"
	"testify-webserver/platforms/shopify"

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

func RunTestSuite(id string) error {
	// This is just example code
	endpoint_type_uuid, _ := uuid.Parse("fa4b981a-6ebb-48e5-bd70-d781cce29d58")
	CreateEndpointType(&models.EndpointType{
		InternalID: 1,
		ID:         endpoint_type_uuid,
		Name:       "shopify",
	})
	ep, err := platforms.PlatformFactory(&platforms.Platform{PlatformType: "shopify", CredentialID: endpoint_type_uuid.String()})
	fmt.Printf("%v", ep)
	if err != nil {
		panic("OOpsie")
	}
	ep.CreateOrder(shopify.Order{
		ID:        "ID",
		Title:     "title",
		Price:     "price",
		Inventory: 1,
	})
	products, err := ep.GetProducts()
	if err != nil {
		panic("Failed to get products")
	}

	// Type assert the products to the expected type
	if productList, ok := products.([]shopify.Product); ok {
		// Iterate over the products and print their details
		for _, product := range productList {
			fmt.Printf("ID: %s, Title: %s, Price: %s, Inventory: %d\n",
				product.ID, product.Title, product.Price, product.Inventory)
		}
	} else {
		fmt.Println("Failed to type assert products")
	}
	return nil
}
