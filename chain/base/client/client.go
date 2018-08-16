package client

import "github.com/fanyang1988/eos-go"

// EosClient interface to eos Client or eosforce Client
type EosClient interface {
	PushEOSCActions(actions ...*eos.Action) (*eos.PushTransactionFullResp, error)
}
