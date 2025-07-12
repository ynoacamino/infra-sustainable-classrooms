package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
)

type answerRepository struct {
	queries *knowledgedb.Queries
}

func NewAnswerRepository(pool *pgxpool.Pool) ports.AnswerRepository {
	return &answerRepository{
		queries: knowledgedb.New(pool),
	}
}

func (r *answerRepository) CreateAnswerSubmission(ctx context.Context, params knowledgedb.CreateAnswerSubmissionParams) error {
	return r.queries.CreateAnswerSubmission(ctx, params)
}

func (r *answerRepository) GetAnswersBySubmission(ctx context.Context, submissionID int64) ([]knowledgedb.GetAnswersBySubmissionRow, error) {
	return r.queries.GetAnswersBySubmission(ctx, submissionID)
}
