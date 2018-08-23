package db

import (
	"fmt"

	"time"

	"github.com/eosforce/relay/types"
	"github.com/pkg/errors"
)

// Account account in db
type Account struct {
	Name       string    `json:"name"`
	Chain      string    `json:"chain"`
	CreateTime time.Time `json:"create_time"`
}

func (u *Account) String() string {
	return fmt.Sprintf("%s::%s", u.Chain, u.Name)
}

// GetAccount get account from db
func GetAccount(name, chain string) (*Account, error) {
	db := Get()

	var res []Account
	err := db.Model(&res).
		Where("name=? and chain = ?", name, chain).
		Select()

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}

// CreateAccount Account to db TODO add pubkey
func CreateAccount(name, chain string) error {
	return errors.WithMessage(
		Get().Create(&Account{
			Name:  name,
			Chain: chain,
		}),
		"create account err By")
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

// GetAccountPermission get account permissions from db
func GetAccountPermission(name, chain string) ([]types.Permission, error) {
	db := Get()

	var res []AccountPermission
	err := db.Model(&res).
		Where("name=? and chain = ?", name, chain).
		Select()

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	resData := make([]types.Permission, 0, len(res))
	for _, k := range res {
		resData = append(resData, types.Permission{
			Permission: k.Permission,
			PublicKey:  k.Pubkey,
		})
	}

	return resData[:], nil
}

// QueryAccountByPermission query account by permission
func QueryAccountByPermission(permission, pubkey string) ([]AccountPermission, error) {
	db := Get()

	var res []AccountPermission
	err := db.Model(&res).
		Where("permission=? and pubkey=?", permission, pubkey).
		Select()

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res[:], nil
}
