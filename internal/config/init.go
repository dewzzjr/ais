package config

import (
	"encoding/json"
	"sync"

	"github.com/jinzhu/configor"
)

var config *Config
var configLock = &sync.Mutex{}

func Load() error {
	configLock.Lock()
	defer configLock.Unlock()
	tmpConfig := Config{}
	err := configor.Load(&tmpConfig)
	if err != nil {
		return err
	}
	config = &tmpConfig
	return nil
}

func Instance() *Config {
	if config != nil {
		return config
	}
	err := Load()
	if err != nil {
		panic(err)
	}
	return config
}

func (*Config) String() string {
	cfg, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(cfg)
}
