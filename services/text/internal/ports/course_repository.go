package ports

import (
	"context"

	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type CourseRepository interface {
	GetCourse(ctx context.Context, id int64) (textdb.Course, error)
	ListCourses(ctx context.Context) ([]textdb.Course, error)
	CreateCourse(ctx context.Context, arg textdb.CreateCourseParams) error
	DeleteCourse(ctx context.Context, id int64) error
	UpdateCourse(ctx context.Context, arg textdb.UpdateCourseParams) error
}