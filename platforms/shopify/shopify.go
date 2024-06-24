package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const SHOPIFY_URL_SUFFIX = ".myshopify.com"
const SHOPIFY_API_VERSION = "2024-04"

type Shopify struct {
	StoreStub string
	Token     string
	client    http.Client
}

func (s *Shopify) restRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var jsonBody []byte
	var err error

	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("https://%s%s/admin/api/%s%s", s.StoreStub, SHOPIFY_URL_SUFFIX, SHOPIFY_API_VERSION, endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Shopify-Access-Token", s.Token)
	req.Header.Add("Content-Type", "application/json")

	return s.client.Do(req)
}

func (s *Shopify) Authenticate(credentials map[string]string) error {
	s.StoreStub = credentials["store_domain"]
	s.Token = credentials["access_token"]
	s.client = http.Client{Timeout: 10 * time.Second}

	// Validate the connection by making a request to Shopify's API
	resp, err := s.restRequest("GET", "shop.json", nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
	}

	return nil
}
