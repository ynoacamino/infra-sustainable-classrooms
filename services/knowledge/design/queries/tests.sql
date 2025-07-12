-- name: CreateTest :exec
INSERT INTO tests (
    title,
    description,
    category_id,
    difficulty_level,
    duration_minutes,
    passing_score,
    is_active,
    expires_at,
    instructions,
    created_by
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
);

-- name: GetTestById :one
SELECT
    t.*,
    tc.name as category_name,
    COUNT(q.id) as total_questions
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
WHERE t.id = $1
GROUP BY t.id, tc.name;

-- name: GetTestsByCreator :many
SELECT
    t.*,
    tc.name as category_name,
    COUNT(q.id) as total_questions
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
WHERE t.created_by = $1
GROUP BY t.id, tc.name
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetAvailableTests :many
SELECT
    t.*,
    tc.name as category_name,
    COUNT(q.id) as total_questions
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
WHERE t.is_active = true AND (t.expires_at IS NULL OR t.expires_at > NOW())
GROUP BY t.id, tc.name
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateTest :exec
UPDATE tests
SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    category_id = COALESCE($4, category_id),
    duration_minutes = COALESCE($5, duration_minutes),
    passing_score = COALESCE($6, passing_score),
    is_active = COALESCE($7, is_active),
    expires_at = COALESCE($8, expires_at),
    instructions = COALESCE($9, instructions),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteTest :exec
DELETE FROM tests
WHERE id = $1 AND created_by = $2;

-- name: GetTestWithQuestions :one
SELECT
    t.*,
    tc.name as category_name,
    COUNT(q.id) as total_questions
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
WHERE t.id = $1
GROUP BY t.id, tc.name;



-- name: CountTestsByCreator :one
SELECT COUNT(*) FROM tests
WHERE created_by = $1;

-- name: GetTestsWithSubmissionCounts :many
SELECT
    t.*,
    tc.name as category_name,
    COUNT(DISTINCT q.id) as total_questions,
    COUNT(DISTINCT ts.user_id) as participant_count,
    COUNT(CASE WHEN ts.is_completed THEN 1 END) as completed_count
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
LEFT JOIN test_submissions ts ON t.id = ts.test_id
WHERE t.created_by = $1
GROUP BY t.id, tc.name
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetTestPreview :one
SELECT
    t.*,
    tc.name as category_name,
    COUNT(DISTINCT q.id) as total_questions,
    COUNT(DISTINCT ts.user_id) as total_participants
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
LEFT JOIN test_submissions ts ON t.id = ts.test_id
WHERE t.id = $1
GROUP BY t.id, tc.name;

-- name: CheckTestAccess :one
SELECT
    t.id,
    t.is_active,
    t.expires_at
FROM tests t
LEFT JOIN test_submissions ts ON t.id = ts.test_id
WHERE t.id = $1
GROUP BY t.id, t.is_active, t.expires_at;

-- name: GetTestsByCategory :many
SELECT
    t.*,
    tc.name as category_name,
    COUNT(q.id) as total_questions
FROM tests t
LEFT JOIN test_categories tc ON t.category_id = tc.id
LEFT JOIN questions q ON t.id = q.test_id
WHERE t.category_id = $1
AND t.is_active = true
AND (t.expires_at IS NULL OR t.expires_at > NOW())
GROUP BY t.id, tc.name
ORDER BY t.created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateTestActiveStatus :exec
UPDATE tests
SET is_active = $2, updated_at = NOW()
WHERE id = $1;

-- name: GetExpiredTests :many
SELECT * FROM tests
WHERE expires_at < NOW() AND is_active = true;

-- name: GetPopularTests :many
SELECT
    t.*,
    COUNT(DISTINCT ts.user_id) as participant_count,
    COUNT(q.id) as total_questions
FROM tests t
LEFT JOIN test_submissions ts ON t.id = ts.test_id
LEFT JOIN questions q ON t.id = q.test_id
WHERE t.is_active = true
AND (t.expires_at IS NULL OR t.expires_at > NOW())
GROUP BY t.id
HAVING COUNT(DISTINCT ts.user_id) > 0
ORDER BY COUNT(DISTINCT ts.user_id) DESC, t.created_at DESC
LIMIT $1 OFFSET $2;
