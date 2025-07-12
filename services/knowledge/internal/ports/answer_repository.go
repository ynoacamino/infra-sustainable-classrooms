package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type AnswerRepository interface {
	// Simplified answer operations
	CreateAnswerSubmission(ctx context.Context, params knowledgedb.CreateAnswerSubmissionParams) error
	GetAnswersBySubmission(ctx context.Context, submissionID int64) ([]knowledgedb.GetAnswersBySubmissionRow, error)
}
