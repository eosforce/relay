package account

import (
	"net/http"

	"github.com/eosforce/relay/types"
	"github.com/gin-gonic/gin"
)

type getAccountInfoReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
	Chain       string `json:"chain" form:"chain" binding:"required"`
}

type getAccountInfoRsp struct {
	AccountName string `json:"name"`
	Chain       string `json:"chain"`
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
		Chain:       params.Chain,
	}

	c.JSON(http.StatusOK, res)
}

type queryAccountInfoReq struct {
	PublicKey string `json:"pubkey" form:"pubkey" binding:"required"`
}

type queryAccountInfoRsp struct {
	AccountName []types.Account `json:"account"`
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
		AccountName: []types.Account{
			{"main", "eosio.acc1"},
			{"main", "eosio.acc2"},
			{"main", "eosio.acc3"},
			{"main", "eosio.acc4"},
			{"side", "eosio.acc5"},
		},
	}

	c.JSON(http.StatusOK, res)
}
