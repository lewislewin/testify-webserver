package shopify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UnmarshalOrder unmarshals the HTTP response body into an Order struct.
func UnmarshalOrder(resp *http.Response) (Order, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Order{}, fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Printf("Response body: %s\n", body)

	var orderWrapper struct {
		Order Order `json:"order"`
	}

	if err := json.Unmarshal(body, &orderWrapper); err != nil {
		return Order{}, fmt.Errorf("failed to decode order response: %v", err)
	}

	return orderWrapper.Order, nil
}

// MarshalOrder marshals an Order struct into a JSON byte slice.
func MarshalOrder(order Order) ([]byte, error) {
	orderWrapper := struct {
		Order Order `json:"order"`
	}{
		Order: order,
	}

	return json.Marshal(orderWrapper)
}

// IsEmptyOrder checks if the order is empty.
func IsEmptyOrder(order Order) bool {
	return order.ID == 0 && len(order.LineItems) == 0 && len(order.DiscountCodes) == 0
}
