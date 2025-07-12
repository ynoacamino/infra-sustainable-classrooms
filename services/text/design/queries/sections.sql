-- name: CreateSection :exec
INSERT INTO sections (course_id, title, description, "order")
VALUES ($1, $2, $3, $4);

-- name: GetSection :one
SELECT * FROM sections WHERE id = $1;

-- name: ListSectionsByCourse :many
SELECT * FROM sections WHERE course_id = $1 ORDER BY "order" ASC;

-- name: UpdateSection :exec
UPDATE sections
SET title = COALESCE($2, title),
    description = COALESCE($3, description),
    "order" = COALESCE($4, "order"),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteSection :exec
DELETE FROM sections WHERE id = $1;

-- name: GetNextOrderForCourse :one
SELECT COALESCE(MAX("order"), 0) + 1 FROM sections WHERE course_id = $1;
