package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanh2k4/Chat-app/configs"
	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
	"github.com/thanh2k4/Chat-app/pkg/database/redis"
)

func main() {

	// Load the config
	cfg, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("❗ Falied to load config: %v", err)
	}

	// Connect to Postgres
	pgPool, err := postgres.NewPostgresDB(cfg.Database.Postgres)

	if err != nil {
		log.Fatalf("❗ Failed to connect to Postgres: %v", err)
	}

	defer pgPool.Close()

	log.Println("Connected to PostgreSQL successfully 🚀🎉🎊")

	// Connect to Redis
	redisPool := redis.NewRedisDB(cfg.Database.Redis)

	defer redisPool.Close()

	log.Println("Connected to Redis successfully 🚀🎉🎊")

	// Start the server
	r := gin.Default()
	serverPort := cfg.Server.ServerPort["auth"]
	log.Printf("Starting Auth Service on port %s", serverPort)

	err = r.Run(serverPort)
	if err != nil {
		panic(err)
	}

}
