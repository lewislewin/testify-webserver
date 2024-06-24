package shopify

const SHOPIFY_URL_SUFFIX = ".myshopify.com"

type Shopify struct {
	StoreStub string
	Token     string
}

func (s *Shopify) Authenticate(credentials map[string]string) error {
	s.StoreStub = credentials["store_stub"]
	s.Token = credentials["token"]

	//make sure connection is valid and works

	return nil
}
