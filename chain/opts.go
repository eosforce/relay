package chain

import (
	"fmt"

	"github.com/eosforce/relay/const"

	"github.com/eosforce/relay/cmd/config"
)

// WatchOpt optional for watch a chain
type WatchOpt struct {
	Name         string
	ApiURL       string
	P2PAddresses []string
	Type         int
}

// NewWatchOptByCfg init a opt from cfg
func NewWatchOptByCfg(cfg *config.ChainCfg) WatchOpt {
	typ := 0
	switch cfg.Type {
	case "eos":
		typ = consts.TypeBaseEos
	case "eosforce":
		typ = consts.TypeBaseEosforce
	default:
		panic(fmt.Errorf("chain %s typ %s unknown", cfg.Name, cfg.Type))
	}
	return WatchOpt{
		Name:         cfg.Name,
		ApiURL:       cfg.ApiURL,
		P2PAddresses: cfg.P2PAddresses,
		Type:         typ,
	}
}
