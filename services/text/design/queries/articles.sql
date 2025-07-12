-- name: CreateArticle :exec
INSERT INTO articles (section_id, title, content)
VALUES ($1, $2, $3);

-- name: GetArticle :one
SELECT * FROM articles WHERE id = $1;

-- name: ListArticlesBySection :many
SELECT * FROM articles WHERE section_id = $1 ORDER BY created_at ASC;

-- name: UpdateArticle :exec
UPDATE articles
SET title = COALESCE($2, title),
    content = COALESCE($3, content),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteArticle :exec
DELETE FROM articles WHERE id = $1;
