package api_gateway

import (
	"context"
	"github.com/thanh2k4/Chat-app/internal/chat/infras/postgres"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	ChatClient gen.ChatServiceClient
}

func NewClient(chatAddress string) *Client {
	conn, err := grpc.NewClient(chatAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Chat Service: %v", err)
	}

	return &Client{ChatClient: gen.NewChatServiceClient(conn)}
}
func (c *Client) SendMessage(msg *postgres.ChatMessage) (*gen.ChatMessage, error) {
	message := &gen.ChatMessage{
		ChatId:   msg.ChatID.String(),
		SenderId: msg.SenderID.String(),
		Content:  msg.Content.String,
		Type:     msg.Type,
		MediaUrl: msg.MediaUrl.String,
		Status:   msg.Status,
	}
	return c.ChatClient.SendMessage(context.Background(), message)
}
