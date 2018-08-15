package types

import "github.com/fanyang1988/eos-go"

// Account is account in relay, it map from the chains
type Account struct {
	Chain ChainName       `json:"chain_name"`
	Name  eos.AccountName `json:"name"`
}
