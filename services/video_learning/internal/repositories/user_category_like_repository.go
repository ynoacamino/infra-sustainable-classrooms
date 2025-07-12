package repositories

import (
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

// UserCategoryLikeRepository implementa ports.UserCategoryLikeRepository usando embedding directo
type UserCategoryLikeRepository struct {
	*videolearningdb.Queries
}

// NewUserCategoryLikeRepository crea una nueva instancia del repositorio de likes de categorías por usuario
func NewUserCategoryLikeRepository(db videolearningdb.DBTX) ports.UserCategoryLikeRepository {
	return &UserCategoryLikeRepository{
		Queries: videolearningdb.New(db),
	}
}

// NewUserCategoryLikeRepositoryWithTx crea una nueva instancia del repositorio de likes de categorías por usuario con transacción
func NewUserCategoryLikeRepositoryWithTx(tx videolearningdb.DBTX) ports.UserCategoryLikeRepository {
	return &UserCategoryLikeRepository{
		Queries: videolearningdb.New(tx),
	}
}
