package ports

import (
	"context"

	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
)

// VideoService defines the business logic operations for video management
type VideoService interface {
	// Video search and discovery
	SearchVideos(ctx context.Context, payload *videolearning.SearchVideosPayload) (*videolearning.VideoList, error)
	GetRecommendations(ctx context.Context, payload *videolearning.GetRecommendationsPayload) (*videolearning.VideoList, error)
	GetVideo(ctx context.Context, payload *videolearning.GetVideoPayload) (*videolearning.VideoDetails, error)
	GetVideosByCategory(ctx context.Context, payload *videolearning.GetVideosByCategoryPayload) (*videolearning.VideoList, error)
	GetSimilarVideos(ctx context.Context, payload *videolearning.GetSimilarVideosPayload) (*videolearning.VideoList, error)

	// Video management
	GetOwnVideos(ctx context.Context, payload *videolearning.GetOwnVideosPayload) ([]*videolearning.OwnVideo, error)
	DeleteVideo(ctx context.Context, payload *videolearning.DeleteVideoPayload) (*videolearning.SimpleResponse, error)
	UpdateVideo(ctx context.Context, payload *videolearning.UpdateVideoPayload) (*videolearning.SimpleResponse, error)

	// Video interactions
	ToggleVideoLike(ctx context.Context, payload *videolearning.ToggleVideoLikePayload) (*videolearning.SimpleResponse, error)
	IncrementViews(ctx context.Context, payload *videolearning.IncrementViewsPayload) (*videolearning.SimpleResponse, error)
}
