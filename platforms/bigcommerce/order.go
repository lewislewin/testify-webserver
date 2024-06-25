package bigcommerce

type Order struct {
	ID        string
	Title     string
	Price     string
	Inventory int
}

func (s *BigCommerce) CreateOrder(order interface{}) error {

	return nil
}
