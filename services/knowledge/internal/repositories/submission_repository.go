package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
)

type submissionRepository struct {
	*knowledgedb.Queries
}

func NewSubmissionRepository(pool *pgxpool.Pool) ports.SubmissionRepository {
	return &submissionRepository{
		Queries: knowledgedb.New(pool),
	}
}
