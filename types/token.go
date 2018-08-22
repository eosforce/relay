package types

import (
	"fmt"
	"strings"

	eos "github.com/eoscanada/eos-go"
	eosforce "github.com/fanyang1988/eos-go"
)

// Asset is token map in relay
type Asset struct {
	Amount int64
	Symbol
}

// Add add two asset if err panic
func (a Asset) Add(other Asset) Asset {
	if a.Symbol != other.Symbol {
		panic("Add applies only to assets with the same symbol")
	}
	return Asset{Amount: a.Amount + other.Amount, Symbol: a.Symbol}
}

// Sub sub two asset if err panic
func (a Asset) Sub(other Asset) Asset {
	if a.Symbol != other.Symbol {
		panic("Sub applies only to assets with the same symbol")
	}
	return Asset{Amount: a.Amount - other.Amount, Symbol: a.Symbol}
}

// GetSymbol get Symbol string
func (a Asset) GetSymbol() string {
	return a.Symbol.Symbol
}

func (a Asset) String() string {
	strInt := fmt.Sprintf("%d", a.Amount)
	if len(strInt) < int(a.Symbol.Precision+1) {
		// prepend `0` for the difference:
		strInt = strings.Repeat("0", int(a.Symbol.Precision+uint8(1))-len(strInt)) + strInt
	}

	var result string
	if a.Symbol.Precision == 0 {
		result = strInt
	} else {
		result = strInt[:len(strInt)-int(a.Symbol.Precision)] + "." + strInt[len(strInt)-int(a.Symbol.Precision):]
	}

	return fmt.Sprintf("%s:%s %s", a.Chain, result, a.Symbol.Symbol)
}

// NewAsset create asset
func NewAsset(a int64, symbol Symbol) Asset {
	return Asset{
		Amount: a,
		Symbol: symbol,
	}
}

// Symbol is token's symbol mapped in relay
type Symbol struct {
	Precision uint8
	Symbol    string
	Chain     ChainName `json:"chain"`
}

// FromEosSymbol gain symbol from eos
func FromEosSymbol(chainName string, symbol eos.Symbol) Symbol {
	return Symbol{
		Precision: symbol.Precision,
		Symbol:    symbol.Symbol,
		Chain:     ChainName(chainName),
	}
}

// FromEosforceSymbol gain symbol from eosforce
func FromEosforceSymbol(chainName string, symbol eosforce.Symbol) Symbol {
	return Symbol{
		Precision: symbol.Precision,
		Symbol:    symbol.Symbol,
		Chain:     ChainName(chainName),
	}
}

func (s Symbol) String() string {
	return fmt.Sprintf("%s:%s", s.Chain, s.Symbol)
}
