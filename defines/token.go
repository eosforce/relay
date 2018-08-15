package defines

import (
	"fmt"
	"strings"

	"github.com/fanyang1988/eos-go"
)

// Asset is token map in relay
type Asset struct {
	Amount int64
	Symbol
}

func (a Asset) Add(other Asset) Asset {
	if a.Symbol != other.Symbol {
		panic("Add applies only to assets with the same symbol")
	}
	return Asset{Amount: a.Amount + other.Amount, Symbol: a.Symbol}
}

func (a Asset) Sub(other Asset) Asset {
	if a.Symbol != other.Symbol {
		panic("Sub applies only to assets with the same symbol")
	}
	return Asset{Amount: a.Amount - other.Amount, Symbol: a.Symbol}
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

	return fmt.Sprintf("%s:%s %s", a.Chain, result, a.Symbol.Symbol.Symbol)
}

// Symbol is token's symbol mapped in relay
type Symbol struct {
	eos.Symbol
	Chain ChainName `json:"chain"`
}

func (s Symbol) String() string {
	return fmt.Sprintf("%s:%s", s.Chain, s.Symbol.Symbol)
}
