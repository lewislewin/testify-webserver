package shopify

type Order struct {
	ID              int
	LineItems       []Product
	BillingAddress  Address
	ShippingAddress Address
}
