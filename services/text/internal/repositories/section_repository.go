package repositories

import (
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
)

type SectionRepository struct {
	*textdb.Queries
}

func NewSectionRepository(db textdb.DBTX) ports.SectionRepository {
	return &SectionRepository{
		Queries: textdb.New(db),
	}
}
