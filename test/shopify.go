package main

import (
	"fmt"
	"testify-webserver/database"
	"testify-webserver/models"
	"testify-webserver/platforms"
	"testify-webserver/platforms/shopify"
	"testify-webserver/services"

	"github.com/google/uuid"
)

func main() {
	database.InitDB()
	endpoint_type_uuid, _ := uuid.Parse("fa4b981a-6ebb-48e5-bd70-d781cce29d58")
	services.CreateEndpointType(&models.EndpointType{
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
}
