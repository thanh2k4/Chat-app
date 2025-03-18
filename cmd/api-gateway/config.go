package api_gateway

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		ServerPort string `yaml:"port"`
	} `yaml:"http"`
	GRPC struct {
		AuthHost string `yaml:"auth_host"`
		UserHost string `yaml:"user_host"`
		ChatHost string `yaml:"chat_host"`
	} `yaml:"grpc"`
}

func LoadConfig() (*Config, error) {
	path := "cmd/api-api-gateway/config/config.yml"
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
