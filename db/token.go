package db

import (
	"fmt"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
	pg "gopkg.in/pg.v4"
)

// AccountToken Account's Token in db
type AccountToken struct {
	Name       string `json:"name"`
	Chain      string `json:"chain"`
	TokenChain string `json:"token_chain"`
	Symbol     string `json:"symbol"`
	Amount     int64  `json:"amount"`
}

func (u *AccountToken) String() string {
	return fmt.Sprintf("%s::%s %s::%s",
		u.Chain, u.Name, u.TokenChain, eos.Asset{
			Amount: u.Amount,
			Symbol: eos.Symbol{
				Symbol: u.Symbol,
			}}.String())
}

// GetAccountTokens get account tokens, not REALTIME!!!
func GetAccountTokens(name, chain string) ([]types.Asset, error) {
	var res []AccountToken
	err := Get().Model(&res).Where("name=?,chain=?", name, chain).Select()

	if err != nil {
		return nil, err
	}

	r := make([]types.Asset, 0, len(res))
	for _, rr := range res {
		typ, err := token.GetSymbol(rr.TokenChain, rr.Symbol)
		if err != nil {
			seelog.Warnf("get token  err by %")
		}
		r = append(r, types.NewAsset(
			rr.Chain,
			rr.Amount,
			typ.Symbol,
		))
	}

	return r, nil
}

// GetAccountToken get account tokens, not REALTIME!!!
func GetAccountToken(db *pg.DB, name, chain, tokenChain, symbol string) (types.Asset, error) {
	typ, err := token.GetSymbol(chain, symbol)
	if err != nil {
		return types.Asset{}, err
	}

	var rs []AccountToken
	err = Get().Model(&rs).
		Where("name=?,chain=?,token_chain=?,symbol=?",
			name, chain, tokenChain, symbol).
		Select()

	if err != nil || len(rs) == 0 {
		return types.NewAsset(
			tokenChain,
			0,
			typ.Symbol,
		), err
	}

	r := rs[0]

	return types.NewAsset(
		r.Chain,
		r.Amount,
		typ.Symbol,
	), nil
}

// AddToken add token to account
func AddToken(name, chain string, asset types.Asset) error {
	return nil
}

// CostToken cost token from account if no token return a error
func CostToken(name, chain string, asset types.Asset) error {
	return nil
}
