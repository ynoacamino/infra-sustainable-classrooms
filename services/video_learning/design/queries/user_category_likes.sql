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

-- name: GetUserCategoryLikes :many
SELECT vc.name, ucl.likes
FROM
    video_categories vc
    LEFT JOIN user_category_likes ucl ON vc.id = ucl.category_id
WHERE
    ucl.user_id = $1;

-- name: GetUserVideoLike :one
SELECT *
FROM user_video_likes
WHERE
    user_id = $1
    AND video_id = $2;

-- name: UpsertUserVideoLike :exec
INSERT INTO user_video_likes (user_id, video_id, liked)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, video_id) DO UPDATE
SET liked = $3, updated_at = NOW();

-- name: UpsertUserCategoryLike :exec
INSERT INTO user_category_likes (user_id, category_id, likes)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, category_id) DO UPDATE
SET likes = likes + $3, updated_at = NOW();