-- name: GetUserCategoryLike :one
SELECT *
FROM user_category_likes
WHERE
    user_id = $1
    AND category_id = $2;

-- name: GetUserCategoryLikes :many
SELECT *
FROM user_category_likes
WHERE
    user_id = $1
ORDER BY likes DESC;

-- name: IncrementUserCategoryLike :exec
INSERT INTO
    user_category_likes (user_id, category_id, likes)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, category_id) DO
UPDATE
SET
    likes = user_category_likes.likes + $3;

-- name: DeleteAllUserCategoryLikes :exec
DELETE FROM user_category_likes WHERE user_id = $1;