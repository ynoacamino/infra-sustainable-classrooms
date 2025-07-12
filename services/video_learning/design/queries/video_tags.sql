-- name: GetTagByName :one
SELECT * FROM video_tags WHERE name = $1;

-- name: GetAllTags :many
SELECT * FROM video_tags ORDER BY name ASC;

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