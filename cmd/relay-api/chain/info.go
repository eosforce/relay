package chain

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getChainInfoReq struct {
	ChainName string `json:"name" form:"name" binding:"required"`
}

type getChainInfoRsp struct {
	ChainName string `json:"account"`
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
	}

	c.JSON(http.StatusOK, res)
}

type getChainsReq struct {
}

type getChainsRsp struct {
	ChainNames []string `json:"chains"`
}

// getAccountInfo get account info by name
func getChains(c *gin.Context) {
	var params getChainsReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getChainsRsp{
		ChainNames: []string{
			"main",
			"eos.test",
		},
	}

	c.JSON(http.StatusOK, res)
}
