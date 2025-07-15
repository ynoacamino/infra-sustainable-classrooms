-- name: CreateExercise :exec
INSERT INTO exercises (
    title, description, initial_code, solution, difficulty, created_by
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetExerciseById :one
SELECT * FROM exercises WHERE id = $1;

-- name: GetExerciseToResolveById :one
SELECT 
    id,
    title,
    description,
    initial_code,
    difficulty,
    created_by,
    created_at,
    updated_at
FROM exercises WHERE id = $1;

-- name: ListExercises :many
SELECT * FROM exercises ORDER BY created_at DESC;

-- name: ListExercisesToResolve :many
SELECT
    id,
    title,
    description,
    initial_code,
    difficulty,
    created_by,
    created_at,
    updated_at
FROM exercises
ORDER BY created_at DESC;

-- name: UpdateExercise :exec
UPDATE exercises SET
    title = $2,
    description = $3,
    initial_code = $4,
    solution = $5,
    difficulty = $6,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteExercise :exec
DELETE FROM exercises WHERE id = $1;