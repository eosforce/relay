package eos

import (
	"github.com/cihub/seelog"
	"github.com/eoscanada/eos-go"
	"github.com/eosforce/relay/chain/base/chain-msg"
	"github.com/eosforce/relay/const"
	"github.com/eosforce/relay/types"
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
	return &chainMsg.Transfer2MsgBuilder{
		ChainName: h.Name(),
	}
}

func (h *Handler) Reg(msgHandler *chainMsg.Handler) {
	msgHandler.AddHandler(h.Name(), consts.AccountMapAcc, func(msg *chainMsg.ChainMsg) {
		name := msg.Account
		err := h.onMapAccount(name)
		if err != nil {
			seelog.Errorf("do map account cmd err %s by %s",
				string(name), err.Error())
		}
	})

	msgHandler.AddHandler(h.Name(), consts.TokenInAcc, func(msg *chainMsg.ChainMsg) {
		name := msg.Account
		asset := msg.PA

		// TODO Relay Fee

		err := h.onTokenIn(name, asset)
		if err != nil {
			seelog.Errorf("do token in cmd err %s %s by %s",
				string(name), asset.String(), err.Error())
		}
	})

	msgHandler.AddHandler(h.Name(), consts.TokenOutAcc, func(msg *chainMsg.ChainMsg) {
		name := msg.Account

		if len(msg.ExtParams) < 1 {
			seelog.Errorf("token out no params")
			return
		}

		asset, err := eos.NewAsset(msg.ExtParams[0])

		if err != nil {
			seelog.Errorf("token out err by asset %s %s",
				err.Error(), msg.ExtParams[0])
			return
		}

		err = h.onTokenOut(name, types.NewAsset(asset.Amount, types.FromEosSymbol(msg.ChainName, asset.Symbol)))
		if err != nil {
			seelog.Errorf("do token out cmd err %s %s %s by %s",
				string(name), h.Name(), asset.String(), err.Error())
		}
	})
}
