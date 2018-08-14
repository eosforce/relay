package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAccountMapsReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
}

type getAccountMapsRsp struct {
	AccountName string   `json:"account"`
	Maps        []string `json:"exchanges"`
}

// getAccountInfo get account info by name
func getAccountMaps(c *gin.Context) {
	var params getAccountMapsReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getAccountMapsRsp{
		AccountName: params.AccountName,
		Maps:        []string{},
	}

	c.JSON(http.StatusOK, res)
}
