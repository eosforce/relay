package relayCliCmd

import (
	"fmt"
	"os"

	"github.com/eosforce/relay/chain/wallets"
	"github.com/eosforce/relay/cmd/config"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "relay-cli",
	Short: "relay-cli is a client send do command to eoforce relay",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("log-cfg", "", "./log_cfg.xml", "seelog cfg file")
	RootCmd.PersistentFlags().StringP("cfg", "", "./cfg.json", "chain cfg file")

	for _, flag := range []string{"log-cfg", "cfg"} {
		if err := viper.BindPFlag(flag, RootCmd.PersistentFlags().Lookup(flag)); err != nil {
			panic(err)
		}
	}

	configFilePath := viper.GetString("log-cfg")
	if configFilePath != "" && IsFileExists(configFilePath) {
		logger, err := seelog.LoggerFromConfigAsFile(configFilePath)
		if err != nil {
			panic(err)
		}
		seelog.ReplaceLogger(logger)
	}

	initConfig()
}

func initConfig() {
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

	var cfg config.RelayCfg
	configFilePath := viper.GetString("cfg")
	err := config.LoadJsonCfg(configFilePath, &cfg)
	if err != nil {
		seelog.Errorf("load cfg err by %s", err.Error())
		return
	}

	// init wallets

	wallets.Get().RegWallet(wallets.NewWalletInfoFromCfg(&cfg.Main))
	for _, sideCfg := range cfg.Side {
		wallets.Get().RegWallet(wallets.NewWalletInfoFromCfg(&sideCfg))
	}
}

// Exists reports whether the named file or directory exists.
func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
