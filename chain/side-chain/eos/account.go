package eos

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/db"
)

// Msg Handler About Account

// onMapAccount map account from side chain
func (h *Handler) onMapAccount(account string) error {
	seelog.Debugf("OnMapAccount %s %s", h.Name(), account)
	return db.CreateAccount(account, h.Name())
}
