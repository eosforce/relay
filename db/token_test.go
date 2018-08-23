package db

import (
	"testing"

	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
)

const (
	accountName  = "fanyang"
	accountChain = "main"
)

// test add
func TestAccountToken(t *testing.T) {
	initDBs()

	for i := 0; i < 4; i++ {
		err := AddToken(accountName, accountChain, types.NewAssetFromEosforce("main", 1111111, eos.EOSSymbol))
		if err != nil {
			t.Errorf("add token err by %s", err.Error())
			t.FailNow()
			return
		}
	}

	for i := 0; i < 4; i++ {
		err := AddToken(accountName, accountChain, types.NewAssetFromEosforce("main", 1111111, eos.EOSSymbol))
		if err != nil {
			t.Errorf("add token err by %s", err.Error())
			t.FailNow()
			return
		}
	}

	res, err := GetAccountTokens(accountName, accountChain)
	if err != nil {
		t.Error("get all token err ", err.Error())
		t.FailNow()
		return
	}

	for i, r := range res {
		t.Log("token ", i, " ", r.String())
	}

	for i := 0; i < 4; i++ {
		err = CostToken(accountName, accountChain, types.NewAssetFromEosforce("main", 11112, eos.EOSSymbol))
		if err != nil {
			t.Errorf("cost token err by %s", err.Error())
			t.FailNow()
			return
		}
	}

	res, err = GetAccountTokens(accountName, accountChain)
	if err != nil {
		t.Error("get all token err ", err.Error())
		t.FailNow()
		return
	}

	for i, r := range res {
		t.Log("token ", i, " ", r.String())
	}
}
