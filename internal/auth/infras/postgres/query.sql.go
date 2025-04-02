// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO auth.users ( id ,username, password)
VALUES ($1, $2 , $3)
    RETURNING id, username, password, created_at, updated_at
`

type CreateUserParams struct {
	ID       pgtype.UUID `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (AuthUser, error) {
	row := q.db.QueryRow(ctx, createUser, arg.ID, arg.Username, arg.Password)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, created_at, updated_at FROM auth.users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAuthUserByID = `-- name: UpdateAuthUserByID :exec
UPDATE auth.users SET username = $2, password = $3
WHERE id = $1
`

type UpdateAuthUserByIDParams struct {
	ID       pgtype.UUID `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
}

func (q *Queries) UpdateAuthUserByID(ctx context.Context, arg UpdateAuthUserByIDParams) error {
	_, err := q.db.Exec(ctx, updateAuthUserByID, arg.ID, arg.Username, arg.Password)
	return err
}
