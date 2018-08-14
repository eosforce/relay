package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getMapsReq struct {
}

type getMapsRsp struct {
	Maps []string `json:"maps"`
}

// getAccountInfo get account info by name
func getMaps(c *gin.Context) {
	var params getMapsReq
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO By FanYang imp
	res := getMapsRsp{}

	c.JSON(http.StatusOK, res)
}
