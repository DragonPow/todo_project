package config

import "github.com/DragonPow/todo_project/library/server"

type ServerConfig struct {
	GrpcConfig server.ListenAddr `json:"grpc" mapstructure:"grpc"`
	HttpConfig server.ListenAddr `json:"http" mapstructure:"http"`
}

type Listen struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		GrpcConfig: server.ListenAddr{
			Host: "0.0.0.0",
			Port: 10051,
		},
		HttpConfig: server.ListenAddr{
			Host: "0.0.0.0",
			Port: 10050,
		},
	}
}
