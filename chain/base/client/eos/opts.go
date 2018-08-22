package eosClient

import (
	"fmt"
	"time"

	"github.com/eosforce/relay/types"
	eos "github.com/fanyang1988/eos-go"
	"github.com/fanyang1988/eos-go/token"
	"github.com/pkg/errors"
)

// Transfer transfer token to account
func (e *Client) Transfer(from, to string, asset types.Asset) error {
	_, err := e.PushEOSCActions(token.NewTransfer(
		eos.AN(from),
		eos.AN(to),
		eos.Asset{
			Amount: asset.Amount,
			Symbol: eos.Symbol{
				Precision: asset.Precision,
				Symbol:    asset.Symbol.Symbol,
			},
		},
		fmt.Sprintf("by relay %d %s", e.waterfallNum, time.Now()),
	))
	e.waterfallNum++
	if err != nil {
		return errors.WithMessage(err, "Transfer")
	}

	return nil
}
