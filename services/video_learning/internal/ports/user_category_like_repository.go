package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// UserCategoryLikeRepository define las operaciones de persistencia para los likes de usuarios en categorías
type UserCategoryLikeRepository interface {
	// User category like operations
	IncrementUserCategoryLike(ctx context.Context, params videolearningdb.IncrementUserCategoryLikeParams) error
	DeleteAllUserCategoryLikes(ctx context.Context, userID int64) error
}
