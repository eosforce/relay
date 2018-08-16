package db

import (
	"fmt"

	"github.com/eoscanada/eos-go"
)

// AccountToken Account's Token in db
type AccountToken struct {
	Name       string `json:"name"`
	Chain      string `json:"chain"`
	TokenChain string `json:"token_chain"`
	Precision  uint8  `json:"precision"`
	Symbol     string `json:"symbol"`
	Amount     int64  `json:"amount"`
}

func (u *AccountToken) String() string {
	return fmt.Sprintf("%s::%s %s::%s",
		u.Chain, u.Name, u.TokenChain, eos.Asset{
			Amount: u.Amount,
			Symbol: eos.Symbol{
				Precision: u.Precision,
				Symbol:    u.Symbol,
			}}.String())
}
