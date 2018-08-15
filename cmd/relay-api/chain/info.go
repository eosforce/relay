package chain

import (
	"net/http"

	"github.com/eosforce/relay/types"
	"github.com/gin-gonic/gin"
)

type getChainInfoReq struct {
	ChainName string `json:"name" form:"name" binding:"required"`
}

type getChainInfoRsp struct {
	ChainName string `json:"account"`
	Note      string `json:"note"`
}

// getAccountInfo get account info by name
func getChainInfo(c *gin.Context) {
	var params getChainInfoReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getChainInfoRsp{
		ChainName: params.ChainName,
		Note:      "ddddddddd",
	}

	c.JSON(http.StatusOK, res)
}

type getChainsRsp struct {
	ChainNames []types.Chain `json:"chains"`
}

// getAccountInfo get account info by name
func getChains(c *gin.Context) {
	// TODO By FanYang imp
	res := getChainsRsp{
		ChainNames: []types.Chain{
			{"main", "main chain"},
			{"side", "side chain"},
		},
	}

	c.JSON(http.StatusOK, res)
}
