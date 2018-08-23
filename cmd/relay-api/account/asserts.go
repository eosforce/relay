package account

import (
	"net/http"

	"github.com/eosforce/relay/db"
	"github.com/eosforce/relay/types"
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

	resData, err := db.GetAccountTokens(params.AccountName, params.Chain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := getAccountAssetsRsp{
		AccountName: params.AccountName,
		Chain:       params.Chain,
		Assets:      resData,
	}

	c.JSON(http.StatusOK, res)
}
