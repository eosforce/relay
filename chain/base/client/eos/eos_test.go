package eosClient

import (
	"testing"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/wallets"
	"github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/eosforce"
)

func TestClient(t *testing.T) {
	defer seelog.Flush()

	wallets.Get().RegWallet(wallets.WalletInfo{
		APIURL:     "http://127.0.0.1:8888",
		URL:        "http://127.0.0.1:6666",
		ChainName:  "side",
		WalletName: "default",
	})

	client := NewClient("side")
	res, err := client.PushEOSCActions(eosforce.NewTransfer(
		eosforce.AN("eosforce"),
		eosforce.AN("biosbpa"),
		eos.NewEOSAsset(121),
		"testest",
	))

	if err != nil {
		t.Errorf("send err %s", err.Error())
		t.FailNow()
	}

	seelog.Infof("res %v", res)
}
