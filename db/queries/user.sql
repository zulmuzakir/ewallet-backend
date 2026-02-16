-- COMMAND QUERIES

-- name: CreateUser :one
INSERT INTO users (id, full_name, username, password, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
