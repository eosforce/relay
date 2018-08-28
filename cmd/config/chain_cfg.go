package config

import "github.com/eosforce/relay/db"

// ChainCfg config for a chain
type ChainCfg struct {
	ApiURL       string   `json:"apiUrl"`
	P2PAddresses []string `json:"p2p"`
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	URL          string   `json:"url"` // url for wallet
}

type RelayCfg struct {
	Main ChainCfg       `json:"main"`
	Side []ChainCfg     `json:"side"`
	DB   db.PostgresCfg `json:"db"`
}
