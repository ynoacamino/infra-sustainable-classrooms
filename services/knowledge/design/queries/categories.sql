-- name: CreateCategory :exec
INSERT INTO test_categories (
    name,
    description
) VALUES (
    $1, $2);

-- name: GetCategoryById :one
SELECT * FROM test_categories
WHERE id = $1;

-- name: GetAllCategories :many
SELECT * FROM test_categories
ORDER BY name ASC;

-- name: UpdateCategory :one
UPDATE test_categories
SET
    name = COALESCE($2, name),
    description = COALESCE($3, description)
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM test_categories
WHERE id = $1;


-- name: CountTestsByCategory :one
SELECT COUNT(*) FROM tests
WHERE category_id = $1;

-- name: GetCategoriesWithTestCounts :many
SELECT
    tc.*,
    COUNT(t.id) as test_count,
    COUNT(CASE WHEN t.is_active THEN 1 END) as active_test_count
FROM test_categories tc
LEFT JOIN tests t ON tc.id = t.category_id
GROUP BY tc.id, tc.name, tc.description, tc.created_at
ORDER BY tc.name ASC;
