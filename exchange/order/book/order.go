package orderBook

import (
	"time"

	"fmt"

	"github.com/eosforce/relay/types"
)

const (
	// OrderSell sell order type
	OrderSell = iota + 1
	// OrderBuy buy order type
	OrderBuy
)

// Order a order to buy tokenA by tokenB with tokenB * price, or sell tokenA to get tokenB with tokenB * price
type Order struct {
	Account    types.Account      `json:"account"`
	UpdateTime time.Time          `json:"update_time"`
	Pair       types.ExchangePair `json:"pair"`
	Type       int                `json:"type"` // sell or buy
	Asset      types.Asset        `json:"asset"`
	Price      int64              `json:"price"`
}

// NewBuy make a buy order
func NewBuy(account types.Account, pair types.ExchangePair, asset types.Asset, price int64) Order {
	return Order{
		Account:    account,
		Type:       OrderBuy,
		Pair:       pair,
		Asset:      asset,
		Price:      price,
		UpdateTime: time.Now(),
	}
}

// NewSell make a sell order
func NewSell(account types.Account, pair types.ExchangePair, asset types.Asset, price int64) Order {
	return Order{
		Account:    account,
		Type:       OrderSell,
		Pair:       pair,
		Asset:      asset,
		Price:      price,
		UpdateTime: time.Now(),
	}
}

// String like `chain:name buy 1000.0000 main:EOS From side:SYS by 10.12`
func (o Order) String() string {
	typStr := "buy"
	forSymbol := o.Pair.TokenB

	if o.Type == OrderSell {
		typStr = "sell"
		forSymbol = o.Pair.TokenA
	}

	return fmt.Sprintf(
		"%s %s %s From %s By %d",
		o.Account, typStr, o.Asset, forSymbol, o.Price,
	)
}
