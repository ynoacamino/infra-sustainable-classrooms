package repositories

import (
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

// VideoTagRepository implementa ports.VideoTagRepository usando embedding directo
type VideoTagRepository struct {
	*videolearningdb.Queries
}

// NewVideoTagRepository crea una nueva instancia del repositorio de tags de video
func NewVideoTagRepository(db videolearningdb.DBTX) ports.VideoTagRepository {
	return &VideoTagRepository{
		Queries: videolearningdb.New(db),
	}
}

// NewVideoTagRepositoryWithTx crea una nueva instancia del repositorio de tags de video con transacción
func NewVideoTagRepositoryWithTx(tx videolearningdb.DBTX) ports.VideoTagRepository {
	return &VideoTagRepository{
		Queries: videolearningdb.New(tx),
	}
}
