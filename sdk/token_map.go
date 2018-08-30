package sdk

import (
	"strings"

	"github.com/eosforce/relay/chain/base/client"
	"github.com/eosforce/relay/const"
	"github.com/eosforce/relay/types"
)

// TokenMapInRelay
//  `Token in` is map token from a its chain to relay name space.
//  In a chain use transfer to trans token to the account define in token map,
//  this will make relay give the account in relay same token map in relay.
func TokenMapInRelay(account, chain string, asset types.Asset) error {
	cli := client.Get(chain)
	asset.Amount += 2000 // fee
	return cli.Transfer(account, consts.TokenInAcc, asset, "token map to relay")
}

// TokenMapOutRelay
// `r.token.out` is fee account in main chain, after test that may be changed.
// memo like `acc:chain:100.0000 EOS` format is:
func TokenMapOutRelay(account, chain string, toAccount string, asset types.Asset) error {
	cli := client.Get(chain)

	memo := strings.Join([]string{
		toAccount,
		string(asset.Chain),
		asset.StringInChain()}, ":")

	return cli.Transfer(account, consts.TokenOutAcc, types.NewAsset(1000, types.Symbol{
		Precision: 4,
		Symbol:    "EOS",
		Chain:     types.ChainName(chain),
	}), memo)
}
