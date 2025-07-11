package repositories

import (
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/ports"
)

type ProfileRepository struct {
	*profilesdb.Queries
}

func NewProfileRepository(db profilesdb.DBTX) ports.ProfileRepository {
	return &ProfileRepository{
		Queries: profilesdb.New(db),
	}
}
