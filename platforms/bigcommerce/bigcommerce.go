package bigcommerce

type BigCommerce struct {
	ClientID    string
	AccessToken string
	StoreHash   string
}

func (bc *BigCommerce) Authenticate(credentials map[string]string) error {
	bc.ClientID = credentials["client_id"]
	bc.AccessToken = credentials["access_token"]
	bc.StoreHash = credentials["store_hash"]
	return nil
}
