package token

import (
	"sync"

	"errors"

	"github.com/eosforce/relay/types"
)

type chainSymbolKey struct {
	chain  string
	symbol string
}

var (
	typs  map[chainSymbolKey]types.Symbol
	mutex sync.RWMutex
)

// Reg reg a symbol
func Reg(symbol types.Symbol) {
	mutex.Lock()
	defer mutex.Unlock()

	typs[chainSymbolKey{
		chain:  string(symbol.Chain),
		symbol: symbol.Symbol.Symbol,
	}] = symbol
}

// GetSymbol get symbol by token chain from and symbol
func GetSymbol(chain, symbol string) (types.Symbol, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	r, ok := typs[chainSymbolKey{
		chain:  string(chain),
		symbol: symbol,
	}]

	if !ok {
		return types.Symbol{}, errors.New("no symbol found")
	}

	return r, nil

}
