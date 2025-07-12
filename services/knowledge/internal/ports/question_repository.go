package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type QuestionRepository interface {
	// Simplified Question CRUD operations
	CreateQuestion(ctx context.Context, params knowledgedb.CreateQuestionParams) error
	GetQuestionById(ctx context.Context, id int64) (knowledgedb.Question, error)
	GetQuestionsByTestId(ctx context.Context, testID int64) ([]knowledgedb.Question, error)
	UpdateQuestion(ctx context.Context, params knowledgedb.UpdateQuestionParams) error
	DeleteQuestion(ctx context.Context, id int64) error
}
