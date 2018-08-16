package mainChain

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base/chain-msg"
	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/eosforce"
)

// Handler process ChainMsg From main chain
type Handler struct {
}

// NewHandler
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Name() string {
	return "chain.main"
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

	msgHandler.AddHandler("r.token.in", func(msg *chainMsg.ChainMsg) {
		name := msg.Account
		transferAct, ok := msg.ActionData.Action.ActionData.Data.(*eosforce.Transfer)
		if !ok {
			seelog.Errorf("token in data error")
			return
		}
		asset := transferAct.Quantity

		// TODO Relay Fee

		err := h.onTokenIn(name, asset)
		if err != nil {
			seelog.Errorf("do token in cmd err %s %s by %s",
				string(name), asset.String(), err.Error())
		}
	})

	msgHandler.AddHandler("r.token.out", func(msg *chainMsg.ChainMsg) {
		name := msg.Account

		if len(msg.ExtParams) < 2 {
			seelog.Errorf("token out no params")
			return
		}

		outChainName := msg.ExtParams[0]
		asset, err := eos.NewAsset(msg.ExtParams[1])

		if err != nil {
			seelog.Errorf("token out err by asset %s %s",
				err.Error(), msg.ExtParams[1])
			return
		}

		err = h.onTokenOut(name, types.NewAsset(outChainName, asset.Amount, asset.Symbol))
		if err != nil {
			seelog.Errorf("do token out cmd err %s %s %s by %s",
				string(name), outChainName, asset.String(), err.Error())
		}
	})
}
