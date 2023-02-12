package util

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort int `json:"api_port"`
}

const configFile = "config.json"

var GlobalConfig Config

func InitConfig() error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil
	}

	return json.Unmarshal(data, &GlobalConfig)
}
