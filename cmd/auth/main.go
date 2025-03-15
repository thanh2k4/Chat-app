package main

import (
	"fmt"
	"github.com/thanh2k4/Chat-app/internal/auth"
	postgres2 "github.com/thanh2k4/Chat-app/internal/auth/infras/postgres"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
	"github.com/thanh2k4/Chat-app/pkg/database/redis"
	auth2 "github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	// Load the config
	cfg, err := config.LoadConfig("cmd/auth/config/config.yml")
	if err != nil {
		log.Fatalf("❗ Failed to load config: %v", err)
	}

	// Connect to Postgres
	pgPool, err := postgres.NewPostgresDB(cfg.Database.Postgres)
	if err != nil {
		log.Fatalf("❗ Failed to connect to Postgres: %v", err)
	}
	defer pgPool.Close()
	query := postgres2.New(pgPool)

	// Connect to Redis
	redisPool := redis.NewRedisDB(cfg.Database.Redis)
	defer redisPool.Close()
	log.Println("Connected to Redis successfully 🚀")

	grpcServer := grpc.NewServer()
	authServer := &auth.AuthServer{
		Postgres:    query,
		RedisClient: redisPool,
		Cfg:         *cfg,
	}

	auth2.RegisterAuthServiceServer(grpcServer, authServer)
	listener, err := net.Listen("tcp", ":"+cfg.Server.ServerPort)
	if err != nil {
		log.Fatalf("❗ Failed to listen: %v", err)
	}
	fmt.Println("Auth Service is running on port " + cfg.Server.ServerPort + "...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
