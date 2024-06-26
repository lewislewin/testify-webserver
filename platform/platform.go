package platform

// package endpoint is a reusable package to map the high level functionality of each environment

import (
	"fmt"
	"net/http"

	"testify-webserver/models"
	"testify-webserver/platform_implementations/shopify"
)

const ShopifyEndpointType = "shopfiy"
const BigcommerceEndpointType = "bigcommerce"
const OtherEndpointType = "other"

type Client struct {
	Endpoint       models.Endpoint
	EndpointClient EndpointClient
}

type EndpointClient interface {
	Authenticate() error
	GetProducts() (http.Response, error)
	CreateOrder() (http.Response, error)
}

type Credentials struct {
	UserName string
	Token    string
}

func NewClient(ep models.Endpoint) (*Client, error) {

	client := Client{
		Endpoint: ep,
	}

	if ep.EndpointType.String() == ShopifyEndpointType {
		client.EndpointClient = shopify.NewClient(ep.Name)
	}

	return &client, nil
}

func (c *Client) AuthoriseConnectionToEndpoint() error {
	fmt.Println("I'm in AuthoriseConnectionToEndpoint()")

	c.EndpointClient.Authenticate()
	return nil
}

func (c *Client) GetProducts() (http.Response, error) {

	return c.EndpointClient.GetProducts()
}
