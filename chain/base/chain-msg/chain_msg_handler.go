package chainMsg

import (
	"sync"

	"github.com/cihub/seelog"
)

// HandleFunc msg handle func
type HandleFunc func(msg *ChainMsg)

// Handler manager handler functions to process msg
type Handler struct {
	msgChan  chan ChainMsg
	stopChan chan interface{}
	handlers map[string][]HandleFunc
	wg       sync.WaitGroup
}

// NewChainMsgHandler create a handler need set handler func before start
func NewChainMsgHandler() *Handler {
	return &Handler{
		msgChan:  make(chan ChainMsg, 40960),
		handlers: make(map[string][]HandleFunc, 32),
		stopChan: make(chan interface{}),
	}
}

// Start start to process msg
func (c *Handler) Start() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.stopChan:
				seelog.Warnf("ChainMsg Handler close")
				// TODO By FanYang process all msg before stop
				return
			case msg := <-c.msgChan:
				hs, ok := c.handlers[msg.MsgTyp]
				if ok {
					for _, h := range hs {
						h(&msg)
					}
				}
			}
		}
	}()
}

// Stop stop process, need use wait to wait gorountinue exit
func (c *Handler) Stop() {
	close(c.stopChan)
}

// Wait wait gorountinue exit
func (c *Handler) Wait() {
	c.wg.Wait()
}

// AddHandler add a handler to a class typ msg, need call before start
func (c *Handler) AddHandler(typ string, f HandleFunc) {
	hs, ok := c.handlers[typ]
	if ok {
		hs = append(hs, f)
		c.handlers[typ] = hs
	} else {
		hs := make([]HandleFunc, 0, 32)
		hs = append(hs, f)
		c.handlers[typ] = hs[:]
	}
}

// PushMsg push msg to handler
func (c *Handler) PushMsg(msg *ChainMsg) {
	c.msgChan <- *msg
}
