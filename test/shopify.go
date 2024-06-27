package main

import (
	"fmt"
	"testify-webserver/database"
	"testify-webserver/models"
	"testify-webserver/platform"
	"testify-webserver/services"

	"github.com/google/uuid"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create an endpoint type for Shopify
	endpointTypeUUID, _ := uuid.Parse("fa4b981a-6ebb-48e5-bd70-d781cce29d58")
	services.CreateEndpointType(&models.EndpointType{
		InternalID: 1,
		ID:         endpointTypeUUID,
		Name:       "shopify",
	})

	// Create a new client for Shopify platform
	platformType := platform.Platform{
		PlatformType: platform.ShopifyEndpointType,
		CredentialID: "your-credential-id-here",
	}

	ep, err := platform.NewClient(platformType)
	if err != nil {
		panic(fmt.Sprintf("Failed to create platform client: %v", err))
	}

	// Authenticate the client
	err = ep.Authenticate()
	if err != nil {
		panic(fmt.Sprintf("Failed to authenticate: %v", err))
	}

	// Create an order (example order data should be passed)
	orderData := map[string]interface{}{
		"order": map[string]interface{}{
			"line_items": []map[string]interface{}{
				{
					"title":    "Example Item",
					"price":    "10.00",
					"quantity": 1,
				},
			},
		},
	}

	_, err = ep.CreateOrder(orderData)
	if err != nil {
		panic(fmt.Sprintf("Failed to create order: %v", err))
	}

	// Get products
	products, err := ep.ValidateProducts()
	if err != nil {
		panic(fmt.Sprintf("Failed to get products: %v", err))
	}

	fmt.Println(products)
}
