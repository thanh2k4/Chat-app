package main

import (
	"github.com/thanh2k4/Chat-app/internal/user"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/thanh2k4/Chat-app/pkg/database/postgres"
)

func main() {

	// Load the config
	cfg, err := config.LoadConfig("cmd/user/config/config.yml")
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
	grpcServer := grpc.NewServer()
	userServer := user.NewUserServer()
	gen.RegisterUserServiceServer(grpcServer, userServer)
	listener, err := net.Listen("tcp", ":"+cfg.Server.ServerPort)
	if err != nil {
		log.Fatalf("❗ Failed to listen: %v", err)
	}
	log.Printf("User Service is running on port %s", cfg.Server.ServerPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC for User Service: %v", err)
	}

}
