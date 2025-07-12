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

-- name: GetCommentByID :one
SELECT * FROM video_comments WHERE id = $1;

-- name: GetCommentsForVideo :many
SELECT *
FROM video_comments
WHERE
    video_id = $1
ORDER BY updated_at DESC
LIMIT $2
OFFSET
    $3;

-- name: UpdateComment :one
UPDATE video_comments
SET
    title = COALESCE($2, title),
    content = COALESCE($3, content)
WHERE
    id = $1
    AND user_id = $4
RETURNING
    *;

-- name: DeleteComment :exec
DELETE FROM video_comments WHERE id = $1 AND user_id = $2;

-- name: CreateCommentReply :one
INSERT INTO
    video_comment_replies (comment_id, user_id, content)
VALUES ($1, $2, $3)
RETURNING
    *;

-- name: GetCommentReplyByID :one
SELECT * FROM video_comment_replies WHERE id = $1;

-- name: GetRepliesForComment :many
SELECT *
FROM video_comment_replies
WHERE
    comment_id = $1
ORDER BY updated_at ASC;

-- name: GetUserCommentsAndReplies :many
SELECT c.id as id, c.updated_at as updated_at, c.title as title, c.content as content
FROM video_comments c
WHERE
    c.user_id = $1
UNION ALL
SELECT r.id as id, r.updated_at as updated_at, NULL as title, r.content as content
FROM video_comment_replies r
WHERE
    r.user_id = $1
ORDER BY updated_at DESC
LIMIT $2
OFFSET
    $3;

-- name: UpdateCommentReply :one
UPDATE video_comment_replies
SET
    content = $2
WHERE
    id = $1
    AND user_id = $3
RETURNING
    *;

-- name: DeleteCommentReply :exec
DELETE FROM video_comment_replies WHERE id = $1 AND user_id = $2;