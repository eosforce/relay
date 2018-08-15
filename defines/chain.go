package defines

// ChainName chain name in relay
type ChainName string

// Chain is the chain connected to relay
type Chain struct {
	Name ChainName `json:"chain_name"`
	Note string    `json:"note"`
}
