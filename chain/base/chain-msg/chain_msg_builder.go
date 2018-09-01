package chainMsg

import (
	"strings"

	"github.com/cihub/seelog"
	"github.com/eoscanada/eos-go/token"
	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go/eosforce"
	"github.com/pkg/errors"
)

// Builder build msg from action got by watcher
type Builder interface {
	BuildFromEOS(action *types.ActionData) (*ChainMsg, error)
	BuildFromEOSForce(action *types.ActionData) (*ChainMsg, error)
}

// Transfer2MsgBuilder now we use transfer to emit some msg, just for develop fast
type Transfer2MsgBuilder struct {
	ChainName string
}

// Build build msg from transfer action, other will gen nil
func (t *Transfer2MsgBuilder) BuildFromEOS(action *types.ActionData) (*ChainMsg, error) {
	if (action.Name != "transfer" || action.Account != "eosio") &&
		(action.Name != "transfer" || action.Account != "eosio.token") {
		return nil, errors.New("no transfer action in main chain")
	}

	transferAct, ok := action.Data.(*token.Transfer)
	if !ok {
		return nil, seelog.Errorf("transfer act data err")
	}

	res := &ChainMsg{
		Account:   string(transferAct.From),
		MsgTyp:    string(transferAct.To),
		ExtParams: strings.Split(transferAct.Memo, ":"),
		ChainName: t.ChainName,
		PA: types.Asset{
			Amount: transferAct.Quantity.Amount,
			Symbol: types.FromEosSymbol(t.ChainName, transferAct.Quantity.Symbol),
		},
	}

	return res, nil

}

// Build build msg from transfer action, other will gen nil
func (t *Transfer2MsgBuilder) BuildFromEOSForce(action *types.ActionData) (*ChainMsg, error) {
	if (action.Name != "transfer" || action.Account != "eosio") &&
		(action.Name != "transfer" || action.Account != "eosio.token") {
		return nil, errors.New("no transfer action in main chain")
	}

	transferAct, ok := action.Data.(*eosforce.Transfer)
	if !ok {
		return nil, seelog.Errorf("transfer act data err")
	}

	res := &ChainMsg{
		Account:   string(transferAct.From),
		MsgTyp:    string(transferAct.To),
		ExtParams: strings.Split(transferAct.Memo, ":"),
		ChainName: t.ChainName,
		PA: types.Asset{
			Amount: transferAct.Quantity.Amount,
			Symbol: types.FromEosforceSymbol(t.ChainName, transferAct.Quantity.Symbol),
		},
	}

	return res, nil

}
