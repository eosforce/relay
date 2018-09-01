package base

import (
	"encoding/hex"
	"encoding/json"

	eos "github.com/eoscanada/eos-go"
	force "github.com/fanyang1988/eos-go"
)

// SHA256Bytes type for sha256 bytes
type SHA256Bytes []byte // should always be 32 bytes

// SHA256BytesFromEos transfer from eos.SHA256Bytes
func SHA256BytesFromEos(s eos.SHA256Bytes) SHA256Bytes {
	return SHA256Bytes(s)
}

// SHA256BytesFromForce transfer from eosforce.SHA256Bytes
func SHA256BytesFromForce(s force.SHA256Bytes) SHA256Bytes {
	return SHA256Bytes(s)
}

// MarshalJSON to json
func (t SHA256Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(t))
}

// UnmarshalJSON from json
func (t *SHA256Bytes) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = hex.DecodeString(s)
	return
}

func (t SHA256Bytes) String() string {
	return hex.EncodeToString(t)
}
