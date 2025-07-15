-- name: CreateComment :one
INSERT INTO
    video_comments (
        video_id,
        user_id,
        title,
        content
    )
VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetCommentsForVideo :many
SELECT *
FROM video_comments
WHERE
    video_id = $1
ORDER BY updated_at DESC
LIMIT $2
OFFSET
    $3;

-- name: DeleteComment :exec
DELETE FROM video_comments WHERE id = $1 AND user_id = $2;