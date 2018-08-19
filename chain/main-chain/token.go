package mainChain

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/db"
	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
)

// Msg Handler About Token

// onTokenIn map token from a its chain to relay name space.
func (h *Handler) onTokenIn(account eos.AccountName, asset eos.Asset) error {
	seelog.Debugf("onTokenIn %s %s %s",
		h.Name(), string(account), asset.String())

	// TODO Fee
	chainName := h.Name()
	return db.AddToken(string(account), chainName, types.NewAsset(chainName, asset.Amount, asset.Symbol))
}

// onTokenOut trans token from relay to its chain
func (h *Handler) onTokenOut(account eos.AccountName, asset types.Asset) error {
	seelog.Debugf("onTokenOut %s %s %s",
		h.Name(), string(account), asset.String())
	return nil

}

// TODO By FanYang onTokenExchange
