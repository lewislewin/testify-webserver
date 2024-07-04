package main

import (
	"fmt"
	"log"

	"github.com/lewislewin/testify-webserver/internal/database"
	"github.com/lewislewin/testify-webserver/internal/models"
	"github.com/lewislewin/testify-webserver/internal/platform"
	"github.com/lewislewin/testify-webserver/internal/platform_implementations/shopify"
	"github.com/lewislewin/testify-webserver/internal/services"

	"github.com/google/uuid"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create an endpoint type for Shopify
	endpointTypeUUID, err := uuid.Parse("fa4b981a-6ebb-48e5-bd70-d781cce29d58")
	if err != nil {
		log.Fatalf("Failed to parse UUID: %v", err)
	}

	err = services.CreateEndpointType(&models.EndpointType{
		InternalID: 1,
		ID:         endpointTypeUUID,
		Name:       "shopify",
	})
	if err != nil {
		log.Fatalf("Failed to create endpoint type: %v", err)
	}

	// Create a new client for Shopify platform
	platformType := platform.Platform{
		PlatformType: platform.ShopifyEndpointType,
		CredentialID: "your-credential-id-here",
	}

	ep, err := platform.NewClient(platformType)
	if err != nil {
		log.Fatalf("Failed to create platform client: %v", err)
	}

	// Authenticate the client
	err = ep.Authenticate()
	if err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	// Get an order
	resp, err := ep.GetOrder(6102131802458)
	if err != nil {
		log.Fatalf("Failed to get order: %v", err)
	}
	defer resp.Body.Close()

	order, err := shopify.UnmarshalOrder(resp)
	if err != nil {
		log.Fatalf("Failed to unmarshal order: %v", err)
	}

	// Check if order is populated
	if shopify.IsEmptyOrder(order) {
		log.Fatalf("Order is empty: %v", order)
	} else {
		fmt.Printf("Retrieved order: %+v\n", order)
	}

	// Create new orders in a loop
	for i := 0; i < 20; i++ {
		resp, err = ep.CreateOrder(order)
		if err != nil {
			log.Fatalf("Failed to create order: %v", err)
		}
		defer resp.Body.Close()

		fmt.Println("Order created successfully with status code:", resp.StatusCode)
	}

	// Get products
	products, err := ep.ValidateProducts()
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	fmt.Println(products)
}
