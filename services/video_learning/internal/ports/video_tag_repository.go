package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// VideoTagRepository define las operaciones de persistencia para etiquetas de videos
type VideoTagRepository interface {
	// Tag CRUD operations
	GetOrCreateTag(ctx context.Context, name string) (videolearningdb.VideoTag, error)
	GetTagByName(ctx context.Context, name string) (videolearningdb.VideoTag, error)
	GetAllTags(ctx context.Context) ([]videolearningdb.VideoTag, error)
	GetTagsByVideoID(ctx context.Context, videoID int64) ([]videolearningdb.VideoTag, error)
}
