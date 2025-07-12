package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
)

type testCategoryRepository struct {
	*knowledgedb.Queries
}

func NewTestCategoryRepository(pool *pgxpool.Pool) ports.TestCategoryRepository {
	return &testCategoryRepository{
		Queries: knowledgedb.New(pool),
	}
}
