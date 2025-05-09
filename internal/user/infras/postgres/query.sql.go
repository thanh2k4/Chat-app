// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const acceptFriendRequest = `-- name: AcceptFriendRequest :exec
UPDATE user_service.friends
SET status = 'accepted'
WHERE user_id = $2 AND friend_id = $1
`

type AcceptFriendRequestParams struct {
	FriendID pgtype.UUID `json:"friend_id"`
	UserID   pgtype.UUID `json:"user_id"`
}

func (q *Queries) AcceptFriendRequest(ctx context.Context, arg AcceptFriendRequestParams) error {
	_, err := q.db.Exec(ctx, acceptFriendRequest, arg.FriendID, arg.UserID)
	return err
}

const checkFriendship = `-- name: CheckFriendship :one
SELECT status FROM user_service.friends
WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)
    LIMIT 1
`

type CheckFriendshipParams struct {
	UserID   pgtype.UUID `json:"user_id"`
	FriendID pgtype.UUID `json:"friend_id"`
}

func (q *Queries) CheckFriendship(ctx context.Context, arg CheckFriendshipParams) (string, error) {
	row := q.db.QueryRow(ctx, checkFriendship, arg.UserID, arg.FriendID)
	var status string
	err := row.Scan(&status)
	return status, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO user_service.users (id,full_name, email, phone, avatar, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, full_name, email, phone, avatar, status, created_at, updated_at
`

type CreateUserParams struct {
	ID       pgtype.UUID `json:"id"`
	FullName string      `json:"full_name"`
	Email    pgtype.Text `json:"email"`
	Phone    pgtype.Text `json:"phone"`
	Avatar   pgtype.Text `json:"avatar"`
	Status   string      `json:"status"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (UserServiceUser, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Phone,
		arg.Avatar,
		arg.Status,
	)
	var i UserServiceUser
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFriendByID = `-- name: DeleteFriendByID :exec
DELETE FROM user_service.friends
WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)
`

type DeleteFriendByIDParams struct {
	UserID   pgtype.UUID `json:"user_id"`
	FriendID pgtype.UUID `json:"friend_id"`
}

func (q *Queries) DeleteFriendByID(ctx context.Context, arg DeleteFriendByIDParams) error {
	_, err := q.db.Exec(ctx, deleteFriendByID, arg.UserID, arg.FriendID)
	return err
}

const getFriendByUserID = `-- name: GetFriendByUserID :many
SELECT id, user_id, friend_id, status, created_at FROM user_service.friends WHERE user_id = $1 OR friend_id = $1 ORDER BY created_at DESC
`

func (q *Queries) GetFriendByUserID(ctx context.Context, userID pgtype.UUID) ([]UserServiceFriend, error) {
	rows, err := q.db.Query(ctx, getFriendByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserServiceFriend
	for rows.Next() {
		var i UserServiceFriend
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FriendID,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, full_name, email, phone, avatar, status, created_at, updated_at FROM user_service.users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email pgtype.Text) (UserServiceUser, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i UserServiceUser
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, full_name, email, phone, avatar, status, created_at, updated_at FROM user_service.users WHERE phone = $1 LIMIT 1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone pgtype.Text) (UserServiceUser, error) {
	row := q.db.QueryRow(ctx, getUserByPhone, phone)
	var i UserServiceUser
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUserID = `-- name: GetUserByUserID :one
SELECT id, full_name, email, phone, avatar, status, created_at, updated_at FROM user_service.users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByUserID(ctx context.Context, id pgtype.UUID) (UserServiceUser, error) {
	row := q.db.QueryRow(ctx, getUserByUserID, id)
	var i UserServiceUser
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const makeFriend = `-- name: MakeFriend :one
INSERT INTO user_service.friends (id, user_id, friend_id, status) VALUES ($1, $2, $3, $4) RETURNING id, user_id, friend_id, status, created_at
`

type MakeFriendParams struct {
	ID       pgtype.UUID `json:"id"`
	UserID   pgtype.UUID `json:"user_id"`
	FriendID pgtype.UUID `json:"friend_id"`
	Status   string      `json:"status"`
}

func (q *Queries) MakeFriend(ctx context.Context, arg MakeFriendParams) (UserServiceFriend, error) {
	row := q.db.QueryRow(ctx, makeFriend,
		arg.ID,
		arg.UserID,
		arg.FriendID,
		arg.Status,
	)
	var i UserServiceFriend
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FriendID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const rejectFriendRequest = `-- name: RejectFriendRequest :exec
DELETE FROM user_service.friends
WHERE user_id = $2 AND friend_id = $1
`

type RejectFriendRequestParams struct {
	FriendID pgtype.UUID `json:"friend_id"`
	UserID   pgtype.UUID `json:"user_id"`
}

func (q *Queries) RejectFriendRequest(ctx context.Context, arg RejectFriendRequestParams) error {
	_, err := q.db.Exec(ctx, rejectFriendRequest, arg.FriendID, arg.UserID)
	return err
}

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE user_service.users SET full_name = $2, email = $3, phone = $4, avatar = $5, status = $6, updated_at = NOW() WHERE id = $1 RETURNING id, full_name, email, phone, avatar, status, created_at, updated_at
`

type UpdateUserByIDParams struct {
	ID       pgtype.UUID `json:"id"`
	FullName string      `json:"full_name"`
	Email    pgtype.Text `json:"email"`
	Phone    pgtype.Text `json:"phone"`
	Avatar   pgtype.Text `json:"avatar"`
	Status   string      `json:"status"`
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (UserServiceUser, error) {
	row := q.db.QueryRow(ctx, updateUserByID,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Phone,
		arg.Avatar,
		arg.Status,
	)
	var i UserServiceUser
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
