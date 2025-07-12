-- Questions queries - simplified

-- name: CreateQuestion :exec
INSERT INTO questions (
    test_id,
    question_text,
    option_a,
    option_b,
    option_c,
    option_d,
    correct_answer,
    question_order
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetQuestionsByTestId :many
SELECT * FROM questions
WHERE test_id = $1
ORDER BY question_order ASC;

-- name: GetQuestionById :one
SELECT * FROM questions WHERE id = $1;

-- name: UpdateQuestion :exec
UPDATE questions
SET
    question_text = $2,
    option_a = $3,
    option_b = $4,
    option_c = $5,
    option_d = $6,
    correct_answer = $7
WHERE id = $1;

-- name: DeleteQuestion :exec
DELETE FROM questions WHERE id = $1;
