package main

// relay node

import (
	"flag"

	"time"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain"
	"github.com/eosforce/relay/chain/wallets"
	"github.com/eosforce/relay/cmd/logger-cfg"
	"github.com/eosforce/relay/db"
)

var logCfg = flag.String("logCfg", "", "log xml cfg file path")
var isDebug = flag.Bool("d", false, "is in debug mode")

func main() {
	defer seelog.Flush()
	flag.Parse()

	loggerCfg.InitLogger("relay", *logCfg)

	// for debug
	mainOpt := chain.WatchOpt{
		ApiURL: "http://127.0.0.1:8001",
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
		Type: chain.TypeBaseEosforce,
		Name: "main",
	}
	wallets.Get().RegWallet(wallets.WalletInfo{
		APIURL:     "http://127.0.0.1:8001",
		URL:        "http://127.0.0.1:6666",
		ChainName:  "main",
		WalletName: "default",
	})

	sideOpt := chain.WatchOpt{
		ApiURL: "http://127.0.0.1:18001",
		P2PAddresses: []string{
			"127.0.0.1:19002",
			"127.0.0.1:19003",
			"127.0.0.1:19004",
			"127.0.0.1:19005",
			"127.0.0.1:19006",
			"127.0.0.1:19007",
			"127.0.0.1:19008",
			"127.0.0.1:19009",
			"127.0.0.1:19010",
		},
		Type: chain.TypeBaseEos,
		Name: "side",
	}
	wallets.Get().RegWallet(wallets.WalletInfo{
		APIURL:     "http://127.0.0.1:18001",
		URL:        "http://127.0.0.1:16666",
		ChainName:  "side",
		WalletName: "default",
	})

	db.InitDB(db.PostgresCfg{
		Address:  "127.0.0.1:5432",
		User:     "pgfy",
		Password: "123456",
		Database: "test3",
	})

	manager := chain.NewManager(mainOpt, sideOpt)
	err := manager.Start()
	if err != nil {
		seelog.Errorf("start chain mng err By %s", err.Error())
		return
	}

	defer manager.Stop()
	for {
		time.Sleep(5 * time.Second)
	}
}
