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

-- name: GetCategoryById :one
SELECT *
FROM video_categories
WHERE id = $1;

-- name: GetAllCategories :many
SELECT * FROM video_categories ORDER BY name ASC;