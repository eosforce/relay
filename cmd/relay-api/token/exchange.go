package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getExchangesReq struct {
}

type getExchangesRsp struct {
	Exchanges []string `json:"exchanges"`
}

// getAccountInfo get account info by name
func getExchanges(c *gin.Context) {
	var params getExchangesReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getExchangesRsp{
		Exchanges: []string{},
	}

	c.JSON(http.StatusOK, res)
}
