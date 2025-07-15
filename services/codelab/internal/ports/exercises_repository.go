package ports

import (
	"context"

	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

type ExercisesRepository interface {
	CreateExercise(ctx context.Context, arg codelabdb.CreateExerciseParams) error
	GetExerciseById(ctx context.Context, id int64) (codelabdb.Exercise, error)
	GetExerciseToResolveById(ctx context.Context, id int64) (codelabdb.GetExerciseToResolveByIdRow, error)
	ListExercises(ctx context.Context) ([]codelabdb.Exercise, error)
	ListExercisesToResolve(ctx context.Context) ([]codelabdb.ListExercisesToResolveRow, error)
	UpdateExercise(ctx context.Context, arg codelabdb.UpdateExerciseParams) error
	DeleteExercise(ctx context.Context, id int64) error
}
