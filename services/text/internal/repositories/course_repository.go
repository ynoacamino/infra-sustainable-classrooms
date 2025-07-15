package repositories

import (
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
)

type CourseRepository struct {
	*textdb.Queries
}

func NewCourseRepository(db textdb.DBTX) ports.CourseRepository {
	return &CourseRepository{
		Queries: textdb.New(db),
	}
}
