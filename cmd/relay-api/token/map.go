package token

import (
	"net/http"

	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
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
			{
				Symbol: eos.Symbol{Precision: 4, Symbol: "EOS"},
				Chain:  "main",
			},
			{
				Symbol: eos.Symbol{Precision: 4, Symbol: "SYS"},
				Chain:  "side",
			},
		},
	}

	c.JSON(http.StatusOK, res)
}
