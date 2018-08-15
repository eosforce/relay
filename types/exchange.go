package types

// ExchangeName exchange pair name
type ExchangeName string

// ExchangePair exchange pair between 2 token, is has different exchange types, now is just bancor
type ExchangePair struct {
	Name   ExchangeName `json:"name"`
	Type   string       `json:"typ"`
	TokenA Symbol       `json:"a"`
	TokenB Symbol       `json:"b"`
}

// ExchangeHistory account exchange history
type ExchangeHistory struct {
	Exchange  ExchangeName `json:"name"`
	From      Account      `json:"from_account"`
	To        Account      `json:"to_account"`
	FromToken Asset        `json:"from_token"`
	ToToken   Asset        `json:"to_token"`
}
