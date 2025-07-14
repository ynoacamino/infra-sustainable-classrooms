-- name: MarkArticleAsCompleted :exec
INSERT INTO article_progress (user_id, article_id, completed_at)
VALUES ($1, $2, NOW())
ON CONFLICT (user_id, article_id) DO NOTHING;

-- name: UnmarkArticleAsCompleted :exec
DELETE FROM article_progress
WHERE user_id = $1 AND article_id = $2;

-- name: CheckArticleCompleted :one
SELECT EXISTS(
    SELECT 1 FROM article_progress
    WHERE user_id = $1 AND article_id = $2
) as completed;

-- name: GetUserCompletedArticles :many
SELECT ap.article_id, ap.completed_at, a.title, a.section_id
FROM article_progress ap
JOIN articles a ON ap.article_id = a.id
WHERE ap.user_id = $1
ORDER BY ap.completed_at DESC;

-- name: GetUserProgressForCourse :many
SELECT 
    a.id as article_id,
    a.title as article_title,
    a.section_id,
    s.title as section_title,
    ap.completed_at,
    CASE WHEN ap.article_id IS NOT NULL THEN true ELSE false END as completed
FROM articles a
JOIN sections s ON a.section_id = s.id
LEFT JOIN article_progress ap ON a.id = ap.article_id AND ap.user_id = $1
WHERE s.course_id = $2
ORDER BY s."order", a.id;

-- name: GetCourseCompletionStats :one
SELECT 
    COUNT(a.id) as total_articles,
    COUNT(ap.article_id) as completed_articles,
    CASE 
        WHEN COUNT(a.id) > 0 THEN 
            ROUND((COUNT(ap.article_id)::numeric / COUNT(a.id)::numeric) * 100, 2)
        ELSE 0 
    END as completion_percentage
FROM articles a
JOIN sections s ON a.section_id = s.id
LEFT JOIN article_progress ap ON a.id = ap.article_id AND ap.user_id = $1
WHERE s.course_id = $2;
