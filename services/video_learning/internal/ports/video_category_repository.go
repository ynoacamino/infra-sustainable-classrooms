package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

type VideoCategoryRepository interface {
	// Category CRUD operations
	CreateCategory(ctx context.Context, name string) (videolearningdb.VideoCategory, error)
	GetCategoryByID(ctx context.Context, id int64) (videolearningdb.VideoCategory, error)
	GetCategoryByName(ctx context.Context, name string) (videolearningdb.VideoCategory, error)
	GetAllCategories(ctx context.Context) ([]videolearningdb.VideoCategory, error)
	UpdateCategory(ctx context.Context, params videolearningdb.UpdateCategoryParams) (videolearningdb.VideoCategory, error)
	DeleteCategory(ctx context.Context, id int64) error
}
