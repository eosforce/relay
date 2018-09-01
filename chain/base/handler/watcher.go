package handler

import (
	"fmt"
	"sync"

	"github.com/eosforce/relay/const"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/types"
)

// ActionHandlerFunc handler func type
type ActionHandlerFunc func(action types.ActionData)

// EosWatcher watch eos actions by some p2pAdds
type EosWatcher struct {
	name    string
	p2pAdds []string
	apiURL  string

	peers map[string]*P2PPeer

	blockChan chan types.SignedBlockInterface
	errChan   chan ErrP2PPeer
	stopChan  chan interface{}

	handlers          []ActionHandlerFunc
	handler           handlerInterface
	processedBlockNum uint32

	chainTyp int

	waitter sync.WaitGroup
}

// TODO BY FanYang support change p2p address in running

// NewEosWatcher create a watcher by some p2pAdds to a eos chain
func NewEosWatcher(chainTyp int, name, apiURL string, p2pAdds []string) *EosWatcher {
	res := &EosWatcher{
		chainTyp: chainTyp,
		name:     name,
		apiURL:   apiURL,
		p2pAdds:  p2pAdds,

		peers:     make(map[string]*P2PPeer, 64),
		blockChan: make(chan types.SignedBlockInterface, len(p2pAdds)*64+32),
		errChan:   make(chan ErrP2PPeer),
		stopChan:  make(chan interface{}),

		handlers: make([]ActionHandlerFunc, 0, 32),
	}

	switch chainTyp {
	case consts.TypeBaseEos:
		res.handler = eosHandler{}
	case consts.TypeBaseEosforce:
		res.handler = eosforceHandler{}
	default:
		panic(fmt.Errorf("chain %s typ %s err", name, chainTyp))
	}

	return res
}

func (w *EosWatcher) Name() string {
	return w.name
}

// Start start watching
func (w *EosWatcher) Start() error {
	err := w.handler.Start(w)
	if err != nil {
		return err
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
		return false, w.processBlock(b)
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

func (w *EosWatcher) processBlock(block types.SignedBlockInterface) error {
	blockNum := block.BlockNumber()
	if w.processedBlockNum >= blockNum {
		return nil
	}
	w.processedBlockNum = blockNum
	seelog.Infof("process block %d", blockNum)

	// TODO By FanYang just process block util it be IrreversibleBlock
	// TODO By FanYang next version will to process fork from diff peer

	return w.handler.ProcessBlock(w, block)
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
