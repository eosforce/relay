package main

// relay node

import (
	"flag"

	"time"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/chain"
	"github.com/eosforce/relay/chain/wallets"
	"github.com/eosforce/relay/cmd/config"
	"github.com/eosforce/relay/db"
)

var logCfg = flag.String("logCfg", "", "log xml cfg file path")
var isDebug = flag.Bool("d", false, "is in debug mode")

func main() {
	defer seelog.Flush()
	flag.Parse()

	var cfg config.RelayCfg

	config.InitLogger("relay", *logCfg)

	err := config.LoadJsonCfg("./cfg.json", &cfg)
	if err != nil {
		seelog.Errorf("load cfg err by %s", err.Error())
		return
	}

	mainOpt := chain.NewWatchOptByCfg(&cfg.Main)
	wallets.Get().RegWallet(wallets.NewWalletInfoFromCfg(&cfg.Main))

	sideOpts := make([]chain.WatchOpt, 0, len(cfg.Side))
	for _, sideCfg := range cfg.Side {
		sideOpt := chain.NewWatchOptByCfg(&sideCfg)
		sideOpts = append(sideOpts, sideOpt)
		wallets.Get().RegWallet(wallets.NewWalletInfoFromCfg(&sideCfg))
	}

	db.InitDB(cfg.DB)

	manager := chain.NewManager(mainOpt, sideOpts...)
	err = manager.Start()
	if err != nil {
		seelog.Errorf("start chain mng err By %s", err.Error())
		return
	}

	defer manager.Stop()
	for {
		time.Sleep(5 * time.Second)
		// TODO exit and ping
	}
}
