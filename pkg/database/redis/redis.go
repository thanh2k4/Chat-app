package redis

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewRedisDB(cfg RedisConfig) *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = cfg.Host
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb
}
