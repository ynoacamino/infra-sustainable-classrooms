-- name: GetOrCreateCategory :one
INSERT INTO
    video_categories (name)
VALUES ($1)
ON CONFLICT (name) DO
UPDATE
SET
    name = EXCLUDED.name
RETURNING
    *;

-- name: GetAllCategories :many
SELECT * FROM video_categories ORDER BY name ASC;