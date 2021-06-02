package public

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppHost string   `json:"app_host"`
	AppPort string   `json:"app_port"`
	Nodes   []string `json:"nodes"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&_cfg)
	if err != nil {
		return err
	}
	return nil
}
