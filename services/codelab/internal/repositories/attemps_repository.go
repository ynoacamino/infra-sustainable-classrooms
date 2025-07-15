package repositories

import codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"

type AttempsRepository struct {
	*codelabdb.Queries
}

func NewAttemptsRepository(db codelabdb.DBTX) *AttempsRepository {
	return &AttempsRepository{
		Queries: codelabdb.New(db),
	}
}
