package platforms

import (
	"fmt"
	"testify-webserver/credentials"
	"testify-webserver/platforms/bigcommerce"
	"testify-webserver/platforms/shopify"
)

type Platform struct {
	PlatformType string
	CredentialID string
}

type PlatformClient interface {
	Authenticate(credentials map[string]string) error
	GetProducts() (interface{}, error)
	CreateOrder(order interface{}) error
}

func PlatformFactory(platform *Platform) (PlatformClient, error) {
	var endpoint PlatformClient

	switch platform.PlatformType {
	case "shopify":
		endpoint = &shopify.Shopify{}
	case "bigcommerce":
		endpoint = &bigcommerce.BigCommerce{}
	default:
		return nil, fmt.Errorf("unsupported endpoint type")
	}

	credentials, err := credentials.GetCredentialById(platform.CredentialID)
	if err != nil {
		return nil, err
	}

	err = endpoint.Authenticate(credentials)
	if err != nil {
		return nil, err
	}
	return endpoint, nil
}

// Generic order struct
type Order struct {
	ID   int
	Name string
}
