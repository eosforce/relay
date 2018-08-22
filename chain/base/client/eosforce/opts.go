package eosforce

import (
	"fmt"
	"time"

	"github.com/eosforce/relay/types"
	eos "github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/eosforce"
	"github.com/fanyang1988/eos-go/token"
	"github.com/pkg/errors"
)

// Transfer transfer token to account
func (e *Client) Transfer(from, to string, asset types.Asset) error {
	var err error
	memo := fmt.Sprintf("by relay %d %s", e.waterfallNum, time.Now())

	if asset.Symbol.Symbol == eos.EOSSymbol.Symbol {
		_, err = e.PushEOSCActions(eosforce.NewTransfer(
			eosforce.AN(from),
			eosforce.AN(to),
			eos.Asset{
				Amount: asset.Amount,
				Symbol: eos.EOSSymbol,
			},
			memo,
		))
	} else {
		_, err = e.PushEOSCActions(token.NewTransfer(
			eosforce.AN(from),
			eosforce.AN(to),
			eos.Asset{
				Amount: asset.Amount,
				Symbol: eos.Symbol{
					Precision: asset.Precision,
					Symbol:    asset.Symbol.Symbol,
				},
			},
			memo,
		))
	}
	e.waterfallNum++
	if err != nil {
		return errors.WithMessage(err, "Transfer")
	}

	return nil
}
