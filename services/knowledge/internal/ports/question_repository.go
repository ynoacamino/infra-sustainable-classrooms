package ports

import (
	"context"

	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

type QuestionRepository interface {
	// Question CRUD operations
	CreateQuestion(ctx context.Context, params knowledgedb.CreateQuestionParams) (knowledgedb.Question, error)
	GetQuestionById(ctx context.Context, id int64) (knowledgedb.Question, error)
	GetQuestionsByTestId(ctx context.Context, testID int64) ([]knowledgedb.Question, error)
	UpdateQuestion(ctx context.Context, params knowledgedb.UpdateQuestionParams) error
	DeleteQuestion(ctx context.Context, id int64) error
	DeleteQuestionsByTestId(ctx context.Context, testID int64) error

	// Questions for test taking (without correct answers)
	GetQuestionForTaking(ctx context.Context, id int64) (knowledgedb.GetQuestionForTakingRow, error)
	GetQuestionsForTaking(ctx context.Context, testID int64) ([]knowledgedb.GetQuestionsForTakingRow, error)

	// Question management
	CountQuestionsByTestId(ctx context.Context, testID int64) (int64, error)
	UpdateQuestionOrder(ctx context.Context, params knowledgedb.UpdateQuestionOrderParams) error

	// Questions with answers (for results)
	GetQuestionWithAnswer(ctx context.Context, params knowledgedb.GetQuestionWithAnswerParams) (knowledgedb.GetQuestionWithAnswerRow, error)
	GetQuestionsWithAnswers(ctx context.Context, params knowledgedb.GetQuestionsWithAnswersParams) ([]knowledgedb.GetQuestionsWithAnswersRow, error)

	// Bulk operations
	BulkCreateQuestions(ctx context.Context, questions []knowledgedb.BulkCreateQuestionsParams) (int64, error)
}
