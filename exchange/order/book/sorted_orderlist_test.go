package orderBook

import (
	"testing"

	"github.com/eosforce/relay/types"

	"github.com/cihub/seelog"
	eos "github.com/eoscanada/eos-go"
	eosforce "github.com/fanyang1988/eos-go"
)

// TestSortedOrderList_Sort sort
func TestSortedOrderList_Sort(t *testing.T) {
	defer seelog.Flush()
	account := types.Account{
		Name:  "ff",
		Chain: "main",
	}

	pair := types.ExchangePair{
		TokenA: types.FromEosforceSymbol("main", eosforce.EOSSymbol),
		TokenB: types.FromEosSymbol("side", eos.EOSSymbol),
	}

	sol := NewSortedOrderList(false, 1024)
	sol.Append(NewBuy(
		account, pair,
		types.NewAssetFromEosforce("main", 11, eosforce.EOSSymbol), 11))
	sol.Append(NewBuy(
		account, pair,
		types.NewAssetFromEosforce("main", 22, eosforce.EOSSymbol), 2))
	sol.Append(NewBuy(
		account, pair,
		types.NewAssetFromEosforce("main", 33, eosforce.EOSSymbol), 3))
	sol.Append(NewBuy(
		account, pair,
		types.NewAssetFromEosforce("main", 44, eosforce.EOSSymbol), 4))
	sol.Append(NewBuy(
		account, pair,
		types.NewAssetFromEosforce("main", 66, eosforce.EOSSymbol), 5))
	sol.Append(NewBuy(
		account, pair,
		types.NewAssetFromEosforce("main", 55, eosforce.EOSSymbol), 5))

	sol.Sort()

	for i, s := range sol.orderList {
		seelog.Infof("%d %s", i, s)
	}
}
