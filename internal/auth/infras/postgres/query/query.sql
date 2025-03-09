-- name:GetUserByUsername
SELECT * FROM users WHERE username = $1;