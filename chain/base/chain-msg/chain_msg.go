package chainMsg

import "github.com/eosforce/relay/types"

// ChainMsg actions from chain can be see as a msg from chain,
// some will be process to complete relay.
type ChainMsg struct {
	Account   string // account which send the msg
	MsgTyp    string
	ExtParams []string
	ChainName string
	PA        types.Asset // PA asset param
}
