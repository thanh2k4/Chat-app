package chat

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsGroup   bool      `json:"isGroup"`
	CreatorID uuid.UUID `json:"creatorId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChatMember struct {
	ChatID   uuid.UUID `json:"chatId"`
	UserID   uuid.UUID `json:"userId"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joinedAt"`
}

type Message struct {
	ID        uuid.UUID `json:"id"`
	ChatID    uuid.UUID `json:"chatId"`
	SenderID  uuid.UUID `json:"senderId"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MessageMedia struct {
	ID        uuid.UUID `json:"id"`
	MessageID uuid.UUID `json:"message_id"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
}

type MessageStatus struct {
	MessageID uuid.UUID `json:"message_id"`
	UserID    uuid.UUID `json:"user_id"`
	Status    string    `json:"status"`
	ReadAt    time.Time `json:"read_at"`
}
