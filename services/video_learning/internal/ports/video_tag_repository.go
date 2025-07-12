package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

type VideoTagRepository interface {
	// Tag CRUD operations
	GetOrCreateTag(ctx context.Context, name string) (videolearningdb.VideoTag, error)
	GetTagByID(ctx context.Context, id int64) (videolearningdb.VideoTag, error)
	GetTagByName(ctx context.Context, name string) (videolearningdb.VideoTag, error)
	GetAllTags(ctx context.Context) ([]videolearningdb.VideoTag, error)
	UpdateTag(ctx context.Context, params videolearningdb.UpdateTagParams) (videolearningdb.VideoTag, error)
	DeleteTag(ctx context.Context, id int64) error
}
