package main

import (
	"github.com/thanh2k4/Chat-app/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
)

func main() {

	// Load the config
	cfg, err := config.LoadConfig("cmd/chat/config/config.yaml")
	if err != nil {
		log.Fatalf("❗ Falied to load config: %v", err)
	}

	// Connect to Postgres
	pgPool, err := postgres.NewPostgresDB(cfg.Database.Postgres)
	if err != nil {
		log.Fatalf("❗ Failed to connect to Postgres: %v", err)
	}
	defer pgPool.Close()
	log.Println("Connected to PostgreSQL successfully 🚀")

	// Start the server
	r := gin.Default()
	serverPort := cfg.Server.ServerPort
	log.Printf("Starting Auth Service on port %s", serverPort)
	err = r.Run(":" + serverPort)
	if err != nil {
		panic(err)
	}

}
