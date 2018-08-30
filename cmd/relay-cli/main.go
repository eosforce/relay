package main

import (
	"github.com/cihub/seelog"
	"github.com/eosforce/relay/cmd/relay-cli/cmd"
)

func main() {
	defer seelog.Flush()
	relayCliCmd.Execute()
}
