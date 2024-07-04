package main

import (
	"fmt"

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
		ID:                    6098106057050,
		Email:                 "jane@example.com",
		ContactEmail:          "jane@example.com",
		FinancialStatus:       "paid",
		Currency:              "GBP",
		SubtotalPrice:         "220.64",
		TotalDiscounts:        "32.36",
		TotalPrice:            "222.64",
		CurrentSubtotalPrice:  "191.86",
		CurrentTotalDiscounts: "28.14",
		CurrentTotalPrice:     "193.86",
		CurrentTotalTax:       "0.00",
		Confirmed:             true,
		TotalLineItemsPrice:   "253.00",
		TotalOutstanding:      "0.00",
		TotalShippingPriceSet: shopify.PriceSet{
			ShopMoney: shopify.Money{
				Amount:       "2.00",
				CurrencyCode: "GBP",
			},
			PresentmentMoney: shopify.Money{
				Amount:       "2.00",
				CurrencyCode: "GBP",
			},
		},
		DiscountCodes: []shopify.DiscountCode{
			{
				Code:   "dtest",
				Amount: "2.00",
				Type:   "fixed_amount",
			},
			{
				Code:   "B",
				Amount: "30.36",
				Type:   "percentage",
			},
		},
		Tags:        "tag2, testtag",
		Test:        true,
		OrderNumber: 1007,
		BillingAddress: shopify.Address{
			FirstName: "Jane",
			LastName:  "Smith",
			Address1:  "123 Fake Street",
			Phone:     "+447454333056",
			City:      "Fakecity",
			Province:  "London",
			Country:   "UK",
			Zip:       "EC1A 1BB",
		},
		ShippingAddress: shopify.Address{
			FirstName: "Jane",
			LastName:  "Smith",
			Address1:  "123 Fake Street",
			Phone:     "+44 777 777 7777",
			City:      "Fakecity",
			Province:  "London",
			Country:   "UK",
			Zip:       "EC1A 1BB",
		},
		LineItems: []shopify.LineItem{
			{
				VariantID: 48478204526938,
				Title:     "Oh no - Medium",
				Quantity:  1,
				Price:     "44.00",
			},
			{
				VariantID: 48478204559706,
				Title:     "Oh no - Large",
				Quantity:  1,
				Price:     "33.00",
			},
			{
				VariantID: 48478204592474,
				Title:     "Oh no - Small",
				Quantity:  4,
				Price:     "44.00",
			},
			{
				VariantID: 48478171398490,
				Title:     "Test product two",
				Quantity:  1,
				Price:     "0.00",
			},
			{
				VariantID: 48478322786650,
				Title:     "Test product two 2",
				Quantity:  1,
				Price:     "0.00",
			},
		},
		Transactions: []shopify.Transaction{
			{
				Kind:    "capture",
				Status:  "success",
				Amount:  "222.64",
				Gateway: "visa",
			},
		},
	}

	resp, err := ep.CreateOrder(order)
	if err != nil {
		panic(fmt.Sprintf("Failed to create order: %v", err))
	}

	fmt.Println("Order created successfully with status code:", resp.StatusCode)

	// Get products
	products, err := ep.ValidateProducts()
	if err != nil {
		panic(fmt.Sprintf("Failed to get products: %v", err))
	}

	fmt.Println(products)
}
