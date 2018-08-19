package db

import (
	"testing"

	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
)

const (
	accountName  = "fanyang"
	accountChain = "main"
)

func initTest(t *testing.T) {
	InitDB(PostgresCfg{
		Address:  "127.0.0.1:5432",
		User:     "pgfy",
		Password: "123456",
		Database: "test3",
	})

	token.Reg(types.NewAsset("main", 0, eos.EOSSymbol).Symbol)
	token.Reg(types.NewAsset("side", 0, eos.EOSSymbol).Symbol)

	CreateAccount(accountName, accountChain)
}

// test add
func TestAccountToken(t *testing.T) {
	initTest(t)
	for i := 0; i < 4; i++ {
		err := AddToken(accountName, accountChain, types.NewAsset("main", 1111111, eos.EOSSymbol))
		if err != nil {
			t.Errorf("add token err by %s", err.Error())
			t.FailNow()
			return
		}
	}

	for i := 0; i < 4; i++ {
		err := AddToken(accountName, accountChain, types.NewAsset("main", 1111111, eos.EOSSymbol))
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
		err := CostToken(accountName, accountChain, types.NewAsset("main", 11112, eos.EOSSymbol))
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
