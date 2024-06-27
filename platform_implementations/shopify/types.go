package shopify

type Order struct {
	LineItems       []LineItem     `json:"line_items"`
	Email           string         `json:"email"`
	Phone           string         `json:"phone"`
	BillingAddress  Address        `json:"billing_address"`
	ShippingAddress Address        `json:"shipping_address"`
	Transactions    []Transaction  `json:"transactions"`
	FinancialStatus string         `json:"financial_status"`
	DiscountCodes   []DiscountCode `json:"discount_codes"`
}

type LineItem struct {
	VariantID int `json:"variant_id"`
	Quantity  int `json:"quantity"`
}

type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address1"`
	Phone     string `json:"phone"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Zip       string `json:"zip"`
}

type Transaction struct {
	Kind   string  `json:"kind"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}
