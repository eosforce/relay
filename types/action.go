package types

import (
	"github.com/eoscanada/eos-go"
	force "github.com/fanyang1988/eos-go"
)

// ActionData action data with block and transaction id
type ActionData struct {
	BlockNum uint32
	BlockID  SHA256Bytes
	TrxID    SHA256Bytes
	Name     string
	Account  string
	Data     interface{}
}

func ActionDataFromEos(blockID eos.SHA256Bytes, blockNum uint32, trxID eos.SHA256Bytes, act *eos.Action) ActionData {
	return ActionData{
		BlockNum: blockNum,
		TrxID:    SHA256BytesFromEos(trxID),
		BlockID:  SHA256BytesFromEos(blockID),
		Name:     string(act.Name),
		Account:  string(act.Account),
		Data:     act.Data,
	}
}

func ActionDataFromForce(blockID force.SHA256Bytes, blockNum uint32, trxID force.SHA256Bytes, act *force.Action) ActionData {
	return ActionData{
		BlockNum: blockNum,
		TrxID:    SHA256BytesFromForce(trxID),
		BlockID:  SHA256BytesFromForce(blockID),
		Name:     string(act.Name),
		Account:  string(act.Account),
		Data:     act.Data,
	}
}
