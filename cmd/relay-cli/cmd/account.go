package relayCliCmd

import (
	"github.com/cihub/seelog"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account commands",
}

var accountMapCmd = &cobra.Command{
	Use:   "map [chain] [name]",
	Short: "map account from chain to relay",

	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		defer seelog.Flush()
	},
}

func init() {
	accountCmd.AddCommand(accountMapCmd)
	RootCmd.AddCommand(accountCmd)
}
