-- name: CreateAttempt :exec
INSERT INTO attempts (answer_id, code, success)
VALUES ($1, $2, $3);

-- name: GetAttempt :one
SELECT * FROM attempts WHERE id = $1;

-- name: GetAttemptsByAnswer :many
SELECT * FROM attempts 
WHERE answer_id = $1 
ORDER BY created_at DESC;

-- name: GetAttemptsByUserAndExercise :many
SELECT a.* FROM attempts a
JOIN answers ans ON a.answer_id = ans.id
WHERE ans.user_id = $1 AND ans.exercise_id = $2
ORDER BY a.created_at DESC;

-- name: GetLatestAttemptByAnswer :one
SELECT * FROM attempts 
WHERE answer_id = $1 
ORDER BY created_at DESC 
LIMIT 1;

-- name: CountAttemptsByAnswer :one
SELECT COUNT(*) FROM attempts 
WHERE answer_id = $1;

-- name: CountSuccessfulAttemptsByAnswer :one
SELECT COUNT(*) FROM attempts 
WHERE answer_id = $1 AND success = true;

-- name: CountTotalAttemptsByExercise :one
SELECT COUNT(a.*) FROM attempts a
JOIN answers ans ON a.answer_id = ans.id
WHERE ans.exercise_id = $1;

-- name: CountSuccessfulAttemptsByExercise :one
SELECT COUNT(a.*) FROM attempts a
JOIN answers ans ON a.answer_id = ans.id
WHERE ans.exercise_id = $1 AND a.success = true;

-- name: GetAttemptsWithAnswerInfo :many
SELECT 
    a.id,
    a.answer_id,
    a.code,
    a.success,
    a.created_at,
    ans.user_id,
    ans.exercise_id,
    ans.completed as answer_completed
FROM attempts a
JOIN answers ans ON a.answer_id = ans.id
WHERE ans.exercise_id = $1
ORDER BY a.created_at DESC;

-- name: GetUserAttemptsForExercise :many
SELECT 
    a.id,
    a.answer_id,
    a.code,
    a.success,
    a.created_at
FROM attempts a
JOIN answers ans ON a.answer_id = ans.id
WHERE ans.user_id = $1 AND ans.exercise_id = $2
ORDER BY a.created_at DESC;
