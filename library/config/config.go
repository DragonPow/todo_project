package config

import (
	"github.com/DragonPow/todo_project/library/log"
)

type Base struct {
	Log    log.Config   `json:"log" mapstructure:"log"`
	Server ServerConfig `json:"server" mapstructure:"server"`
}

func LoadDefaultBase() Base {
	return Base{
		Log:    log.DefaulConfig(),
		Server: DefaultServerConfig(),
	}
}

func (cfg Base) BuildLogger() (*log.Logger, error) {
	return cfg.Log.Build()
}
