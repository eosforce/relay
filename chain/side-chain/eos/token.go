package eos

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base/client"
	"github.com/eosforce/relay/const"
	"github.com/eosforce/relay/db"
	"github.com/eosforce/relay/types"
)

// Msg Handler About Token

// onTokenIn map token from a its chain to relay name space.
func (h *Handler) onTokenIn(account string, asset types.Asset) error {
	seelog.Debugf("onTokenIn %s %s %s",
		h.Name(), account, asset.String())

	// TODO Fee
	chainName := h.Name()
	return db.AddToken(account, chainName, asset)
}

// onTokenOut trans token from relay to its chain
func (h *Handler) onTokenOut(account string, asset types.Asset) error {
	seelog.Debugf("onTokenOut %s %s %s",
		h.Name(), account, asset.String())

	c := client.Get(h.Name())

	// TODO Fee

	// TODO Process Err

	db.CostToken(account, h.Name(), asset)

	return c.Transfer(consts.TokenInAcc, account, asset)
}

// TODO By FanYang onTokenExchange
