package chain

import "github.com/eosforce/relay/chain/base/chain-msg"

// Handler interface of chain handler
type Handler interface {
	Name() string
	Builder() chainMsg.Builder
	Reg(h *chainMsg.Handler)
}
