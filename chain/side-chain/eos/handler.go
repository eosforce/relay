package eos

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base/chain-msg"
)

// Handler process ChainMsg From main chain
type Handler struct {
	name string
}

// NewHandler
func NewHandler(name string) *Handler {
	return &Handler{
		name: name,
	}
}

func (h *Handler) Name() string {
	return h.name
}

func (h *Handler) Builder() chainMsg.Builder {
	return &chainMsg.Transfer2MsgBuilder{}
}

func (h *Handler) Reg(msgHandler *chainMsg.Handler) {
	msgHandler.AddHandler("r.acc.map", func(msg *chainMsg.ChainMsg) {
		name := msg.Account
		err := h.onMapAccount(name)
		if err != nil {
			seelog.Errorf("do map account cmd err %s by %s",
				string(name), err.Error())
		}
	})
}
