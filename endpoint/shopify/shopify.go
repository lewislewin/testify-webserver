package shopify

import (
	"fmt"
	"net/http"
)

const urlSuffix = ".myshopify.com"

type Client struct {
	shopName string
	url      string
}

func NewClient(name string) *Client {
	client := Client{
		shopName: name,
		url:      name + urlSuffix,
	}

	return &client
}

func (c *Client) AuthenticateWithEndpoint() (bool, error) {
	fmt.Println("I'm in pkg shopify, AuthenticateWithEndpoint()")

	return true, nil
}

func (c *Client) GetProducts() (http.Response, error) {
	fmt.Println("I'm in pkg shopify, GetProducts()")

	return http.Response{}, nil
}

func (c *Client) CreateOrder() (http.Response, error) {
	fmt.Println("I'm in pkg shopify, CreateOrder()")

	return http.Response{}, nil
}

// ValidateBulkProducts implements endpoint.EndpointClient.
func (c *Client) ValidateBulkProducts() (bool, error) {
	panic("unimplemented")
}

// ValidateProduct implements endpoint.EndpointClient.
func (c *Client) ValidateProduct() (bool, error) {
	panic("unimplemented")
}

// ValidateProduct implements endpoint.EndpointClient.
func (c *Client) ValidateProduct2() (bool, error) {
	panic("unimplemented")
}
