package conf

import (
	"github.com/BurntSushi/toml"
)

var global *config

func C() *config {
	if global == nil {
		panic("local config failed")
	}
	return global
}

func LoadConfigFromToml(path string) error {
	cfg := NewConfig()
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return err
	}
	global = cfg
	return nil
}
