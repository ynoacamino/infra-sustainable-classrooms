package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
)

type TestRepository struct {
	*knowledgedb.Queries
}

func NewTestRepository(pool *pgxpool.Pool) ports.TestRepository {
	return &TestRepository{
		Queries: knowledgedb.New(pool),
	}
}
