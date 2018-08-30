package relayCliCmd

import (
	"github.com/cihub/seelog"
	"github.com/spf13/cobra"
)

var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "exchange commands",
}

var exchangeTestCmd = &cobra.Command{
	Use:   "test [chain] [name] [amount] [symbol] [toChain] [toSymbol]",
	Short: "token in relay to",

	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		defer seelog.Flush()
	},
}

func init() {
	exchangeCmd.AddCommand(exchangeTestCmd)
	RootCmd.AddCommand(exchangeCmd)
}
