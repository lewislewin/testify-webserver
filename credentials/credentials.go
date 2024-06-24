package credentials

func GetCredentialById(id string) (map[string]string, error) {
	// Just using default creds for now, should retrieve from some kind of secrets manager in future
	credentials := map[string]string{
		"access_token": "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",
		"store_domain": "d46136-ae.myshopify.com",
	}

	return credentials, nil
}
