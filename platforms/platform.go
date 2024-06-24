package platforms

import (
	"fmt"
	"testify-webserver/credentials"
	"testify-webserver/models"
	"testify-webserver/platforms/bigcommerce"
	"testify-webserver/platforms/shopify"
	"testify-webserver/services"
)

type Platform interface {
	Authenticate(credentials map[string]string) error
	GetProducts() (interface{}, error)
}

func EndpointFactory(ep models.Endpoint) (Platform, error) {
	var endpoint Platform
	endpointType, err := services.GetEndpointType(ep.EndpointType.String())
	if err != nil {
		return nil, err
	}

	switch endpointType.Name {
	case "shopify":
		endpoint = &shopify.Shopify{}
	case "bigcommerce":
		endpoint = &bigcommerce.BigCommerce{}
	default:
		return nil, fmt.Errorf("unsupported endpoint type")
	}

	credentials, err := credentials.GetCredentialById(string(ep.CredentialID.String()))
	if err != nil {
		return nil, err
	}

	err = endpoint.Authenticate(credentials)
	if err != nil {
		return nil, err
	}
	return endpoint, nil
}
