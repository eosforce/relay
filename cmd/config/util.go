package config

import (
	"fmt"
	"os"

	"io/ioutil"

	"encoding/json"

	"github.com/pkg/errors"
)

// LoadJsonCfg load a cfg by json file
func LoadJsonCfg(path string, cfg interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("load json %s err", path))
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("read json file %s err", path))
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("unmarshal json %s err", path))
	}

	return nil
}
