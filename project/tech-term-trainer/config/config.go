package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	MongoDB struct {
		URI          string `yaml:"uri"`
		Database     string `yaml:"database"`
		Collection   string `yaml:"collection"`
		NumDocuments int    `yaml:"num_documents"`
		Theme        string `yaml:"theme"`
	} `yaml:"mongodb"`
}

func LoadConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
