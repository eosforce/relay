package account

import "github.com/eosforce/relay/cmd/relay-api/http"

func init() {
	http.Router.POST("/v1/get/account", getAccountInfo)
	http.Router.POST("/v1/query/account", queryAccountInfo)
	http.Router.POST("/v1/get/account/maps", getAccountMaps)
	http.Router.POST("/v1/get/account/exchanges", getAccountExchanges)
	http.Router.POST("/v1/get/account/asserts", getAccountAssets)
}
