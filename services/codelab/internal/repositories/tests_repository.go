package repositories

import codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"

type TestsRepository struct {
	*codelabdb.Queries
}

func NewTestsRepository(db codelabdb.DBTX) *TestsRepository {
	return &TestsRepository{
		Queries: codelabdb.New(db),
	}
}
