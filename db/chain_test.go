package db

import "testing"

func TestGetChains(t *testing.T) {
	initDBs()

	data, err := GetChain("main")
	if err != nil {
		t.Errorf("err by get %s", err.Error())
	}
	t.Logf("get chain %v", *data)

	pers, err := GetChains()
	if err != nil {
		t.Errorf("err by get %s", err.Error())
	}
	t.Logf("get chain %v", pers)

}
