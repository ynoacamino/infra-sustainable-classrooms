package ports

import (
	"context"

	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

type AttemptsRepository interface {
	CreateAttempt(ctx context.Context, arg codelabdb.CreateAttemptParams) error
	GetAttempt(ctx context.Context, id int64) (codelabdb.Attempt, error)
	GetAttemptsByAnswer(ctx context.Context, answerID int64) ([]codelabdb.Attempt, error)
	GetAttemptsByUserAndExercise(ctx context.Context, arg codelabdb.GetAttemptsByUserAndExerciseParams) ([]codelabdb.Attempt, error)
	GetLatestAttemptByAnswer(ctx context.Context, answerID int64) (codelabdb.Attempt, error)
	CountAttemptsByAnswer(ctx context.Context, answerID int64) (int64, error)
	CountSuccessfulAttemptsByAnswer(ctx context.Context, answerID int64) (int64, error)
	CountTotalAttemptsByExercise(ctx context.Context, exerciseID int64) (int64, error)
	CountSuccessfulAttemptsByExercise(ctx context.Context, exerciseID int64) (int64, error)
	GetAttemptsWithAnswerInfo(ctx context.Context, exerciseID int64) ([]codelabdb.GetAttemptsWithAnswerInfoRow, error)
	GetUserAttemptsForExercise(ctx context.Context, arg codelabdb.GetUserAttemptsForExerciseParams) ([]codelabdb.Attempt, error)
}
