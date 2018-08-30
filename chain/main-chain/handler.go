package mainChain

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base/chain-msg"
	"github.com/eosforce/relay/const"
	"github.com/eosforce/relay/db"
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
)

// Handler process ChainMsg From main chain
type Handler struct {
}

// NewHandler
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Name() string {
	return "main"
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

		if len(msg.ExtParams) < 3 {
			seelog.Errorf("token out no params")
			return
		}

		asset, err := eos.NewAsset(msg.ExtParams[2])

		if err != nil {
			seelog.Errorf("token out err by asset %s %s",
				err.Error(), msg.ExtParams[2])
			return
		}

		err = h.onTokenOut(name, msg.ExtParams[0], msg.ExtParams[1],
			types.NewAsset(asset.Amount, types.FromEosforceSymbol(msg.ExtParams[1], asset.Symbol)))
		if err != nil {
			seelog.Errorf("do token out cmd err %s %s %s by %s",
				string(name), h.Name(), asset.String(), err.Error())
		}
	})
	msgHandler.AddHandler(h.Name(), consts.TestExchange, func(msg *chainMsg.ChainMsg) {
		name := msg.Account

		// "1000.0000 EOS:main:SYS:side"

		if len(msg.ExtParams) < 4 {
			seelog.Errorf("exchange no params")
			return
		}

		af, err := eos.NewAsset(msg.ExtParams[0])
		if err != nil {
			seelog.Errorf("exchange err by asset %s %s",
				err.Error(), msg.ExtParams[0])
			return
		}

		asset_to_exchange := types.NewAssetFromEosforce(msg.ExtParams[1], af.Amount, af.Symbol)

		symbol_to_get, err := token.GetSymbol(msg.ExtParams[3], msg.ExtParams[2])
		if err != nil {
			seelog.Errorf("exchange err by exchange symbol %s %s:%s",
				err.Error(), msg.ExtParams[3], msg.ExtParams[2])
			return
		}

		seelog.Debugf("exchange test from %s --> %s", asset_to_exchange, symbol_to_get)
		err = db.ExchangeTokenTest(name, h.Name(), asset_to_exchange, symbol_to_get)

		if err != nil {
			seelog.Errorf("do token out cmd err %s %s -> %s %s by %s",
				string(name), h.Name(), asset_to_exchange, symbol_to_get, err.Error())
		}
	})
}
