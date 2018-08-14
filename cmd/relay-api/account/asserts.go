package account

import (
	"net/http"

	"github.com/fanyang1988/eos-go"
	"github.com/gin-gonic/gin"
)

type getAccountAssetsReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
}

type getAccountAssetsRsp struct {
	AccountName string   `json:"account"`
	Assets      []string `json:"asset"`
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
		Assets: []string{
			eos.Asset{Amount: 100300, Symbol: eos.Symbol{Precision: 4, Symbol: "EOS"}}.String(),
			eos.Asset{Amount: 100200, Symbol: eos.Symbol{Precision: 4, Symbol: "SYS"}}.String(),
			eos.Asset{Amount: 100400, Symbol: eos.Symbol{Precision: 4, Symbol: "CCC"}}.String(),
			eos.Asset{Amount: 103000, Symbol: eos.Symbol{Precision: 4, Symbol: "AAA"}}.String(),
			eos.Asset{Amount: 101000, Symbol: eos.Symbol{Precision: 4, Symbol: "DDD"}}.String(),
		},
	}

	c.JSON(http.StatusOK, res)
}
