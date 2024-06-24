package endpoint

// package endpoint is a reusable package to map the high level functionality of each environment

import (
	"fmt"
	"net/http"

	"testify-webserver/endpoint/shopify"
)

const ShopifyEndpointType = "shopify"
const BigcommerceEndpointType = "bigcommerce"
const OtherEndpointType = "other"

type Client struct {
	Endpoint       Endpoint
	EndpointClient EndpointClient
}

type EndpointClient interface {
	// Authentication
	AuthenticateWithEndpoint() (bool, error)

	// Product Methods
	GetProducts() (*http.Response, error)
	ValidateProduct() (bool, error)
	ValidateBulkProducts() (bool, error)

	// Order Methods
	CreateOrder(orderData map[string]interface{}) (*http.Response, error)
}

// Endpoint:
// Contains generic information & interfaces related to a platform's shop
type Endpoint struct {
	Name      string
	Type      string
	CompanyID int64
	URL       string
	Credentials
}

type Credentials struct {
	UserName string
	Token    string
}

func NewClient(ep Endpoint) (*Client, error) {
	client := &Client{
		Endpoint: ep,
	}

	if ep.Type == ShopifyEndpointType {
		client.EndpointClient = shopify.NewClient(ep.Name, ep.Credentials.Token)
	}

	return client, nil
}

func (c *Client) AuthoriseConnectionToEndpoint() error {
	fmt.Println("I'm in AuthoriseConnectionToEndpoint()")

	authenticated, err := c.EndpointClient.AuthenticateWithEndpoint()
	if err != nil {
		return err
	}
	if !authenticated {
		return fmt.Errorf("authentication failed")
	}
	return nil
}

func (c *Client) GetProducts() (*http.Response, error) {
	return c.EndpointClient.GetProducts()
}
