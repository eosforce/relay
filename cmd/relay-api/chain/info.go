package chain

import (
	"net/http"

	"github.com/eosforce/relay/db"
	"github.com/gin-gonic/gin"
)

type getChainInfoReq struct {
	ChainName string `json:"name" form:"name" binding:"required"`
}

type getChainInfoRsp struct {
	ChainName string `json:"account"`
	Note      string `json:"note"`
	Typ       string `json:"typ"`
}

// getAccountInfo get account info by name
func getChainInfo(c *gin.Context) {
	var params getChainInfoReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := db.GetChain(params.ChainName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data == nil {
		c.JSON(http.StatusOK, getChainInfoRsp{})
		return
	}

	res := getChainInfoRsp{
		ChainName: params.ChainName,
		Note:      data.Note,
		Typ:       data.Typ,
	}

	c.JSON(http.StatusOK, res)
}

type getChainsRsp struct {
	ChainNames []db.ChainData `json:"chains"`
}

// getAccountInfo get account info by name
func getChains(c *gin.Context) {
	data, err := db.GetChains()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := getChainsRsp{
		ChainNames: data,
	}

	c.JSON(http.StatusOK, res)
}
