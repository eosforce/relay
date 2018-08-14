package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAccountInfoReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
}

type getAccountInfoRsp struct {
	AccountName string `json:"name"`
	// TODO fill other data
}

// getAccountInfo get account info by name
func getAccountInfo(c *gin.Context) {
	var params getAccountInfoReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getAccountInfoRsp{
		AccountName: params.AccountName,
	}

	c.JSON(http.StatusOK, res)
}

type queryAccountInfoReq struct {
	PublicKey string `json:"pubkey" form:"pubkey" binding:"required"`
}

type queryAccountInfoRsp struct {
	AccountName []string `json:"account"`
	// TODO fill other data
}

// getAccountInfo get account info by name
func queryAccountInfo(c *gin.Context) {
	var params queryAccountInfoReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := queryAccountInfoRsp{
		AccountName: []string{
			"eosio.acc1",
			"eosio.acc2",
			"eosio.acc3",
			"eosio.acc4",
			"eosio.acc5",
		},
	}

	c.JSON(http.StatusOK, res)
}
