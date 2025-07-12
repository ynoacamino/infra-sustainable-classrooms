package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// VideoCommentRepository define las operaciones de persistencia para comentarios de videos
type VideoCommentRepository interface {
	// Comment CRUD operations
	CreateComment(ctx context.Context, params videolearningdb.CreateCommentParams) (videolearningdb.VideoComment, error)
	GetCommentByID(ctx context.Context, id int64) (videolearningdb.VideoComment, error)
	UpdateComment(ctx context.Context, params videolearningdb.UpdateCommentParams) (videolearningdb.VideoComment, error)
	DeleteComment(ctx context.Context, params videolearningdb.DeleteCommentParams) error

	// Comment queries
	GetCommentsForVideo(ctx context.Context, params videolearningdb.GetCommentsForVideoParams) ([]videolearningdb.VideoComment, error)

	// Comment reply CRUD operations
	CreateCommentReply(ctx context.Context, params videolearningdb.CreateCommentReplyParams) (videolearningdb.VideoCommentReply, error)
	GetCommentReplyByID(ctx context.Context, id int64) (videolearningdb.VideoCommentReply, error)
	UpdateCommentReply(ctx context.Context, params videolearningdb.UpdateCommentReplyParams) (videolearningdb.VideoCommentReply, error)
	DeleteCommentReply(ctx context.Context, params videolearningdb.DeleteCommentReplyParams) error

	// Comment reply queries
	GetRepliesForComment(ctx context.Context, commentID int64) ([]videolearningdb.VideoCommentReply, error)

	// User activity
	GetUserCommentsAndReplies(ctx context.Context, params videolearningdb.GetUserCommentsAndRepliesParams) ([]videolearningdb.GetUserCommentsAndRepliesRow, error)
}
