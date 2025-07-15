package repositories

import (
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/ports"
)

type StudentProfileRepository struct {
	*profilesdb.Queries
}

func NewStudentProfileRepository(db profilesdb.DBTX) ports.StudentProfileRepository {
	return &StudentProfileRepository{
		Queries: profilesdb.New(db),
	}
}
