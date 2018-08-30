package relayCliCmd

import (
	"strconv"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/sdk"
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
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

		amount, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}

		fromSymbol, err := token.GetSymbol(args[0], args[3])
		if err != nil {
			panic(err)
		}

		err = sdk.TokenMapInRelay(args[1], args[0], types.Asset{
			Amount: int64(amount),
			Symbol: fromSymbol,
		})

		if err != nil {
			panic(err)
		}
	},
}

var tokenOutCmd = &cobra.Command{
	Use:   "out [chain] [name] [toChain] [toName] [amount] [symbol]",
	Short: "token out from relay",

	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		defer seelog.Flush()

		amount, err := strconv.Atoi(args[4])
		if err != nil {
			panic(err)
		}

		symbol, err := token.GetSymbol(args[2], args[5])
		if err != nil {
			panic(err)
		}

		err = sdk.TokenMapOutRelay(args[1], args[0], args[3], types.Asset{
			Amount: int64(amount),
			Symbol: symbol,
		})
	},
}

func init() {
	tokenCmd.AddCommand(tokenInCmd)
	tokenCmd.AddCommand(tokenOutCmd)
	RootCmd.AddCommand(tokenCmd)
}
