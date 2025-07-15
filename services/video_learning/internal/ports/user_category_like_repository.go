package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// UserCategoryLikeRepository define las operaciones de persistencia para los likes de usuarios en categorías
type UserCategoryLikeRepository interface {
	// User category like operations
	IncrementUserCategoryLike(ctx context.Context, params videolearningdb.IncrementUserCategoryLikeParams) error
	UpsertUserCategoryLike(ctx context.Context, params videolearningdb.UpsertUserCategoryLikeParams) error
	DeleteAllUserCategoryLikes(ctx context.Context, userID int64) error
	GetUserCategoryLikes(ctx context.Context, userID int64) ([]videolearningdb.GetUserCategoryLikesRow, error)

	// User video like operations
	GetUserVideoLike(ctx context.Context, params videolearningdb.GetUserVideoLikeParams) (videolearningdb.UserVideoLike, error)
	UpsertUserVideoLike(ctx context.Context, params videolearningdb.UpsertUserVideoLikeParams) error
}
