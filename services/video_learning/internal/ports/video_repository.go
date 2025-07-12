package ports

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// VideoRepository define las operaciones de persistencia para videos
type VideoRepository interface {
	// Video CRUD operations
	CreateVideo(ctx context.Context, params videolearningdb.CreateVideoParams) (videolearningdb.Video, error)
	GetVideoByID(ctx context.Context, id int64) (videolearningdb.GetVideoByIDRow, error)
	DeleteVideo(ctx context.Context, id int64) error

	// Video search and filtering
	SearchVideos(ctx context.Context, params videolearningdb.SearchVideosParams) ([]videolearningdb.SearchVideosRow, error)
	GetVideosByCategory(ctx context.Context, id int64) ([]videolearningdb.GetVideosByCategoryRow, error)
	GetVideosByUser(ctx context.Context, params videolearningdb.GetVideosByUserParams) ([]videolearningdb.GetVideosByUserRow, error)
	GetSimilarVideos(ctx context.Context, id int64) ([]videolearningdb.GetSimilarVideosRow, error)
	GetRecentVideos(ctx context.Context, interval pgtype.Interval) ([]videolearningdb.GetRecentVideosRow, error)

	// Video interactions
	IncrementVideoViews(ctx context.Context, params videolearningdb.IncrementVideoViewsParams) error
	IncrementVideoLikes(ctx context.Context, params videolearningdb.IncrementVideoLikesParams) error

	// Video tags association
	AssignTagToVideo(ctx context.Context, params videolearningdb.AssignTagToVideoParams) error
	RemoveTagFromVideo(ctx context.Context, params videolearningdb.RemoveTagFromVideoParams) error
}
