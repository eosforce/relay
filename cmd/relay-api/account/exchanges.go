package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAccountExchangesReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
}

type getAccountExchangesRsp struct {
	AccountName string   `json:"account"`
	Exchanges   []string `json:"exchanges"`
}

// getAccountInfo get account info by name
func getAccountExchanges(c *gin.Context) {
	var params getAccountExchangesReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getAccountExchangesRsp{
		AccountName: params.AccountName,
		Exchanges:   []string{},
	}

	c.JSON(http.StatusOK, res)
}
