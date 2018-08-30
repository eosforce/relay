package sdk

import (
	"github.com/eosforce/relay/chain/base/client"
	"github.com/eosforce/relay/const"
	"github.com/eosforce/relay/types"
)

// MapAccountToRelay map account to relay
//  now map account is implement by a transfer to "relay.a.map", the token transfer is the fee to map.
//  For account in main chain, use eosio.transfer to tranfer 0.1 EOSC to "relay.a.map" in main chain.
func MapAccountToRelay(account, chain string) error {
	cli := client.Get(chain)
	return cli.Transfer(account, consts.AccountMapAcc,
		types.NewAsset(1000, types.Symbol{
			Precision: 4,
			Symbol:    "EOS",
			Chain:     types.ChainName(chain),
		}), "account map to relay")
}
