package ports

import (
	"context"

	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

type TestsRepository interface {
	CreateTest(ctx context.Context, arg codelabdb.CreateTestParams) error
	GetTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error)
	GetPublicTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error)
	GetHiddenTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error)
	UpdateTest(ctx context.Context, arg codelabdb.UpdateTestParams) error
	DeleteTest(ctx context.Context, id int64) error
	DeleteTestsByExercise(ctx context.Context, exerciseID int64) error
}
