package chainMsg

import (
	"strings"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/base/eos-handler"
	"github.com/fanyang1988/eos-go/eosforce"
	"github.com/pkg/errors"
)

// Builder build msg from action got by watcher
type Builder interface {
	Build(action *eosHandler.ActionData) (*ChainMsg, error)
}

// Transfer2MsgBuilder now we use transfer to emit some msg, just for develop fast
type Transfer2MsgBuilder struct {
}

// Build build msg from transfer action, other will gen nil
func (t *Transfer2MsgBuilder) Build(action *eosHandler.ActionData) (*ChainMsg, error) {
	if (action.Action.Name != "transfer" || action.Action.Account != "eosio") &&
		(action.Action.Name != "token.transfer" || action.Action.Account != "eosio") {
		return nil, errors.New("no transfer action in main chain")
	}

	transferAct, ok := action.Action.ActionData.Data.(*eosforce.Transfer)
	if !ok {
		return nil, seelog.Errorf("transfer act data err")
	}

	res := &ChainMsg{
		ActionData: *action,
		Account:    transferAct.From,
		MsgTyp:     string(transferAct.To),
		ExtParams:  strings.Split(transferAct.Memo, ":"),
	}

	return res, nil

}
