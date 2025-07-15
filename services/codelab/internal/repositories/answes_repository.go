package repositories

import codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"

type AnswersRepository struct {
	*codelabdb.Queries
}

func NewAnswersRepository(db codelabdb.DBTX) *AnswersRepository {
	return &AnswersRepository{
		Queries: codelabdb.New(db),
	}
}
