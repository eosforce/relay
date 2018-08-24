package types

import "fmt"

// Account is account in relay, it map from the chains
type Account struct {
	Chain string `json:"chain"`
	Name  string `json:"name"`
}

func (a Account) String() string {
	return fmt.Sprintf("%s:%s", a.Chain, a.Name)
}

// NewAccount new account
func NewAccount(name, chain string) Account {
	return Account{
		Name:  name,
		Chain: chain,
	}
}

// Permission permission in a account , like owner or active
type Permission struct {
	Permission string `json:"permission"`
	PublicKey  string `json:"pubkey"`
}
