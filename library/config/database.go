package config

import "fmt"

type Database interface {
	DSN() string
}

type PostgreSql struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	UserName string `json:"user_name" mapstructure:"user_name"`
	Password string `json:"password" mapstructure:"password"`
	DbName   string `json:"db_name" mapstructure:"db_name"`
}

func (cfg PostgreSql) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DbName)
}

func LoadDefaultPostgreSql() PostgreSql {
	return PostgreSql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "",
		DbName:   "",
	}
}
