package db

import (
	"testing"

	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
)

func initDBs() {
	InitDB(PostgresCfg{
		Address:  "127.0.0.1:5432",
		User:     "pgfy",
		Password: "123456",
		Database: "test3",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "main",
		Symbol:    "EOS",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "side",
		Symbol:    "EOS",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "side",
		Symbol:    "SYS",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "main",
		Symbol:    "TST",
	})

}

func TestCreateAccount(t *testing.T) {
	initDBs()
	err := CreateAccount("fanyang", "main")
	if err != nil {
		t.Errorf("err by create %s", err.Error())
	}

	acc, err := GetAccount("fanyang", "main")
	if err != nil {
		t.Errorf("err by get %s", err.Error())
	}
	t.Logf("get account %v", acc)
}

func TestCreateAccountPermissions(t *testing.T) {
	initDBs()

	pers, err := GetAccountPermission("fy", "main")
	if err != nil {
		t.Errorf("err by get %s", err.Error())
	}
	t.Logf("get account %v", pers)

	accs, err := QueryAccountByPermission("owner", "EOS66ncp9xAzdCaAggUDChFMq9YX9wfAJQmN5uMdxf8APV4ecqaZV")
	if err != nil {
		t.Errorf("err by get %s", err.Error())
	}
	t.Logf("get account %v", accs)
}
