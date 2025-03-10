package config

import (
	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Postgres postgres.PostgresConfig `yaml:"postgres"`
	} `yaml:"database"`

	Server struct {
		ServerPort string `yaml:"port"`
	} `yaml:"http"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("cmd/chat/config/config.yml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
