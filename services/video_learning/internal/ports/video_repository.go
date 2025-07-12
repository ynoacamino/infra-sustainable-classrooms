package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// VideoRepository define las operaciones de persistencia para videos
type VideoRepository interface {
	// Video CRUD operations
	CreateVideo(ctx context.Context, params videolearningdb.CreateVideoParams) (videolearningdb.Video, error)
	GetVideoByID(ctx context.Context, id int64) (videolearningdb.GetVideoByIDRow, error)
	UpdateVideo(ctx context.Context, params videolearningdb.UpdateVideoParams) (videolearningdb.Video, error)
	DeleteVideo(ctx context.Context, id int64) error

	// Video search and filtering
	SearchVideos(ctx context.Context, params videolearningdb.SearchVideosParams) ([]videolearningdb.SearchVideosRow, error)
	GetVideosByCategory(ctx context.Context, params videolearningdb.GetVideosByCategoryParams) ([]videolearningdb.GetVideosByCategoryRow, error)
	GetSimilarVideos(ctx context.Context, params videolearningdb.GetSimilarVideosParams) ([]videolearningdb.GetSimilarVideosRow, error)

	// Video interactions
	IncrementVideoViews(ctx context.Context, params videolearningdb.IncrementVideoViewsParams) error
	IncrementVideoLikes(ctx context.Context, params videolearningdb.IncrementVideoLikesParams) error

	// Video tags association
	AssignTagToVideo(ctx context.Context, params videolearningdb.AssignTagToVideoParams) error
	RemoveTagFromVideo(ctx context.Context, params videolearningdb.RemoveTagFromVideoParams) error
}
