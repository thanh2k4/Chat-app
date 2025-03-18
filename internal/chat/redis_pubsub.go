package chat

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"github.com/thanh2k4/Chat-app/internal/chat/infras/postgres"
)

func PublishMessage(redisClient *redis.Client, msg postgres.ChatMessage) error {
	data, _ := json.Marshal(msg)
	return redisClient.Publish(context.Background(), "chat_message:"+msg.ChatID.String(), data).Err()
}

func SubscribeMessages(redisClient *redis.Client, manager *Manager, chatID pgtype.UUID) error {
	ctx := context.Background()
	pubsub := redisClient.Subscribe(ctx, "chat_message:"+chatID.String())
	ch := pubsub.Channel()
	for msg := range ch {
		var message postgres.ChatMessage
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			return err
		}
		manager.SendToUser(ctx, message.ChatID, message.Content.String)
	}
	return nil
}
