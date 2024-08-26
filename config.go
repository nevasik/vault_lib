package vault

import (
	"encoding/json"
	"os"
)

type Config struct {
	Url string
	Credentials
}
type Credentials struct {
	Role   string
	Secret string
}

func CfgFromFile(path string) (cfg *Config, err error) {
	var data []byte
	if path == "" {
		path = "config/config.json"
	}
	data, err = os.ReadFile(path)
	if err != nil {
		return
	}
	cfg = &Config{}
	if err = json.Unmarshal(data, cfg); err != nil {
		return
	}
	return
}
