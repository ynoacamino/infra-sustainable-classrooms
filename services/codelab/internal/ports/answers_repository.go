package ports

import (
	"context"

	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

type AnswersRepository interface {
	CheckIfAnswerExists(ctx context.Context, arg codelabdb.CheckIfAnswerExistsParams) (int32, error)
	CreateAnswer(ctx context.Context, arg codelabdb.CreateAnswerParams) error
	GetAnswerByUserAndExercise(ctx context.Context, arg codelabdb.GetAnswerByUserAndExerciseParams) (codelabdb.Answer, error)
	ListAnswersByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Answer, error)
	ListAnswersByUser(ctx context.Context, userID int64) ([]codelabdb.Answer, error)
	UpdateAnswerCompleted(ctx context.Context, arg codelabdb.UpdateAnswerCompletedParams) error
	CountCompletedAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error)
	CountTotalAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error)
}
