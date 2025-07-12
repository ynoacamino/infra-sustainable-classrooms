package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

type UserCategoryLikeRepository interface {
	// User category preference operations
	GetUserCategoryLike(ctx context.Context, params videolearningdb.GetUserCategoryLikeParams) (videolearningdb.UserCategoryLike, error)
	GetUserCategoryLikes(ctx context.Context, userID int64) ([]videolearningdb.UserCategoryLike, error)
	IncrementUserCategoryLike(ctx context.Context, params videolearningdb.IncrementUserCategoryLikeParams) error
	DeleteAllUserCategoryLikes(ctx context.Context, userID int64) error
}
