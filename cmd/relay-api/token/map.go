package token

import (
	"net/http"

	"github.com/eoscanada/eos-go"
	"github.com/eosforce/relay/types"
	eosforce "github.com/fanyang1988/eos-go"
	"github.com/gin-gonic/gin"
)

type getMapsRsp struct {
	Maps []types.Symbol `json:"maps"`
}

// getAccountInfo get account info by name
func getMaps(c *gin.Context) {

	// TODO By FanYang imp
	res := getMapsRsp{
		Maps: []types.Symbol{
			types.FromEosforceSymbol("main", eosforce.EOSSymbol),
			types.FromEosSymbol("side", eos.EOSSymbol),
		},
	}

	c.JSON(http.StatusOK, res)
}
