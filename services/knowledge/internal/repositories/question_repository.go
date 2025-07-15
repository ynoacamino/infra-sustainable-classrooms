package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
)

type QuestionRepository struct {
	*knowledgedb.Queries
}

func NewQuestionRepository(pool *pgxpool.Pool) ports.QuestionRepository {
	return &QuestionRepository{
		Queries: knowledgedb.New(pool),
	}
}
