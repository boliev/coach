package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %s ", err.Error())
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err = viper.MergeInConfig() // Find and read the config file
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %s ", err.Error())
	}

	if _, err := os.Stat("./.env.local"); err == nil {
		viper.SetConfigName(".env.local")
		err = viper.MergeInConfig() // Find and read the config file
		if err != nil {
			return nil, fmt.Errorf("fatal error config file: %s ", err.Error())
		}
	}
	return &Config{}, nil
}

func (c Config) GetString(key string) string {
	return viper.GetString(key)
}

func (c Config) GetInt(key string) int {
	return viper.GetInt(key)
}
