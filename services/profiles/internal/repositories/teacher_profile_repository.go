package repositories

import (
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/ports"
)

type TeacherProfileRepository struct {
	*profilesdb.Queries
}

func NewTeacherProfileRepository(db profilesdb.DBTX) ports.TeacherProfileRepository {
	return &TeacherProfileRepository{
		Queries: profilesdb.New(db),
	}
}
