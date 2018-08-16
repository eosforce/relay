package eosClient

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain/wallets"
	eos "github.com/fanyang1988/eos-go"
)

// Client client to push trx to eos chain
type Client struct {
	api  *eos.API
	info wallets.WalletInfo
}

func NewClient(chainName string) *Client {
	info, err := wallets.Get().GetInfo(chainName)
	if err != nil {
		seelog.Errorf("no chainName found %s", err.Error())
		return nil
	}

	return &Client{
		api:  eos.New(info.APIURL),
		info: info,
	}
}
