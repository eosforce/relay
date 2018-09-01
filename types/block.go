package types

// SignedBlockInterface interface for signed block to eosforce or eos
type SignedBlockInterface interface {
	BlockNumber() uint32
}
