-- name: CreateCategory :one
INSERT INTO video_categories (name)
VALUES ($1)
RETURNING *;

-- name: GetCategoryByID :one
SELECT * FROM video_categories
WHERE id = $1;

-- name: GetCategoryByName :one
SELECT * FROM video_categories
WHERE name = $1;

-- name: GetAllCategories :many
SELECT * FROM video_categories
ORDER BY name ASC;

-- name: UpdateCategory :one
UPDATE video_categories
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM video_categories
WHERE id = $1;
