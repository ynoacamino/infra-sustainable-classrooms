package repositories

import codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"

type ExercisesRepository struct {
	*codelabdb.Queries
}

func NewExercisesRepository(db codelabdb.DBTX) *ExercisesRepository {
	return &ExercisesRepository{
		Queries: codelabdb.New(db),
	}
}
