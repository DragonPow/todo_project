package config

import (
	"github.com/DragonPow/todo_project/library/config"
	"gorm.io/gorm"
)

type Config struct {
	config.Base   `mapstructure:",squash"`
	Database      config.PostgreSql `json:"database" mapstructure:"database"`
	MigrationPath string            `json:"migration_path" mapstructure:"migration_path"`
	MaxIdleConns  int               `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns  int               `json:"max_open_conns" mapstructure:"max_open_conns"`
}

func loadDefault() *Config {
	return &Config{
		Base:          config.LoadDefaultBase(),
		Database:      config.LoadDefaultPostgreSql(),
		MigrationPath: "file://migrations",
	}
}

func (cfg Config) GormConfig() *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction: true,
	}
}
