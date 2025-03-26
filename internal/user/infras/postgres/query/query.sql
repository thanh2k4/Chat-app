-- name: GetUserByUserID :one
SELECT * FROM user_service.users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM user_service.users WHERE email = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM user_service.users WHERE phone = $1 LIMIT 1;

-- name: UpdateUserByID :one
UPDATE user_service.users SET full_name = $2, email = $3, phone = $4, avatar = $5, status = $6, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: CreateUser :one
INSERT INTO user_service.users (id,full_name, email, phone, avatar, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: MakeFriend :one
INSERT INTO user_service.friends (id, user_id, friend_id, status) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetFriendByUserID :many
SELECT * FROM user_service.friends WHERE user_id = $1 OR friend_id = $1 ORDER BY created_at DESC;

-- name: DeleteFriendByID :exec
DELETE FROM user_service.friends
WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1);

-- name: AcceptFriendRequest :exec
UPDATE user_service.friends
SET status = 'accepted'
WHERE user_id = $2 AND friend_id = $1;

-- name: RejectFriendRequest :exec
DELETE FROM user_service.friends
WHERE user_id = $2 AND friend_id = $1;

-- name: CheckFriendship :one
SELECT status FROM user_service.friends
WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)
    LIMIT 1;

