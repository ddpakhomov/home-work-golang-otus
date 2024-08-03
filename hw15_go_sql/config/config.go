package config

import (
	"github.com/spf13/viper"
)

// Config структура для хранения конфигурации
type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig структура для хранения конфигурации базы данных
type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}

// LoadConfig загружает конфигурацию из файла
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
