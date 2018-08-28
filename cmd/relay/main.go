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
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
)

var logCfg = flag.String("logCfg", "", "log xml cfg file path")
var isDebug = flag.Bool("d", false, "is in debug mode")

func main() {
	defer seelog.Flush()
	flag.Parse()

	config.InitLogger("relay", *logCfg)

	var cfg config.RelayCfg

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
