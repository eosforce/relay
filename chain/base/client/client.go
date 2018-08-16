package client

import (
	"github.com/eosforce/relay/chain/base/client/eos"
	"github.com/eosforce/relay/chain/base/client/eosforce"
	"github.com/fanyang1988/eos-go"
)

// EosClient interface to eos Client or eosforce Client
type EosClient interface {
	PushEOSCActions(actions ...*eos.Action) (*eos.PushTransactionFullResp, error)
}

// Get get a client to chain by chainName now just a simple imp
func Get(chainName string) EosClient {
	// TODO now just main chain is eosforce
	if chainName == "main" || chainName == "chain.main" {
		return eosforce.NewClient(chainName)
	} else {
		return eosClient.NewClient(chainName)
	}
}
