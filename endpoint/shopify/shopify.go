package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	urlSuffix  = ".myshopify.com"
	apiVersion = "2024-04"
)

type Client struct {
	shopName string
	url      string
	apiKey   string
	client   *http.Client
}

func NewClient(name, apiKey string) *Client {
	client := &Client{
		shopName: name,
		url:      "https://" + name + urlSuffix + "/admin/api/" + apiVersion,
		apiKey:   apiKey,
		client:   &http.Client{Timeout: 10 * time.Second},
	}

	return client
}

func (c *Client) AuthenticateWithEndpoint() (bool, error) {
	// In Shopify, usually, authentication is done by including the API key in the headers.
	// So this function might not be necessary for every request, but we can check if the API key is valid.
	req, err := http.NewRequest("GET", c.url+"/shop.json", nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("X-Shopify-Access-Token", c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
}

func (c *Client) restRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var jsonBody []byte
	var err error

	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.url+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Shopify-Access-Token", c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	return c.client.Do(req)
}

func (c *Client) graphqlRequest(query string, variables map[string]interface{}) (*http.Response, error) {
	requestBody := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url+"/graphql.json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Shopify-Access-Token", c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	return c.client.Do(req)
}

func (c *Client) GetProducts() (*http.Response, error) {
	fmt.Println("I'm in pkg shopify, GetProducts()")

	query := `
		{
			products(first: 10) {
				edges {
					node {
						id
						title
						handle
					}
				}
			}
		}
	`
	resp, err := c.graphqlRequest(query, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateOrder(orderData map[string]interface{}) (*http.Response, error) {
	fmt.Println("I'm in pkg shopify, CreateOrder()")

	resp, err := c.restRequest("POST", "/orders.json", orderData)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ValidateBulkProducts implements endpoint.EndpointClient.
func (c *Client) ValidateBulkProducts() (bool, error) {
	panic("unimplemented")
}

// ValidateProduct implements endpoint.EndpointClient.
func (c *Client) ValidateProduct() (bool, error) {
	panic("unimplemented")
}
