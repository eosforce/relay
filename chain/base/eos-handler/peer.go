package eosHandler

import (
	"time"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/p2p"
)

// ErrP2PPeer error from p2p peer, Peer is point to Err P2PPeer
type ErrP2PPeer struct {
	Err  error
	Peer *P2PPeer
}

// P2PPeer connect to eos p2p node to watch action
type P2PPeer struct {
	blockChan chan<- eos.SignedBlock
	errChan   chan<- ErrP2PPeer
	client    *p2p.Client

	p2pAddress     string
	chainID        eos.SHA256Bytes
	networkVersion uint16
}

// TODO By FanYang change too long params

// NewP2PPeer create connection p2p peer
func NewP2PPeer(blockChan chan<- eos.SignedBlock, errChan chan<- ErrP2PPeer, p2pAddr string, chainID eos.SHA256Bytes, networkVersion uint16) *P2PPeer {
	return &P2PPeer{
		blockChan:      blockChan,
		p2pAddress:     p2pAddr,
		chainID:        chainID,
		networkVersion: networkVersion,
		errChan:        errChan,
	}

}

// Connect connect or reconnect to peer, sync from currHeadBlock
func (p *P2PPeer) Connect(headBlock uint32, headBlockID eos.SHA256Bytes, headBlockTime time.Time, lib uint32, libID eos.SHA256Bytes) {
	p.client = p2p.NewClient(p.p2pAddress, p.chainID, p.networkVersion)
	p.client.WithLogger(seelog.Current)
	p.client.RegisterHandler(p2p.HandlerFunc(p.handler))

	go func() {
		err := p.client.ConnectAndSync(headBlock, headBlockID, headBlockTime, lib, libID)

		p.errChan <- ErrP2PPeer{
			Err:  seelog.Errorf("P2PPeer conn %s err by %s", p.p2pAddress, err.Error()),
			Peer: p,
		}
	}()

}

// Close close connect to peer
func (p *P2PPeer) Close() {
	// TODO By FanYang imp close
}

func (p *P2PPeer) handler(msg p2p.Message) {
	switch msg.Envelope.Type {
	case eos.SignedBlockType:
		{
			signedBlockMsg, ok := msg.Envelope.P2PMessage.(*eos.SignedBlock)
			if !ok {
				seelog.Error("typ error by signedBlockMsg")
				return
			}

			_, err := signedBlockMsg.BlockID()
			if err != nil {
				seelog.Errorf("BlockID error by signedBlockMsg %v", err.Error())
				return
			}

			//seelog.Tracef("on block %d %s %v", signedBlockMsg.BlockNumber(), signedBlockMsg.Producer, blockID)

			p.blockChan <- *signedBlockMsg
			return
		}
	}

	//seelog.Tracef("recv msg from %s --> %s", msg.Route.From, string(data))
}
