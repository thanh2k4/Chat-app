package chat

import "github.com/thanh2k4/Chat-app/proto/gen"

type ChatServer struct {
	gen.UnimplementedChatServiceServer
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}
