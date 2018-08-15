package chainMsg

import "github.com/eosforce/relay/chain/base/eos-handler"

// ChainMsg actions from chain can be see as a msg from chain,
// some will be process to complete relay.
type ChainMsg struct {
	ActionData eosHandler.ActionData
	MsgTyp     string
	ExtParams  []string
}
