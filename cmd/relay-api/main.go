package main

import (
	"flag"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/cmd/logger-cfg"
	"github.com/eosforce/relay/cmd/relay-api/http"
	"github.com/gin-gonic/gin"

	_ "github.com/eosforce/relay/cmd/relay-api/account"
	_ "github.com/eosforce/relay/cmd/relay-api/chain"
	_ "github.com/eosforce/relay/cmd/relay-api/token"
)

// relay http api

var url = flag.String("url", ":8080", "api url")
var logCfg = flag.String("logCfg", "", "log xml cfg file path")
var isDebug = flag.Bool("d", false, "is in debug mode")

func main() {
	defer seelog.Flush()
	flag.Parse()

	loggerCfg.InitLogger("relay_api", *logCfg)

	http.InitRouter(*isDebug)

	http.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := http.Router.Run(*url)
	if err != nil {
		seelog.Errorf("run err by %s", err.Error())
	}
}
