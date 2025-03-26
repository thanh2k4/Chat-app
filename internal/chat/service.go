package chat

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thanh2k4/Chat-app/internal/chat/infras/postgres"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"time"
)

type ChatServer struct {
	gen.UnimplementedChatServiceServer
	Postgres *postgres.Queries
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}

func (s *ChatServer) SendMessage(ctx context.Context, req *gen.ChatMessage) (*gen.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	chatID, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, err
	}
	senderID, err := uuid.Parse(req.SenderId)
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.Postgres.SendMessage(ctx, postgres.SendMessageParams{
		ID:       pgtype.UUID{Bytes: id, Valid: true},
		ChatID:   pgtype.UUID{Bytes: chatID, Valid: true},
		Content:  pgtype.Text{String: req.Content, Valid: true},
		SenderID: pgtype.UUID{Bytes: senderID, Valid: true},
		Type:     req.Type,
		MediaUrl: pgtype.Text{String: req.MediaUrl, Valid: true},
	})
	return &gen.ChatMessage{
		Id:       req.Id,
		ChatId:   req.ChatId,
		SenderId: req.SenderId,
		Content:  req.Content,
		Type:     req.Type,
		MediaUrl: req.MediaUrl,
		Status:   req.Status,
	}, nil
}

func (s *ChatServer) CreateChat(ctx context.Context, req *gen.ChatCreateRequest) (*gen.Chat, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	creatorId, err := uuid.Parse(req.CreatorId)
	if err != nil {
		return nil, err
	}
	chat, err := s.Postgres.CreateChat(ctx, postgres.CreateChatParams{
		Type:      req.Type,
		CreatorID: pgtype.UUID{Bytes: creatorId, Valid: true},
		Name:      pgtype.Text{String: req.Name, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &gen.Chat{
		Id:        chat.ID.String(),
		Type:      chat.Type,
		CreatorId: chat.CreatorID.String(),
		Name:      chat.Name.String,
	}, nil

}

func (s *ChatServer) GetChat(ctx context.Context, req *gen.ChatGetRequest) (*gen.Chat, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	chatID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	chat, err := s.Postgres.GetChatByID(ctx, pgtype.UUID{Bytes: chatID, Valid: true})
	if err != nil {
		return nil, err
	}
	return &gen.Chat{
		Id:        chat.ID.String(),
		Type:      chat.Type,
		CreatorId: chat.CreatorID.String(),
		Name:      chat.Name.String,
	}, nil

}

func (s *ChatServer) GetChatsByUser(ctx context.Context, req *gen.ChatGetByUserRequest) (*gen.ChatGetByUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	chats, err := s.Postgres.GetChatByUserID(ctx, pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil {
		return nil, err
	}
	var chatList []*gen.Chat
	for _, chatID := range chats {
		chat, err := s.Postgres.GetChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		chatList = append(chatList, &gen.Chat{
			Id:        chat.ID.String(),
			Type:      chat.Type,
			CreatorId: chat.CreatorID.String(),
			Name:      chat.Name.String,
		})
	}
	return &gen.ChatGetByUserResponse{
		Chats: chatList,
	}, nil

}

func (s *ChatServer) GetMessagesByChat(ctx context.Context, req *gen.ChatGetRequest) (*gen.ChatGetMessagesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	chatID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	messages, err := s.Postgres.GetMessagesByChatID(ctx, pgtype.UUID{Bytes: chatID, Valid: true})
	if err != nil {
		return nil, err
	}
	var messageList []*gen.ChatMessage
	for _, message := range messages {
		messageList = append(messageList, &gen.ChatMessage{
			Id:       message.ID.String(),
			ChatId:   message.ChatID.String(),
			SenderId: message.SenderID.String(),
			Content:  message.Content.String,
			Type:     message.Type,
			MediaUrl: message.MediaUrl.String,
			Status:   message.Status,
		})
	}
	return &gen.ChatGetMessagesResponse{
		Messages: messageList,
	}, nil

}
