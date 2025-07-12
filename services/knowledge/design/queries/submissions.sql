-- === ESSENTIAL TEST SUBMISSIONS QUERIES ===

-- name: CreateSubmission :one
INSERT INTO test_submissions (
    test_id,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSubmissionById :one
SELECT * FROM test_submissions
WHERE id = $1;

-- name: GetSubmissionByUserAndTest :one
SELECT * FROM test_submissions
WHERE user_id = $1 AND test_id = $2;

-- name: CompleteSubmission :one
UPDATE test_submissions
SET
    submitted_at = NOW(),
    score = $2,
    passed = $3,
    time_taken_minutes = $4,
    is_completed = true,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetUserSubmissions :many
SELECT
    ts.*,
    t.title as test_title,
    tc.name as category_name
FROM test_submissions ts
JOIN tests t ON ts.test_id = t.id
LEFT JOIN test_categories tc ON t.category_id = tc.id
WHERE ts.user_id = $1
AND ts.is_completed = true
ORDER BY ts.submitted_at DESC
LIMIT $2 OFFSET $3;

-- name: GetTestParticipants :many
SELECT
    ts.user_id,
    ts.score,
    ts.passed,
    ts.submitted_at,
    ts.time_taken_minutes
FROM test_submissions ts
WHERE ts.test_id = $1
AND ts.is_completed = true
ORDER BY ts.submitted_at DESC
LIMIT $2 OFFSET $3;

-- name: CountTestParticipants :one
SELECT COUNT(*) FROM test_submissions
WHERE test_id = $1 AND is_completed = true;

-- name: GetSubmissionWithAnswers :one
SELECT
    ts.*,
    t.title as test_title,
    tc.name as category_name,
    json_agg(
        json_build_object(
            'question_id', a.question_id,
            'selected_answer', a.selected_answer,
            'is_correct', a.is_correct,
            'points_earned', a.points_earned,
            'question_text', q.question_text,
            'options', q.options,
            'correct_answer', q.correct_answer,
            'explanation', q.explanation,
            'max_points', q.points
        ) ORDER BY q.question_order
    ) as question_results
FROM test_submissions ts
JOIN tests t ON ts.test_id = t.id
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN answer_submissions a ON ts.id = a.submission_id
LEFT JOIN questions q ON a.question_id = q.id
WHERE ts.id = $1
GROUP BY ts.id, t.title, tc.name;

-- name: CheckExistingSubmission :one
SELECT id FROM test_submissions
WHERE user_id = $1 AND test_id = $2 AND is_completed = false;

-- name: CheckUserCompletedTest :one
SELECT EXISTS(SELECT 1 FROM test_submissions WHERE user_id = $1 AND test_id = $2 AND is_completed = true);
