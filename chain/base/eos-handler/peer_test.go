package eosHandler

import (
	"testing"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/eos-go"
)

// TestPeerConnectToEOS test if can connect a local chain
func TestPeerConnectToEOS(t *testing.T) {
	defer seelog.Flush()

	actionChan := make(chan eos.Action, 64)
	errChan := make(chan ErrP2PPeer)

	// start a chain http api :8889, p2p address 9001
	apiURL := "http://127.0.0.1:8890"
	p2pAddr := "127.0.0.1:9002"

	api := eos.New(apiURL)
	info, err := api.GetInfo()
	if err != nil {
		t.Error("Error getting info: ", err)
		t.FailNow()
		return
	}
	seelog.Infof("get info %v %v", info.LastIrreversibleBlockNum, info.LastIrreversibleBlockID)
	peer := NewP2PPeer(actionChan, errChan, p2pAddr, info.ChainID, 1)

	peer.Connect(info.HeadBlockNum, info.HeadBlockID, info.HeadBlockTime.Time,
		info.LastIrreversibleBlockNum, info.LastIrreversibleBlockID)

	actionGetted := 0

	for {
		select {
		case act := <-actionChan:
			{
				t.Logf("get act %v %v", act.Name, act.Account)
				actionGetted++
				if actionGetted > 32 {
					t.Logf("get 32 action ok")
					return
				}
			}
		case err := <-errChan:
			{
				t.Errorf("err %v", err.Err.Error())
				t.FailNow()
				return
			}
		}
	}
}
