package chain

import (
	"fmt"

	"github.com/eosforce/relay/cmd/config"
)

const (
	// TypeBaseEos chain base eos
	TypeBaseEos = iota + 1

	// TypeBaseEosforce chain base eosforce
	TypeBaseEosforce
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
		typ = TypeBaseEos
	case "eosforce":
		typ = TypeBaseEosforce
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
