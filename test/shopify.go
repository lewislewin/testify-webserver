package main

import (
	"fmt"
	"testify-webserver/database"
	"testify-webserver/models"
	"testify-webserver/platform"
	"testify-webserver/platform_implementations/shopify"
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

	// Create a Shopify order
	order := shopify.Order{
		LineItems: []shopify.LineItem{
			{
				VariantID: 447654529,
				Quantity:  1,
			},
		},
		Email: "jane@example.com",
		Phone: "18885551234",
		BillingAddress: shopify.Address{
			FirstName: "John",
			LastName:  "Smith",
			Address1:  "123 Fake Street",
			Phone:     "555-555-5555",
			City:      "Fakecity",
			Province:  "Ontario",
			Country:   "Canada",
			Zip:       "K2P 1L4",
		},
		ShippingAddress: shopify.Address{
			FirstName: "Jane",
			LastName:  "Smith",
			Address1:  "123 Fake Street",
			Phone:     "777-777-7777",
			City:      "Fakecity",
			Province:  "Ontario",
			Country:   "Canada",
			Zip:       "K2P 1L4",
		},
		Transactions: []shopify.Transaction{
			{
				Kind:   "sale",
				Status: "success",
				Amount: 50.0,
			},
		},
		FinancialStatus: "paid",
	}

	_, err = ep.CreateOrder(order)
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
