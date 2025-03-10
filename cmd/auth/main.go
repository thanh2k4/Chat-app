package main

import (
	"github.com/thanh2k4/Chat-app/cmd/auth/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanh2k4/Chat-app/pkg/database/redis"
)

func main() {

	// Load the config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❗ Falied to load config: %v", err)
	}

	// Connect to Redis
	redisPool := redis.NewRedisDB(cfg.Database.Redis)
	defer redisPool.Close()
	log.Println("Connected to Redis successfully 🚀")

	// Start the server
	r := gin.Default()
	serverPort := cfg.Server.ServerPort
	log.Printf("Starting Auth Service on port %s", serverPort)
	err = r.Run(":" + serverPort)
	if err != nil {
		panic(err)
	}

}
