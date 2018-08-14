package chain

import "github.com/eosforce/relay/cmd/relay-api/http"

func init() {
	http.Router.POST("/v1/get/chain", getChains)
	http.Router.POST("/v1/get/chain/info", getChainInfo)
}
