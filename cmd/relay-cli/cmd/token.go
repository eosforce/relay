package relayCliCmd

import (
	"github.com/cihub/seelog"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "token commands",
}

var tokenInCmd = &cobra.Command{
	Use:   "in [chain] [name] [amount] [symbol]",
	Short: "token in relay to",

	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		defer seelog.Flush()
	},
}

var tokenOutCmd = &cobra.Command{
	Use:   "out [chain] [name] [toChain] [toName] [amount] [symbol]",
	Short: "token out from relay",

	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		defer seelog.Flush()
	},
}

func init() {
	tokenCmd.AddCommand(tokenInCmd)
	tokenCmd.AddCommand(tokenOutCmd)
	RootCmd.AddCommand(tokenCmd)
}
