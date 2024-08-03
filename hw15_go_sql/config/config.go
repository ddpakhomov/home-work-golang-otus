package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
