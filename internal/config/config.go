package config

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"os"
)

//Config - конфиг
type Config struct {
	TgToken      string `envconfig:"TG_TOKEN"`
	PingTimeout  int    `envconfig:"PING_TIMEOUT"`
	WorkersCount int    `envconfig:"WORKERS"`
}

//Load - загрузить конфиг по указанному пути
func Load(path string) (*Config, error) {
	config := Config{}
	file, err := os.Open(path)
	if err != nil {
		panic("load config error: " + err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return &config, err
	}
	err = envconfig.Process("", &config)
	return &config, err
}
