-- name: CreateQuestion :one
INSERT INTO questions (
    test_id,
    question_text,
    options,
    correct_answer,
    explanation,
    points,
    question_order
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetQuestionById :one
SELECT * FROM questions
WHERE id = $1;

-- name: GetQuestionsByTestId :many
SELECT * FROM questions
WHERE test_id = $1
ORDER BY question_order ASC;

-- name: UpdateQuestion :exec
UPDATE questions
SET
    question_text = COALESCE($2, question_text),
    options = COALESCE($3, options),
    correct_answer = COALESCE($4, correct_answer),
    explanation = COALESCE($5, explanation),
    points = COALESCE($6, points),
    question_order = COALESCE($7, question_order),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteQuestion :exec
DELETE FROM questions
WHERE id = $1;

-- name: DeleteQuestionsByTestId :exec
DELETE FROM questions
WHERE test_id = $1;

-- name: GetQuestionForTaking :one
SELECT
    id,
    test_id,
    question_text,
    options,
    points,
    question_order
FROM questions
WHERE id = $1;

-- name: GetQuestionsForTaking :many
SELECT
    id,
    test_id,
    question_text,
    options,
    points,
    question_order
FROM questions
WHERE test_id = $1
ORDER BY question_order ASC;

-- name: CountQuestionsByTestId :one
SELECT COUNT(*) FROM questions
WHERE test_id = $1;

-- name: GetQuestionWithAnswer :one
SELECT
    q.*,
    CASE
        WHEN a.selected_answer = q.correct_answer THEN true
        ELSE false
    END as user_was_correct,
    a.selected_answer as user_answer
FROM questions q
LEFT JOIN answer_submissions a ON q.id = a.question_id
WHERE q.id = $1 AND a.submission_id = $2;

-- name: GetQuestionsWithAnswers :many
SELECT
    q.*,
    CASE
        WHEN a.selected_answer = q.correct_answer THEN true
        ELSE false
    END as user_was_correct,
    a.selected_answer as user_answer,
    a.points_earned
FROM questions q
LEFT JOIN answer_submissions a ON q.id = a.question_id
WHERE q.test_id = $1 AND a.submission_id = $2
ORDER BY q.question_order ASC;

-- name: BulkCreateQuestions :copyfrom
INSERT INTO questions (
    test_id,
    question_text,
    options,
    correct_answer,
    explanation,
    points,
    question_order
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: UpdateQuestionOrder :exec
UPDATE questions
SET question_order = $2, updated_at = NOW()
WHERE id = $1;
