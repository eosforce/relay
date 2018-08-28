package main

import (
	"flag"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/cmd/config"
	_ "github.com/eosforce/relay/cmd/relay-api/account"
	_ "github.com/eosforce/relay/cmd/relay-api/chain"
	"github.com/eosforce/relay/cmd/relay-api/http"
	_ "github.com/eosforce/relay/cmd/relay-api/token"
	"github.com/eosforce/relay/db"
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
)

// relay http api

var url = flag.String("url", ":8080", "api url")
var logCfg = flag.String("logCfg", "", "log xml cfg file path")
var isDebug = flag.Bool("d", false, "is in debug mode")

func main() {
	defer seelog.Flush()
	flag.Parse()

	var cfg config.RelayCfg

	err := config.LoadJsonCfg("./cfg.json", &cfg)
	if err != nil {
		seelog.Errorf("load cfg err by %s", err.Error())
		return
	}

	db.InitDB(cfg.DB)

	// TODO By FanYang user def symbol
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "main",
		Symbol:    "EOS",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "side",
		Symbol:    "EOS",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "side",
		Symbol:    "SYS",
	})
	token.Reg(types.Symbol{
		Precision: 4,
		Chain:     "main",
		Symbol:    "TST",
	})

	config.InitLogger("relay_api", *logCfg)

	http.InitRouter(*isDebug)

	err = http.Router.Run(*url)
	if err != nil {
		seelog.Errorf("run err by %s", err.Error())
	}
}
