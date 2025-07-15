-- name: CreateTest :exec
INSERT INTO tests (input, output, public, exercise_id)
VALUES ($1, $2, $3, $4);

-- name: GetTestsByExercise :many
SELECT * FROM tests 
WHERE exercise_id = $1 
ORDER BY created_at;

-- name: GetPublicTestsByExercise :many
SELECT * FROM tests 
WHERE exercise_id = $1 AND public = true 
ORDER BY created_at;

-- name: GetHiddenTestsByExercise :many
SELECT * FROM tests 
WHERE exercise_id = $1 AND public = false 
ORDER BY created_at;

-- name: UpdateTest :exec
UPDATE tests SET
    input = $2,
    output = $3,
    public = $4,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteTest :exec
DELETE FROM tests WHERE id = $1;

-- name: DeleteTestsByExercise :exec
DELETE FROM tests WHERE exercise_id = $1;
