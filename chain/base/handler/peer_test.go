package handler

import (
	"testing"

	"github.com/eosforce/relay/types"

	"github.com/cihub/seelog"
	"github.com/eoscanada/eos-go"
)

// TestPeerConnectToEOS test if can connect a local chain
func TestPeerConnectToEOS(t *testing.T) {
	defer seelog.Flush()

	blockChan := make(chan types.SignedBlockInterface, 64)
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
	peer := NewP2PPeer(blockChan, errChan, p2pAddr, types.SHA256BytesFromEos(info.ChainID), 1)

	peer.handler = eosHandler{}

	peer.Connect(info.HeadBlockNum, types.SHA256BytesFromEos(info.HeadBlockID), info.HeadBlockTime.Time,
		info.LastIrreversibleBlockNum, types.SHA256BytesFromEos(info.LastIrreversibleBlockID))

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
