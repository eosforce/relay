package eosforceHandler

import (
	"testing"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/eos-go"
)

// TestPeerConnectToEOS test if can connect a local chain
func TestPeerConnectToEOS(t *testing.T) {
	defer seelog.Flush()

	blockChan := make(chan eos.SignedBlock, 64)
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
	peer := NewP2PPeer(blockChan, errChan, p2pAddr, info.ChainID, 1)

	peer.Connect(info.HeadBlockNum, info.HeadBlockID, info.HeadBlockTime.Time,
		info.LastIrreversibleBlockNum, info.LastIrreversibleBlockID)

	got := 0

	for {
		select {
		case b := <-blockChan:
			{
				t.Logf("get act %v", b.BlockNumber())
				got++
				if got > 5 {
					t.Logf("get 5 block ok")
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
