package main

import (
	"fmt"
	apigateway "github.com/thanh2k4/Chat-app/internal/api-gateway"
	"github.com/thanh2k4/Chat-app/internal/api-gateway/router"
	"github.com/thanh2k4/Chat-app/pkg/config"

	"log"
)

func main() {
	cfg, err := config.LoadConfig("cmd/api-gateway/config/config.yml")
	if err != nil {
		log.Fatalf("❗ Falied to load config: %v", err)
	}

	client := apigateway.NewClient(*cfg)

	r := router.SetupAPIGatewayRoutes(client)

	err = r.Run(":" + cfg.Server.ServerPort)
	if err != nil {
		panic(err)
	}
	fmt.Println("API Gateway Service is running on port " + cfg.Server.ServerPort + "...")
}
