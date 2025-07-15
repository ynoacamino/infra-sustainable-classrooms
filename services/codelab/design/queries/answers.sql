-- name: CheckIfAnswerExists :one
SELECT 1 FROM answers
WHERE exercise_id = $1 AND user_id = $2;

-- name: CreateAnswer :exec
INSERT INTO answers (exercise_id, user_id, completed)
VALUES ($1, $2, $3);

-- name: GetAnswerByUserAndExercise :one
SELECT * FROM answers 
WHERE exercise_id = $1 AND user_id = $2;

-- name: ListAnswersByExercise :many
SELECT * FROM answers 
WHERE exercise_id = $1
ORDER BY updated_at DESC;

-- name: ListAnswersByUser :many
SELECT * FROM answers 
WHERE user_id = $1 
ORDER BY updated_at DESC;

-- name: UpdateAnswerCompleted :exec
UPDATE answers SET
    completed = $3,
    updated_at = NOW()
WHERE exercise_id = $1 AND user_id = $2;

-- name: CountCompletedAnswersByExercise :one
SELECT COUNT(*) FROM answers 
WHERE exercise_id = $1 AND completed = true;

-- name: CountTotalAnswersByExercise :one
SELECT COUNT(*) FROM answers 
WHERE exercise_id = $1;
