package repositories

import (
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
)

type ArticleRepository struct {
	*textdb.Queries
}

func NewArticleRepository(db textdb.DBTX) ports.ArticleRepository {
	return &ArticleRepository{
		Queries: textdb.New(db),
	}
}
