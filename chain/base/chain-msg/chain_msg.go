package chainMsg

import (
	"github.com/eosforce/relay/chain/base/eos-handler"
	"github.com/fanyang1988/eos-go"
)

// ChainMsg actions from chain can be see as a msg from chain,
// some will be process to complete relay.
type ChainMsg struct {
	ActionData eosHandler.ActionData
	Account    eos.AccountName // account which send the msg
	MsgTyp     string
	ExtParams  []string
}
