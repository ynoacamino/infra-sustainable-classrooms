-- name: GetTagByName :one
SELECT * FROM video_tags WHERE name = $1;

-- name: GetAllTags :many
SELECT * FROM video_tags ORDER BY name ASC;

-- name: GetTagsByVideoID :many
SELECT vt.*
FROM
    video_tags vt
    JOIN video_video_tags vvt ON vt.id = vvt.tag_id
WHERE
    vvt.video_id = $1;

-- name: GetTagById :one
SELECT * FROM video_tags WHERE id = $1;

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