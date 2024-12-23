package server

import (
	"fmt"
)

type ListenAddr struct {
	Host string
	Port int
}

// Return the string representation of the ListenAddr with format "host:port"
func (l ListenAddr) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

type Config struct {
	grpcConfig *grpcConfig
	httpConfig *httpConfig
	services   []Service
}

func loadDefaultConfig() *Config {
	return &Config{
		grpcConfig: loadDefaultGrpcConfig(),
		httpConfig: loadDefaultHttpConfig(),
	}
}
