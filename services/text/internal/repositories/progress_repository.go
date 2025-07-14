package repositories

import (
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
)

type ProgressRepository struct {
	*textdb.Queries
}

func NewProgressRepository(db textdb.DBTX) ports.ProgressRepository {
	return &ProgressRepository{
		Queries: textdb.New(db),
	}
}