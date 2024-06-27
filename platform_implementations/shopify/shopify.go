package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testify-webserver/credentials"
	"time"
)

const urlSuffix = ".myshopify.com"
const apiVersion = "2024-04"

type Client struct {
	storeDomain string
	token       string
	httpClient  http.Client
}

func NewClient(credentialID string) *Client {
	creds, err := credentials.GetCredentialById(credentialID)
	if err != nil {
		panic("Failed to get credentials")
	}
	client := Client{
		storeDomain: creds["store_domain"],
		token:       creds["access_token"],
		httpClient:  http.Client{Timeout: 10 * time.Second},
	}
	return &client
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

	url := fmt.Sprintf("https://%s%s/admin/api/%s%s", c.storeDomain, urlSuffix, apiVersion, endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Shopify-Access-Token", c.token)
	req.Header.Add("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c *Client) Authenticate() error {
	fmt.Println("I'm in pkg shopify, Authenticate()")
	resp, err := c.restRequest("GET", "/shop.json", nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) ValidateProducts() (*http.Response, error) {
	fmt.Println("I'm in pkg shopify, ValidateProducts()")
	return c.restRequest("GET", "/products.json", nil)
}

func (c *Client) CreateOrder(order interface{}) (*http.Response, error) {
	fmt.Println("I'm in pkg shopify, CreateOrder()")
	orderData := map[string]interface{}{
		"order": order,
	}

	resp, err := c.restRequest("POST", "/orders.json", orderData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Failed to create order with status code: %d\n", resp.StatusCode)
		fmt.Printf("Response Body: %s\n", string(body))
		return resp, fmt.Errorf("failed to create order with status code: %d", resp.StatusCode)
	}

	fmt.Println("Response Body:", string(body))
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Headers:", resp.Header)

	return resp, nil
}
