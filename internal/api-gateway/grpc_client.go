package api_gateway

import (
	"github.com/bufbuild/protovalidate-go"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	ChatClient gen.ChatServiceClient
	UserClient gen.UserServiceClient
	AuthClient gen.AuthServiceClient
	Validator  protovalidate.Validator
}

func NewClient(cfg config.Config) *Client {
	authConn, err := grpc.NewClient("localhost:"+cfg.GRPC.AuthServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Auth Service: %v", err)
	}

	userConn, err := grpc.NewClient("localhost:"+cfg.GRPC.UserServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to User Service: %v", err)
	}

	chatConn, err := grpc.NewClient("localhost:"+cfg.GRPC.ChatServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
