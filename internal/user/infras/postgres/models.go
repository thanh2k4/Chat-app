// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UserServiceFriend struct {
	ID        pgtype.UUID      `json:"id"`
	UserID    pgtype.UUID      `json:"user_id"`
	FriendID  pgtype.UUID      `json:"friend_id"`
	Status    string           `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type UserServiceUser struct {
	ID        pgtype.UUID      `json:"id"`
	FullName  string           `json:"full_name"`
	Email     pgtype.Text      `json:"email"`
	Phone     pgtype.Text      `json:"phone"`
	Avatar    pgtype.Text      `json:"avatar"`
	Status    string           `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
