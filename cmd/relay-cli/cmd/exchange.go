package relayCliCmd

import (
	"strconv"

	"github.com/cihub/seelog"
	"github.com/eosforce/relay/sdk"
	"github.com/eosforce/relay/token"
	"github.com/eosforce/relay/types"
	"github.com/spf13/cobra"
)

var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "exchange commands",
}

var exchangeTestCmd = &cobra.Command{
	Use:   "test [chain] [name] [amount] [symbol] [toChain] [toSymbol]",
	Short: "token in relay to like `test main eosforce 1000 EOS side SYS`",

	Args: cobra.ExactArgs(6),
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

		toSymbol, err := token.GetSymbol(args[4], args[5])
		if err != nil {
			panic(err)
		}

		err = sdk.ExchangeTest(args[1], args[0], types.Asset{
			Amount: int64(amount),
			Symbol: fromSymbol,
		}, toSymbol)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	exchangeCmd.AddCommand(exchangeTestCmd)
	RootCmd.AddCommand(exchangeCmd)
}
