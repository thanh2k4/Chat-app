package api_gateway

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

type Client struct {
	ChatClient gen.ChatServiceClient
	UserClient gen.UserServiceClient
	AuthClient gen.AuthServiceClient
	Validator  protovalidate.Validator
}

func NewClient(cfg config.Config) *Client {
	// Connect to the Auth Service
	authHost := os.Getenv("GRPC_AUTH_HOST")
	if authHost == "" {
		authHost = "localhost"
	}
	authConn, err := grpc.NewClient(authHost+":"+cfg.GRPC.AuthServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Auth Service: %v", err)
	}

	// Connect to the User Service
	userHost := os.Getenv("GRPC_USER_HOST")
	if userHost == "" {
		userHost = "localhost"
	}
	userConn, err := grpc.NewClient(userHost+":"+cfg.GRPC.UserServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to User Service: %v", err)
	}

	// Connect to the Chat Service
	chatHost := os.Getenv("GRPC_CHAT_HOST")
	if chatHost == "" {
		chatHost = "localhost"
	}
	chatConn, err := grpc.NewClient(userHost+":"+cfg.GRPC.ChatServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Chat Service: %v", err)
	}

	v, _ := protovalidate.New()

	return &Client{
		UserClient: gen.NewUserServiceClient(userConn),
		ChatClient: gen.NewChatServiceClient(chatConn),
		AuthClient: gen.NewAuthServiceClient(authConn),
		Validator:  v,
	}
}
