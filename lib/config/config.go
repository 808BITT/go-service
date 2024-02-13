package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Install *InstallConfig `json:"Install"`
}

type InstallConfig struct {
	Name string `json:"Service-Name"`
	Path string `json:"Install-Path"`
}

func NewConfig(path string) *Config {
	config := &Config{}
	f, e := os.ReadFile(path)
	if e != nil {
		return nil
	}
	e = json.Unmarshal(f, config)
	if e != nil {
		return nil
	}
	return config
}
