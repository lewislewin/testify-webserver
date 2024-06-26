package shopify

type Product struct {
	ID        string
	Title     string
	Price     string
	Inventory int
}

func (s *Shopify) GetProducts() (interface{}, error) {
	var products []Product
	// Example logic (make actual API call to Shopify)
	products = append(products, Product{
		ID:        "1",
		Title:     "Hat",
		Price:     "100",
		Inventory: 1,
	})
	return products, nil
}
