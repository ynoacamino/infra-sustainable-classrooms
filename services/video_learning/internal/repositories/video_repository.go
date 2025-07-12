package repositories

import (
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

// VideoRepository implementa ports.VideoRepository usando embedding directo
type VideoRepository struct {
	*videolearningdb.Queries
}

// NewVideoRepository crea una nueva instancia del repositorio de videos
func NewVideoRepository(db videolearningdb.DBTX) ports.VideoRepository {
	return &VideoRepository{
		Queries: videolearningdb.New(db),
	}
}

// NewVideoRepositoryWithTx crea una nueva instancia del repositorio de videos con transacción
func NewVideoRepositoryWithTx(tx videolearningdb.DBTX) ports.VideoRepository {
	return &VideoRepository{
		Queries: videolearningdb.New(tx),
	}
}
