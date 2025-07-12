-- name: CreateVideo :one
INSERT INTO
    video (
        title,
        user_id,
        description,
        views,
        likes,
        video_obj_name,
        thumb_obj_name,
        category_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    )
RETURNING
    *;

-- name: GetVideoByID :one
SELECT v.*, vc.name as category_name
FROM video v
    JOIN video_categories vc ON v.category_id = vc.id
WHERE
    v.id = $1;

-- name: GetVideosByCategory :many
SELECT v.*, vc.name as category_name
FROM video v
    JOIN video_categories vc ON v.category_id = vc.id
WHERE
    v.category_id = $1
ORDER BY (v.likes + v.views) DESC
LIMIT $2
OFFSET
    $3;

-- name: GetRecentVideos :many
SELECT v.*, vc.name as category_name
FROM video v
    JOIN video_categories vc ON v.category_id = vc.id
WHERE
    v.created_at >= NOW() - INTERVAL $1;

-- name: SearchVideos :many
SELECT DISTINCT
    v.*,
    vc.name as category_name
FROM
    video v
    JOIN video_categories vc ON v.category_id = vc.id
    LEFT JOIN video_video_tags vvt ON v.id = vvt.video_id
    LEFT JOIN video_tags vt ON vvt.tag_id = vt.id
WHERE (
        v.title ILIKE '%' || $1 || '%'
        OR vt.name ILIKE '%' || $1 || '%'
    )
    AND (
        $2::bigint IS NULL
        OR v.category_id = $2
    )
ORDER BY (v.likes + v.views) DESC
LIMIT $3
OFFSET
    $4;

-- name: GetSimilarVideos :many
SELECT v.*, vc.name as category_name
FROM video v
    JOIN video_categories vc ON v.category_id = vc.id
WHERE
    v.id != $1
    AND v.category_id = (
        SELECT category_id
        FROM video
        WHERE
            id = $1
    );

-- name: GetVideosByUser :many
SELECT v.*, vc.name as category_name
FROM video v
    JOIN video_categories vc ON v.category_id = vc.id
WHERE
    v.user_id = $1
ORDER BY v.created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: IncrementVideoViews :exec
UPDATE video SET views = views + $2 WHERE id = $1;

-- name: IncrementVideoLikes :exec
UPDATE video SET likes = likes + $2 WHERE id = $1;

-- name: DeleteVideo :exec
DELETE FROM video WHERE id = $1;

-- name: AssignTagToVideo :exec
INSERT INTO
    video_video_tags (video_id, tag_id)
VALUES ($1, $2)
ON CONFLICT (video_id, tag_id) DO NOTHING;

-- name: RemoveTagFromVideo :exec
DELETE FROM video_video_tags WHERE video_id = $1 AND tag_id = $2;