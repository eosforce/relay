package orderBook

import (
	"testing"

	"time"

	"math/rand"

	"github.com/cihub/seelog"
	eos "github.com/eoscanada/eos-go"
	"github.com/eosforce/relay/types"
	eosforce "github.com/fanyang1988/eos-go"
)

func TestOrderBook(t *testing.T) {
	defer seelog.Flush()

	pair := types.ExchangePair{
		TokenA: types.FromEosforceSymbol("main", eosforce.EOSSymbol),
		TokenB: types.FromEosSymbol("side", eos.EOSSymbol),
	}
	book := NewOrderBook(pair, func(book *OrderBook, buy, sell *Order) error {
		seelog.Infof("ok %s --> %s", buy, sell)
		return nil
	})

	buyer := types.NewAccount("fy01", "main")
	seller := types.NewAccount("fy02", "side")
	book.Start()

	go func() {
		asset := types.NewAsset(10000000, pair.TokenB)
		for {
			asset.Amount++
			book.Buy(buyer, asset, 1000+rand.Int63n(1000))
			time.Sleep(1 * time.Microsecond)
		}
	}()

	go func() {
		asset := types.NewAsset(10000000, pair.TokenA)
		for {
			asset.Amount++
			book.Sell(seller, asset, 1000+rand.Int63n(1000))
			time.Sleep(1 * time.Microsecond)
		}
	}()

	time.Sleep(3 * time.Second)
	book.Stop()
	book.Wait()

	//will panic by send cmd to chan closed
}
