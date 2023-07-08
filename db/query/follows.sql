-- name: CreateFollow :one
INSERT INTO follows (
    following_user_id, followed_user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetFollowing :one
SELECT * FROM follows
WHERE following_user_id = $1 LIMIT 1;

-- name: GetFollower :one
SELECT * FROM follows
WHERE followed_user_id = $1 LIMIT 1;

-- name: ListFollowing :many
SELECT * FROM follows
ORDER BY following_user_id
LIMIT $1
OFFSET $2;

-- name: ListFollower :many
SELECT * FROM follows
ORDER BY followed_user_id
LIMIT $1
OFFSET $2;

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE following_user_id = $1;