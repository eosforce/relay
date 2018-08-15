package token

import (
	"net/http"

	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
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
				Name: "eos2sys",
				TokenA: types.Symbol{
					Symbol: eos.Symbol{Precision: 4, Symbol: "SYS"},
					Chain:  "main"},
				TokenB: types.Symbol{
					Symbol: eos.Symbol{Precision: 4, Symbol: "SYS"},
					Chain:  "side"},
				Type: "bancor",
			},
		},
	}

	c.JSON(http.StatusOK, res)
}
