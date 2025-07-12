-- Tests queries - simplified

-- name: CreateTest :exec
INSERT INTO tests (title, created_by) VALUES ($1, $2);

-- name: GetTestById :one
SELECT * FROM tests WHERE id = $1;

-- name: GetMyTests :many
SELECT * FROM tests WHERE created_by = $1 ORDER BY created_at DESC;

-- name: GetAvailableTests :many
SELECT t.*, 
       (SELECT COUNT(*) FROM questions WHERE test_id = t.id) as question_count
FROM tests t 
WHERE t.created_by != $1 
  AND NOT EXISTS (SELECT 1 FROM test_submissions WHERE test_id = t.id AND user_id = $1)
ORDER BY t.created_at DESC;

-- name: UpdateTest :exec
UPDATE tests SET title = $2 WHERE id = $1;

-- name: DeleteTest :exec
DELETE FROM tests WHERE id = $1;
