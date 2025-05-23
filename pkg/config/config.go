package config

import (
	"os"
	"time"

	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
	"github.com/thanh2k4/Chat-app/pkg/database/redis"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Postgres postgres.PostgresConfig `yaml:"postgres"`
		Redis    redis.RedisConfig       `yaml:"redis"`
	} `yaml:"database"`

	Server struct {
		ServerPort string `yaml:"port"`
	} `yaml:"http"`

	JWT struct {
		SecretRefreshKey   string        `yaml:"refresh_token_secret"`
		SecretAccessKey    string        `yaml:"access_token_secret"`
		AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
		RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
	} `yaml:"jwt"`

	GRPC struct {
		AuthServer string `yaml:"auth_host"`
		ChatServer string `yaml:"chat_host"`
		UserServer string `yaml:"user_host"`
	} `yaml:"grpc"`
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
