package bigcommerce

type Order struct {
	// Define the fields required for a BigCommerce order
	CustomerID      int           `json:"customer_id"`
	Items           []OrderItem   `json:"items"`
	BillingAddress  Address       `json:"billing_address"`
	ShippingAddress Address       `json:"shipping_address"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	Status          string        `json:"status"`
}

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address1"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
	Phone     string `json:"phone"`
}

type PaymentMethod struct {
	Method string `json:"method"`
}
