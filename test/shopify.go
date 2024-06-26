package main

import (
	"fmt"
	"testify-webserver/models"
	"testify-webserver/platforms"
	"testify-webserver/platforms/shopify"
	"testify-webserver/services"

	"github.com/google/uuid"
)

func main() {
	endpoint_type_uuid, _ := uuid.Parse("fa4b981a-6ebb-48e5-bd70-d781cce29d58")
	services.CreateEndpointType(&models.EndpointType{
		InternalID: 1,
		ID:         endpoint_type_uuid,
		Name:       "shopify",
	})
	ep, err := platforms.EndpointFactory(&models.Endpoint{
		InternalID:   1,
		ID:           uuid.New(),
		Name:         "shopify test",
		CredentialID: uuid.New(),
		CompanyID:    uuid.New(),
		EndpointType: endpoint_type_uuid,
	})
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
}
