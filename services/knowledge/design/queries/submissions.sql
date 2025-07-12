-- Submissions queries - simplified

-- name: CreateSubmission :one
INSERT INTO test_submissions (test_id, user_id, score)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetSubmissionById :one
SELECT * FROM test_submissions WHERE id = $1;

-- name: GetUserSubmissions :many
SELECT 
    ts.*,
    t.title as test_title
FROM test_submissions ts
JOIN tests t ON ts.test_id = t.id
WHERE ts.user_id = $1
ORDER BY ts.submitted_at DESC;

-- name: CheckUserCompletedTest :one
SELECT EXISTS(
    SELECT 1 FROM test_submissions 
    WHERE user_id = $1 AND test_id = $2
);
