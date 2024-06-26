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

// Authenticate implements platform.EndpointClient.
func (c *Client) Authenticate() error {
	panic("unimplemented")
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
