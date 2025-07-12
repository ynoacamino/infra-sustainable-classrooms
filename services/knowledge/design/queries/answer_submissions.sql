-- === ESSENTIAL ANSWER SUBMISSIONS QUERIES ===

-- name: CreateAnswerSubmission :one
INSERT INTO answer_submissions (
    submission_id,
    question_id,
    selected_answer,
    is_correct,
    points_earned
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: BulkCreateAnswerSubmissions :copyfrom
INSERT INTO answer_submissions (
    submission_id,
    question_id,
    selected_answer,
    is_correct,
    points_earned
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: GetAnswerSubmissions :many
SELECT * FROM answer_submissions
WHERE submission_id = $1
ORDER BY answered_at ASC;

-- name: GetAnswersWithQuestions :many
SELECT
    a.*,
    q.question_text,
    q.options,
    q.correct_answer,
    q.explanation,
    q.points as max_points,
    q.question_order
FROM answer_submissions a
JOIN questions q ON a.question_id = q.id
WHERE a.submission_id = $1
ORDER BY q.question_order ASC;

-- name: ValidateAnswerSubmission :one
SELECT
    q.correct_answer,
    q.points,
    CASE WHEN q.correct_answer = $2 THEN true ELSE false END as is_correct,
    CASE WHEN q.correct_answer = $2 THEN q.points ELSE 0 END as points_earned
FROM questions q
WHERE q.id = $1;

-- name: CountCorrectAnswers :one
SELECT COUNT(*) FROM answer_submissions
WHERE submission_id = $1 AND is_correct = true;

-- name: GetTotalPointsEarned :one
SELECT COALESCE(SUM(points_earned), 0) as total_points
FROM answer_submissions
WHERE submission_id = $1;

-- name: IsSubmissionComplete :one
SELECT
    (SELECT COUNT(*) FROM questions q JOIN test_submissions ts ON q.test_id = ts.test_id WHERE ts.id = $1) as expected_answers,
    (SELECT COUNT(*) FROM answer_submissions WHERE submission_id = $1) as actual_answers,
    (SELECT COUNT(*) FROM questions q JOIN test_submissions ts ON q.test_id = ts.test_id WHERE ts.id = $1) =
    (SELECT COUNT(*) FROM answer_submissions WHERE submission_id = $1) as is_complete;
