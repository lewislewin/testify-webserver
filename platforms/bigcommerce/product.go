package bigcommerce

type Product struct {
	ID            string
	Name          string
	Price         string
	StockQuantity int
}

func (bc *BigCommerce) GetProducts() (interface{}, error) {
	var products []Product
	// Example logic (make actual API call to BigCommerce)
	return products, nil
}
