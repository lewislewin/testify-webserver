package shopify

import "fmt"

type Order struct {
	ID        string
	Title     string
	Price     string
	Inventory int
}

func (s *Shopify) CreateOrder(order interface{}) error {
	fmt.Println(order)
	return nil
}
