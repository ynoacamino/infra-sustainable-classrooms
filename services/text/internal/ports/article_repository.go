package ports

import (
	"context"

	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type ArticleRepository interface {
	GetArticle(ctx context.Context, id int64) (textdb.Article, error)
	ListArticlesBySection(ctx context.Context, sectionID int64) ([]textdb.Article, error)
	CreateArticle(ctx context.Context, arg textdb.CreateArticleParams) error
	DeleteArticle(ctx context.Context, id int64) error
	UpdateArticle(ctx context.Context, arg textdb.UpdateArticleParams) error
}
