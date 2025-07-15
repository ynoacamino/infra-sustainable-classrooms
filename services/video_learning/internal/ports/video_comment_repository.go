package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// VideoCommentRepository define las operaciones de persistencia para comentarios de videos
type VideoCommentRepository interface {
	CreateComment(ctx context.Context, params videolearningdb.CreateCommentParams) (videolearningdb.VideoComment, error)
	GetCommentsForVideo(ctx context.Context, params videolearningdb.GetCommentsForVideoParams) ([]videolearningdb.VideoComment, error)
	DeleteComment(ctx context.Context, params videolearningdb.DeleteCommentParams) error
}
