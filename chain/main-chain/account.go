package mainChain

import (
	"github.com/cihub/seelog"
	"github.com/fanyang1988/eos-go"
)

// Msg Handler About Account

// onMapAccount map account from main chain
func (h *Handler) onMapAccount(account eos.AccountName) error {
	seelog.Debugf("OnMapAccount %s %s", h.Name(), string(account))
	return nil
}
