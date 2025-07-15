package repositories

import (
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

// VideoCommentRepository implementa ports.VideoCommentRepository usando embedding directo
type VideoCommentRepository struct {
	*videolearningdb.Queries
}

// NewVideoCommentRepository crea una nueva instancia del repositorio de comentarios de video
func NewVideoCommentRepository(db videolearningdb.DBTX) ports.VideoCommentRepository {
	return &VideoCommentRepository{
		Queries: videolearningdb.New(db),
	}
}

// NewVideoCommentRepositoryWithTx crea una nueva instancia del repositorio de comentarios de video con transacción
func NewVideoCommentRepositoryWithTx(tx videolearningdb.DBTX) ports.VideoCommentRepository {
	return &VideoCommentRepository{
		Queries: videolearningdb.New(tx),
	}
}
