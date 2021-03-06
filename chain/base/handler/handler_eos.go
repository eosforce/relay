package handler

import (
	"fmt"

	"github.com/cihub/seelog"
	eos "github.com/eoscanada/eos-go"
	"github.com/eosforce/relay/chain/base/p2p/types"
	"github.com/eosforce/relay/types"
)

type handlerInterface interface {
	Start(w *EosWatcher) error
	PeerHandler(blockChan chan<- types.SignedBlockInterface, msg p2pTypes.P2PMessage)
	ProcessBlock(w *EosWatcher, sb types.SignedBlockInterface) error
}

type eosHandler struct {
}

func (e eosHandler) Start(w *EosWatcher) error {
	api := eos.New(w.apiURL)
	info, err := api.GetInfo()
	if err != nil {
		return seelog.Errorf("get chain %s %s info err by %s",
			w.name, w.apiURL, err.Error())
	}

	for _, add := range w.p2pAdds {
		w.waitter.Add(1)
		peer := NewP2PPeer(w.blockChan, w.errChan, add, types.SHA256BytesFromEos(info.ChainID), 1)
		peer.chainTyp = w.chainTyp
		peer.handler = w.handler
		peer.Connect(info.HeadBlockNum, types.SHA256BytesFromEos(info.HeadBlockID), info.HeadBlockTime.Time,
			info.LastIrreversibleBlockNum, types.SHA256BytesFromEos(info.LastIrreversibleBlockID))
		w.peers[add] = peer
	}

	return nil
}

func (e eosHandler) PeerHandler(blockChan chan<- types.SignedBlockInterface, msg p2pTypes.P2PMessage) {
	switch eos.P2PMessageType(msg.GetType()) {
	case eos.SignedBlockType:
		{
			signedBlockMsg, ok := msg.GetP2PMsg().(*eos.SignedBlock)
			if !ok {
				seelog.Error("typ error by signedBlockMsg")
				return
			}

			blockChan <- signedBlockMsg
			return
		}
	}
}

func (e eosHandler) ProcessBlock(w *EosWatcher, sb types.SignedBlockInterface) error {
	block, ok := sb.(*eos.SignedBlock)
	if !ok {
		return fmt.Errorf("block type err %d", sb.BlockNumber())
	}

	blockNum := block.BlockNumber()
	blockID, err := block.BlockID()
	if err != nil {
		return seelog.Errorf("get block ID err by %s", err.Error())
	}

	for _, tr := range block.Transactions {
		trx, err := tr.Transaction.Packed.Unpack()
		if err != nil {
			seelog.Errorf("transaction unpack err by %s", err.Error())
			continue
		}

		for _, action := range trx.Actions {
			for _, h := range w.handlers {
				h(types.ActionDataFromEos(blockID, blockNum, tr.Transaction.ID, action))
			}
		}
	}

	return nil
}
