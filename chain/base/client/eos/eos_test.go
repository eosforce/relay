package eosClient

import (
	"testing"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/wallets"
	"github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/token"
)

func TestClient(t *testing.T) {
	defer seelog.Flush()

	wallets.Get().RegWallet(wallets.WalletInfo{
		APIURL:     "http://127.0.0.1:18888",
		URL:        "http://127.0.0.1:16666",
		ChainName:  "side",
		WalletName: "default",
	})

	client := NewClient("side")
	res, err := client.PushEOSCActions(token.NewTransfer(
		eos.AN("eosforce"),
		eos.AN("biosbpa"),
		eos.NewEOSAsset(121),
		"testest",
	))

	if err != nil {
		t.Errorf("send err %s", err.Error())
		t.FailNow()
	}

	seelog.Infof("res %v", res)
}
