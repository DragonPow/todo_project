package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	fmt.Println("Loading config...")

	// Step 1: load default
	config := loadDefault()

	// Step 2: load from file if exists
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Step 3: load from env
	viper.AutomaticEnv()
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	fmt.Println(config)

	return config, nil
}
