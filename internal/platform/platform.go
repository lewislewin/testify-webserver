package platform

import (
	"fmt"
	"net/http"

	"github.com/lewislewin/testify-webserver/internal/platform_implementations/shopify"
)

const (
	ShopifyEndpointType     = "shopify"
	BigcommerceEndpointType = "bigcommerce"
	OtherEndpointType       = "other"
)

type Client struct {
	Platform       Platform
	PlatformClient PlatformClient
}

type Platform struct {
	PlatformType string
	CredentialID string
}

type PlatformClient interface {
	Authenticate() error
	ValidateProducts() (*http.Response, error)
	CreateOrder(order interface{}) (*http.Response, error)
	GetOrder(id int) (*http.Response, error)
}

func NewClient(platform Platform) (*Client, error) {
	client := &Client{
		Platform: platform,
	}

	switch platform.PlatformType {
	case ShopifyEndpointType:
		client.PlatformClient = shopify.NewClient(platform.CredentialID)
	default:
		return nil, fmt.Errorf("unsupported platform type: %s", platform.PlatformType)
	}

	return client, nil
}

func (c *Client) Authenticate() error {
	fmt.Println("Authenticating with platform")
	return c.PlatformClient.Authenticate()
}

func (c *Client) ValidateProducts() (*http.Response, error) {
	return c.PlatformClient.ValidateProducts()
}

func (c *Client) CreateOrder(order interface{}) (*http.Response, error) {
	return c.PlatformClient.CreateOrder(order)
}

func (c *Client) GetOrder(id int) (*http.Response, error) {
	return c.PlatformClient.GetOrder(id)
}
