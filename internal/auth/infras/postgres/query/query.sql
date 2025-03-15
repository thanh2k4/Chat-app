-- name: GetUserByUsername :one

SELECT * FROM auth WHERE username = $1 LIMIT 1;

-- name: CreateUser :one

INSERT INTO auth (id ,username, password)
VALUES ($1, $2, $3)
RETURNING *;