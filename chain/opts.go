package chain

const (
	// TypeBaseEos chain base eos
	TypeBaseEos = iota + 1

	// TypeBaseEosforce chain base eosforce
	TypeBaseEosforce
)

// WatchOpt optional for watch a chain
type WatchOpt struct {
	Name         string
	ApiURL       string
	P2PAddresses []string
	Type         int
}
