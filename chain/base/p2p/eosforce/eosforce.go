package eosforce

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base"
	"github.com/eosforce/relay/chain/base/p2p/types"
	"github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/p2p"
)

type Message struct {
	*p2p.Message
}

func (m Message) GetP2PMsg() interface{} {
	return m.Envelope.P2PMessage
}

func (m Message) GetType() byte {
	return byte(m.Envelope.Type)
}

type P2PClient struct {
	*p2p.Client
}

func (p P2PClient) RegHandler(h types.MessageHandler) {
	p.RegisterHandler(p2p.HandlerFunc(func(msg p2p.Message) {
		h(Message{
			&msg,
		})
	}))
}

func NewClient(p2pAddr string, chainID base.SHA256Bytes) P2PClient {
	c := p2p.NewClient(p2pAddr, eos.SHA256Bytes(chainID), 1)
	c.WithLogger(seelog.Current)
	return P2PClient{c}
}
