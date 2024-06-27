package shopify

type Order struct {
	ID              int            `json:"id,omitempty"`
	LineItems       []LineItem     `json:"line_items,omitempty"`
	Email           string         `json:"email,omitempty"`
	Phone           string         `json:"phone,omitempty"`
	BillingAddress  Address        `json:"billing_address,omitempty"`
	ShippingAddress Address        `json:"shipping_address,omitempty"`
	Transactions    []Transaction  `json:"transactions,omitempty"`
	FinancialStatus string         `json:"financial_status,omitempty"`
	DiscountCodes   []DiscountCode `json:"discount_codes,omitempty"`
}

type LineItem struct {
	VariantID int     `json:"variant_id,omitempty"`
	Title     string  `json:"title,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

type Address struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Address1  string `json:"address1,omitempty"`
	Phone     string `json:"phone,omitempty"`
	City      string `json:"city,omitempty"`
	Province  string `json:"province,omitempty"`
	Country   string `json:"country,omitempty"`
	Zip       string `json:"zip,omitempty"`
}

type Transaction struct {
	Kind   string  `json:"kind,omitempty"`
	Status string  `json:"status,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type DiscountCode struct {
	Code   string `json:"code,omitempty"`
	Amount string `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
}
