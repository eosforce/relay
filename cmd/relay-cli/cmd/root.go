package relayCliCmd

import (
	"fmt"
	"os"

	"github.com/cihub/seelog"
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

	for _, flag := range []string{"log-cfg"} {
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
}

func initConfig() {
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
