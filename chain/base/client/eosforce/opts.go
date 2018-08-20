package eosforce

import (
	"fmt"
	"time"

	eos "github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/eosforce"
	"github.com/fanyang1988/eos-go/token"
	"github.com/pkg/errors"
)

// Transfer transfer token to account
func (e *Client) Transfer(from, to string, asset eos.Asset) error {
	var err error
	memo := fmt.Sprintf("by relay %d %s", e.waterfallNum, time.Now())

	if asset.Symbol == eos.EOSSymbol {
		_, err = e.PushEOSCActions(eosforce.NewTransfer(
			eosforce.AN(from),
			eosforce.AN(to),
			asset,
			memo,
		))
	} else {
		_, err = e.PushEOSCActions(token.NewTransfer(
			eosforce.AN(from),
			eosforce.AN(to),
			asset,
			memo,
		))
	}
	e.waterfallNum++
	if err != nil {
		return errors.WithMessage(err, "Transfer")
	}

	return nil
}
