package p2p

import (
	"fmt"

	"github.com/eosforce/relay/chain/base/p2p/eos"
	"github.com/eosforce/relay/chain/base/p2p/eosforce"

	"github.com/eosforce/relay/chain/base"
	"github.com/eosforce/relay/chain/base/p2p/types"
	"github.com/eosforce/relay/const"
)

// Client p2p client interface to eosforce or eos
type Client interface {
	ConnectRecent() error
	RegHandler(h types.MessageHandler)
}

func NewClient(chainTyp int, p2pAddr string, chainID base.SHA256Bytes) Client {
	switch chainTyp {
	case consts.TypeBaseEos:
		return eos.NewClient(p2pAddr, chainID)
	case consts.TypeBaseEosforce:
		return eosforce.NewClient(p2pAddr, chainID)
	default:
		panic(fmt.Errorf("unknown p2p client by %d %s %s", chainTyp, p2pAddr, chainID))
	}
}
