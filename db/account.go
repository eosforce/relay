package db

import "fmt"

// Account account in db
type Account struct {
	Name  string `json:"name"`
	Chain string `json:"chain"`
}

func (u *Account) String() string {
	return fmt.Sprintf("%s::%s", u.Chain, u.Name)
}

// AccountPermission account's permission
type AccountPermission struct {
	Name       string `json:"name"`
	Chain      string `json:"chain"`
	Permission string `json:"permission"`
	Pubkey     string `json:"pubkey"`
}

func (u *AccountPermission) String() string {
	return fmt.Sprintf("%s::%s %s %s",
		u.Chain, u.Name, u.Permission, u.Pubkey)
}
