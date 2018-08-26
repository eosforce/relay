package utils

import (
	"fmt"

	"github.com/eosforce/relay/types"
	"github.com/loocash/go-rationals"
)

// Price to exchange the token from diff chain, we need give every token exchange pair a price
// Warning !!! donot reuse Price as it imp by ref
// the Price 10.00000000, note the exchange between diff token need each token 's min amount
// so the price is imp by a Rational number by To/From, which in a exchange from token * A = to token * B
type Price struct {
	Pair  types.ExchangePair `json:"pair"`
	Value rational.T         `json:"v"`
}

// NewPrice new price from token to to token
func NewPrice(pair types.ExchangePair, from types.Asset, to types.Asset) Price {
	return Price{
		Pair:  pair,
		Value: rational.NewRational(from.Amount, to.Amount),
	}
}

// NewPriceFromInt new price use v from token to get 1 to token
func NewPriceFromInt(pair types.ExchangePair, v int64) Price {
	return Price{
		Pair:  pair,
		Value: rational.NewRational(v, 1),
	}
}

// Equal Return if p == o
func (p *Price) Equal(o *Price) bool {
	return p.Value.Equals(o.Value)
}

// Less Return if p < o
func (p *Price) Less(o *Price) bool {
	// TODO check exchange pair is same
	// TODO overflow check (big num is no need)
	return p.Value.LessThan(o.Value)
}

// More Return if p > o
func (p *Price) More(o *Price) bool {
	return !(p.Equal(o) || p.Less(o))
}

func (p Price) String() string {
	return fmt.Sprintf("%s %s", p.Pair, p.Value)
}

// CalcPrice calc price in asset to by asset from
func CalcPrice(from types.Asset, to types.Asset) int64 {

	return 0
}
