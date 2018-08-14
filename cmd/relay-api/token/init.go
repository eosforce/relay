package token

import "github.com/eosforce/relay/cmd/relay-api/http"

func init() {
	http.Router.POST("/v1/get/token/maps", getMaps)
	http.Router.POST("/v1/get/token/exchanges", getExchanges)
}
