package user

import (
	"github.com/google/uuid"
	"time"
)

type UserProfile struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Avatar   string    `json:"avatar"`
	Status   string    `json:"status"`
}

type Friend struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FriendID  uuid.UUID `json:"friend_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStatus struct {
	UserID   uuid.UUID `json:"user_id"`
	Status   string    `json:"status"`
	LastSeen time.Time `json:"last_seen"`
}
