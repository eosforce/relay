package account

import (
	"net/http"

	"github.com/eosforce/relay/types"
	"github.com/fanyang1988/eos-go"
	"github.com/gin-gonic/gin"
)

type getAccountExchangesReq struct {
	AccountName string `json:"name" form:"name" binding:"required"`
	Chain       string `json:"chain" form:"chain" binding:"required"`
}

type getAccountExchangesRsp struct {
	AccountName string                  `json:"account"`
	Chain       string                  `json:"chain"`
	Exchanges   []types.ExchangeHistory `json:"exchanges"`
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
		Chain:       params.Chain,
		Exchanges: []types.ExchangeHistory{
			{
				Exchange: "test",
				From: types.Account{
					Chain: types.ChainName(params.Chain),
					Name:  eos.AccountName(params.AccountName),
				},
				To: types.Account{
					Chain: "main",
					Name:  "others",
				},
				FromToken: types.NewAsset("main", 11111, eos.Symbol{Precision: 4, Symbol: "EOS"}),
				ToToken:   types.NewAsset("main", 22222, eos.Symbol{Precision: 4, Symbol: "SYS"}),
			},
			{
				Exchange: "test",
				From: types.Account{
					Chain: types.ChainName(params.Chain),
					Name:  eos.AccountName(params.AccountName),
				},
				To: types.Account{
					Chain: "main",
					Name:  "others",
				},
				FromToken: types.NewAsset("main", 11111, eos.Symbol{Precision: 4, Symbol: "EOS"}),
				ToToken:   types.NewAsset("main", 22222, eos.Symbol{Precision: 4, Symbol: "SYS"}),
			},
		},
	}

	c.JSON(http.StatusOK, res)
}
