package main

import (
	"github.com/thanh2k4/Chat-app/internal/chat"
	postgres2 "github.com/thanh2k4/Chat-app/internal/chat/infras/postgres"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/pkg/database/redis"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
)

func main() {

	// Load the config
	cfg, err := config.LoadConfig("cmd/chat/config/config.yml")
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
	queries := postgres2.New(pgPool)

	// Connect to Redis
	redisPool := redis.NewRedisDB(cfg.Database.Redis)
	defer redisPool.Close()
	log.Println("Connected to Redis successfully 🚀")

	// Start the gRPC server
	grpcServer := grpc.NewServer()
	chatServer := chat.NewChatServer()
	gen.RegisterChatServiceServer(grpcServer, chatServer)
	go func() {
		listener, err := net.Listen("tcp", ":"+cfg.Server.ServerPort)
		if err != nil {
			log.Fatalf("❗ Failed to listen: %v", err)
		}
		log.Printf("Chat Service is running on port %s", cfg.Server.ServerPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC for Chat Service: %v", err)
		}
	}()

	manager := chat.NewManager(redisPool, queries)
	r := chat.SetupChatRoutes(manager)

	err = r.Run(":" + "8080")
	if err != nil {
		panic(err)
	}

}
