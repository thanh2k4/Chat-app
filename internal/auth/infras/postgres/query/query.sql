-- name: GetUserByUsername :one
SELECT * FROM auth.users WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO auth.users ( id ,username, password)
VALUES ($1, $2 , $3)
    RETURNING *;
