package account

import (
	"net/http"

	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
	"github.com/gin-gonic/gin"
)

type getAccountAssetsReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
	Chain       string `json:"chain" form:"chain" binding:"required"`
}

type getAccountAssetsRsp struct {
	AccountName string        `json:"account"`
	Chain       string        `json:"chain"`
	Assets      []types.Asset `json:"asset"`
}

// getAccountInfo get account info by name
func getAccountAssets(c *gin.Context) {
	var params getAccountAssetsReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getAccountAssetsRsp{
		AccountName: params.AccountName,
		Chain:       params.Chain,
		Assets: []types.Asset{
			types.NewAsset("main", 11111, eos.Symbol{Precision: 4, Symbol: "EOS"}),
			types.NewAsset("main", 22222, eos.Symbol{Precision: 4, Symbol: "EEE"}),
			types.NewAsset("main", 33333, eos.Symbol{Precision: 4, Symbol: "AAA"}),
			types.NewAsset("side", 44444, eos.Symbol{Precision: 4, Symbol: "SYS"}),
		},
	}

	c.JSON(http.StatusOK, res)
}
