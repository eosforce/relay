package account

import (
	"net/http"

	"time"

	"github.com/eosforce/relay/db"
	"github.com/eosforce/relay/types"
	"github.com/gin-gonic/gin"
)

type getAccountInfoReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
	Chain       string `json:"chain" form:"chain" binding:"required"`
}

type getAccountInfoRsp struct {
	AccountName string             `json:"name"`
	Chain       string             `json:"chain"`
	CreateTime  time.Time          `json:"create_time"`
	Permissions []types.Permission `json:"permissions"`
	// TBD fill other data
}

// getAccountInfo get account info by name
func getAccountInfo(c *gin.Context) {
	var params getAccountInfoReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := db.GetAccount(params.AccountName, params.Chain)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data == nil {
		c.JSON(http.StatusOK, getAccountInfoRsp{})
		return
	}

	permissions, err := db.GetAccountPermission(params.AccountName, params.Chain)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := getAccountInfoRsp{
		AccountName: params.AccountName,
		Chain:       params.Chain,
		CreateTime:  data.CreateTime,
		Permissions: permissions,
	}

	c.JSON(http.StatusOK, res)
}

type queryAccountInfoReq struct {
	PublicKey string `json:"pubkey" form:"pubkey" binding:"required"`
}

type queryAccountInfoRsp struct {
	AccountName []types.Account `json:"account"`
}

// getAccountInfo get account info by name
func queryAccountInfo(c *gin.Context) {
	var params queryAccountInfoReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// use owner public key
	data, err := db.QueryAccountByPermission("owner", params.PublicKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := queryAccountInfoRsp{
		AccountName: make([]types.Account, 0, len(data)),
	}

	for _, d := range data {
		res.AccountName = append(res.AccountName, types.Account{
			Name:  d.Name,
			Chain: d.Chain,
		})
	}

	c.JSON(http.StatusOK, res)
}
