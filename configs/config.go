package configs

import (
	"os"
	"time"

	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
	"github.com/thanh2k4/Chat-app/pkg/database/redis"
	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	Database struct {
		Postgres postgres.PostgresConfig `yaml:"postgres"`
		Redis    redis.RedisConfig       `yaml:"redis"`
	} `yaml:"database"`

	Server struct {
		ServerPort map[string]string
	} `yaml:"server"`

	JWT struct {
		SecretRefreshKey   string        `yaml:"refresh_token_secret"`
		SecretAccessKey    string        `yaml:"access_token_secret"`
		AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
		RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
	} `yaml:"jwt"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
