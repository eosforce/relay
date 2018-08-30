package sdk

import (
	"strings"

	"github.com/eosforce/relay/const"

	"github.com/eosforce/relay/chain/base/client"
	"github.com/eosforce/relay/types"
)

// ExchangeTest a exchange for test
// {fromAsset:fromChain:ToSymbol:ToChain}
func ExchangeTest(account, chain string, from types.Asset, to types.Symbol) error {
	cli := client.Get(chain)

	memo := strings.Join([]string{
		from.StringInChain(),
		string(from.Chain),
		to.Symbol,
		string(to.Chain)}, ":")

	return cli.Transfer(account, consts.TestExchange, types.NewAsset(1000, types.Symbol{
		Precision: 4,
		Symbol:    "EOS",
		Chain:     types.ChainName(chain),
	}), memo)
}
