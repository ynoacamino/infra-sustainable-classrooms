package repositories

import (
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

// VideoCategoryRepository implementa ports.VideoCategoryRepository usando embedding directo
type VideoCategoryRepository struct {
	*videolearningdb.Queries
}

// NewVideoCategoryRepository crea una nueva instancia del repositorio de categorías de video
func NewVideoCategoryRepository(db videolearningdb.DBTX) ports.VideoCategoryRepository {
	return &VideoCategoryRepository{
		Queries: videolearningdb.New(db),
	}
}

// NewVideoCategoryRepositoryWithTx crea una nueva instancia del repositorio de categorías de video con transacción
func NewVideoCategoryRepositoryWithTx(tx videolearningdb.DBTX) ports.VideoCategoryRepository {
	return &VideoCategoryRepository{
		Queries: videolearningdb.New(tx),
	}
}
