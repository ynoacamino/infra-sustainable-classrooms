package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// VideoCategoryRepository define las operaciones de persistencia para categorías de videos
type VideoCategoryRepository interface {
	// Category CRUD operations
	GetOrCreateCategory(ctx context.Context, name string) (videolearningdb.VideoCategory, error)
	GetAllCategories(ctx context.Context) ([]videolearningdb.VideoCategory, error)
}
