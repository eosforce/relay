package db

// ChainData chain data
type ChainData struct {
	Name string `json:"name"`
	Note string `json:"note"`
	Typ  string `json:"typ"`
	// TBD other datas
}

// GetChains get all chains data
func GetChains() ([]ChainData, error) {
	db := Get()

	var res []ChainData
	err := db.Model(&res).Limit(64).Select()

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res[:], nil
}

// GetChain get chain data by name
func GetChain(name string) (*ChainData, error) {
	db := Get()

	var res []ChainData
	err := db.Model(&res).Where("name=?", name).Select()

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}
