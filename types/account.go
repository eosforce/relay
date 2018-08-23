package types

// Account is account in relay, it map from the chains
type Account struct {
	Chain string `json:"chain"`
	Name  string `json:"name"`
}

// Permission permission in a account , like owner or active
type Permission struct {
	Permission string `json:"permission"`
	PublicKey  string `json:"pubkey"`
}
