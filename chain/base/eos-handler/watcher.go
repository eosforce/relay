package eosHandler

import (
	"sync"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/eos-go"
)

// ActionData action data with block and transaction id
type ActionData struct {
	Action   eos.Action
	BlockNum uint32
	BlockID  eos.SHA256Bytes
	TrxID    eos.SHA256Bytes
}

// ActionHandlerFunc handler func type
type ActionHandlerFunc func(action ActionData)

// EosWatcher watch eos actions by some p2pAdds
type EosWatcher struct {
	name    string
	p2pAdds []string
	apiURL  string

	peers map[string]*P2PPeer

	blockChan chan eos.SignedBlock
	errChan   chan ErrP2PPeer
	stopChan  chan interface{}

	handlers          []ActionHandlerFunc
	processedBlockNum uint32

	waitter sync.WaitGroup
}

// TODO BY FanYang support change p2p address in running

// NewEosWatcher create a watcher by some p2pAdds to a eos chain
func NewEosWatcher(name, apiURL string, p2pAdds []string) *EosWatcher {
	return &EosWatcher{
		name:    name,
		apiURL:  apiURL,
		p2pAdds: p2pAdds,

		peers:     make(map[string]*P2PPeer, 64),
		blockChan: make(chan eos.SignedBlock, len(p2pAdds)*64+32),
		errChan:   make(chan ErrP2PPeer),
		stopChan:  make(chan interface{}),

		handlers: make([]ActionHandlerFunc, 0, 32),
	}
}

// Start start watching
func (w *EosWatcher) Start() error {
	api := eos.New(w.apiURL)
	info, err := api.GetInfo()
	if err != nil {
		return seelog.Errorf("get chain %s info err by %s", w.name, w.apiURL)
	}

	for _, add := range w.p2pAdds {
		w.waitter.Add(1)
		peer := NewP2PPeer(w.blockChan, w.errChan, add, info.ChainID, 0)
		peer.Connect(info.HeadBlockNum, info.HeadBlockID, info.HeadBlockTime.Time,
			info.LastIrreversibleBlockNum, info.LastIrreversibleBlockID)
		w.peers[add] = peer
	}

	w.waitter.Add(1)
	go func() {
		defer w.waitter.Done()
		for {
			isStop, _ := w.loop()
			if isStop {
				seelog.Warnf("watcher stop by close")
				return
			}
		}
	}()

	return nil
}

// Stop stop watching then close all
func (w *EosWatcher) Stop() {
	// TODO By FanYang imp Client close
	close(w.stopChan)
}

// Wait wait watching stop
func (w *EosWatcher) Wait() {
	w.waitter.Wait()
}

// RegHandler reg a action handler, this can only call before start watching
func (w *EosWatcher) RegHandler(h ActionHandlerFunc) {
	w.handlers = append(w.handlers, h)
}

// loop main loop for process block from per p2p peer
func (w *EosWatcher) loop() (bool, error) {
	select {
	case <-w.stopChan:
		w.closeAll()
		return true, nil
	case b := <-w.blockChan:
		return false, w.processBlock(&b)
	case err := <-w.errChan:
		{
			if err.Peer == nil {
				return false, seelog.Errorf("peer err nil")
			}
			// reconnect
		}
	}
	return false, nil
}

func (w *EosWatcher) processBlock(block *eos.SignedBlock) error {
	// TODO By FanYang next version will to process fork from diff peer
	blockNum := block.BlockNumber()
	if w.processedBlockNum >= blockNum {
		seelog.Debugf("no process processed block %d", blockNum)
		return nil
	}

	seelog.Infof("process block %d", blockNum)
	w.processedBlockNum = blockNum

	// TODO By FanYang just process block util it be IrreversibleBlock
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
			w.handler(ActionData{
				BlockID:  blockID,
				BlockNum: blockNum,
				TrxID:    tr.Transaction.ID,
				Action:   *action,
			})
		}
	}

	return nil
}

func (w *EosWatcher) handler(act ActionData) {
	for _, h := range w.handlers {
		h(act)
	}
}

func (w *EosWatcher) closeAll() {
	seelog.Errorf("Stop Watcher Last BlockNum is %d", w.processedBlockNum)
	seelog.Flush()

	// TODO wait all connect closed
	for addr, peer := range w.peers {
		seelog.Warnf("stop connect %s", addr)
		peer.Close()
		seelog.Warnf("stop connect %s ok", addr)
		w.waitter.Done()
	}
}
