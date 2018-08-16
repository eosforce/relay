package eosClient

import (
	"encoding/json"
	"time"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/eos-go"
	"github.com/pkg/errors"
)

// PushEOSCActions push transaction with actions to eos net
func (e *Client) PushEOSCActions(actions ...*eos.Action) (*eos.PushTransactionFullResp, error) {
	opts := &eos.TxOptions{}

	if err := opts.FillFromChain(e.api); err != nil {
		return nil, errors.WithMessage(err,
			"Error fetching tapos + chain_id from the chain")
	}

	tx := eos.NewTransaction(actions, opts)

	// TODO By FanYang cfg
	tx.SetExpiration(60 * time.Second)

	e.api.SetSigner(
		eos.NewWalletSigner(eos.New(e.info.URL), e.info.WalletName))

	signedTx, packedTx, err := e.api.SignTransaction(tx, opts.ChainID, eos.CompressionNone)
	if err != nil {
		return nil, errors.WithMessage(err, "signing trx err")
	}

	if packedTx == nil {
		return nil, errors.New(
			"A signed transaction is required if you want to broadcast it. " +
				"Remove --skip-sign (or add --output-transaction ?)")
	}

	err = printTrx(signedTx)
	if err != nil {
		seelog.Error("printf Trx err by ", err.Error())
	}

	resp, err := e.api.PushTransaction(packedTx)
	if err != nil {
		return nil, errors.WithMessage(err, "pushing transaction err")
	}

	return resp, err
}

func printTrx(signedTx *eos.SignedTransaction) error {
	for _, act := range signedTx.Actions {
		act.SetToServer(false)
	}

	cnt, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		return errors.WithMessage(err, "marshalling json")
	}

	seelog.Trace(string(cnt))
	return nil
}
