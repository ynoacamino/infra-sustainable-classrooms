package ports

import (
	"context"

	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type SectionRepository interface {
	GetSection(ctx context.Context, id int64) (textdb.Section, error)
	ListSectionsByCourse(ctx context.Context, courseID int64) ([]textdb.Section, error)
	CreateSection(ctx context.Context, arg textdb.CreateSectionParams) error
	DeleteSection(ctx context.Context, id int64) error
	UpdateSection(ctx context.Context, arg textdb.UpdateSectionParams) error
	GetNextOrderForCourse(ctx context.Context, courseID int64) (int32, error)
}
