package main

// relay node

import (
	"flag"

	"time"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain"
	"github.com/eosforce/relay/cmd/logger-cfg"
)

var logCfg = flag.String("logCfg", "", "log xml cfg file path")
var isDebug = flag.Bool("d", false, "is in debug mode")

func main() {
	defer seelog.Flush()
	flag.Parse()

	loggerCfg.InitLogger("relay", *logCfg)

	// for debug
	mainOpt := chain.WatchOpt{
		ApiURL: "http://127.0.0.1:8890",
		P2PAddresses: []string{
			"127.0.0.1:9002",
			"127.0.0.1:9003",
			"127.0.0.1:9004",
			"127.0.0.1:9005",
			"127.0.0.1:9006",
			"127.0.0.1:9007",
			"127.0.0.1:9008",
			"127.0.0.1:9009",
			"127.0.0.1:9010",
		},
	}

	manager := chain.NewManager(mainOpt)
	err := manager.Start()
	if err != nil {
		seelog.Errorf("start chain mng err By %s", err.Error())
		return
	}

	time.Sleep(5 * time.Second)

	manager.Stop()
}
