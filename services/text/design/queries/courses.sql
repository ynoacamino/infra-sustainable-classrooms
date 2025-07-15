-- name: CreateCourse :exec
INSERT INTO courses (title, description, image_url)
VALUES ($1, $2, $3);

-- name: GetCourse :one
SELECT * FROM courses WHERE id = $1;

-- name: ListCourses :many
SELECT * FROM courses ORDER BY created_at ASC;

-- name: UpdateCourse :exec
UPDATE courses
SET title = COALESCE($2, title),
    description = COALESCE($3, description),
    image_url = COALESCE($4, image_url),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE id = $1;
