-- name: CreateUser :one
INSERT INTO users (
    username, role
) VALUES (
    $1, $2
) RETURNING *;