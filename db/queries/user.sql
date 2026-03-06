-- COMMAND QUERIES

-- name: CreateUser :one
INSERT INTO users (id, full_name, username, password, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1) AS exists;
