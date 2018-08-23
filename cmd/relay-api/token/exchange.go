package token

import (
	"net/http"

	"github.com/eoscanada/eos-go"
	"github.com/eosforce/relay/types"
	eosforce "github.com/fanyang1988/eos-go"
	"github.com/gin-gonic/gin"
)

type getExchangesRsp struct {
	Exchanges []types.ExchangePair `json:"exchanges"`
}

// getAccountInfo get account info by name
func getExchanges(c *gin.Context) {

	// TODO By FanYang imp
	res := getExchangesRsp{
		Exchanges: []types.ExchangePair{
			{
				Name:   "eos2sys",
				TokenA: types.FromEosforceSymbol("main", eosforce.EOSSymbol),
				TokenB: types.FromEosSymbol("side", eos.EOSSymbol),
				Type:   "bancor",
			},
		},
	}

	c.JSON(http.StatusOK, res)
}
