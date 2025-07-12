-- name: GetTagByID :one
SELECT * FROM video_tags WHERE id = $1;

-- name: GetTagByName :one
SELECT * FROM video_tags WHERE name = $1;

-- name: GetAllTags :many
SELECT * FROM video_tags ORDER BY name ASC;

-- name: UpdateTag :one
UPDATE video_tags SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteTag :exec
DELETE FROM video_tags WHERE id = $1;

-- name: GetOrCreateTag :one
INSERT INTO
    video_tags (name)
VALUES ($1)
ON CONFLICT (name) DO
UPDATE
SET
    name = EXCLUDED.name
RETURNING
    *;