package db

import "testing"

func TestCreateAccount(t *testing.T) {
	InitDB(PostgresCfg{
		Address:  "127.0.0.1:5432",
		User:     "pgfy",
		Password: "123456",
		Database: "test3",
	})

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
