package ports

import (
	"context"

	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

type VideoCommentRepository interface {
	// Comment operations
	CreateComment(ctx context.Context, params videolearningdb.CreateCommentParams) (videolearningdb.VideoComment, error)
	GetCommentByID(ctx context.Context, id int64) (videolearningdb.VideoComment, error)
	GetCommentsForVideo(ctx context.Context, params videolearningdb.GetCommentsForVideoParams) ([]videolearningdb.VideoComment, error)
	UpdateComment(ctx context.Context, params videolearningdb.UpdateCommentParams) (videolearningdb.VideoComment, error)
	DeleteComment(ctx context.Context, params videolearningdb.DeleteCommentParams) error

	// Comment reply operations
	CreateCommentReply(ctx context.Context, params videolearningdb.CreateCommentReplyParams) (videolearningdb.VideoCommentReply, error)
	GetCommentReplyByID(ctx context.Context, id int64) (videolearningdb.VideoCommentReply, error)
	GetRepliesForComment(ctx context.Context, commentID int64) ([]videolearningdb.VideoCommentReply, error)
	UpdateCommentReply(ctx context.Context, params videolearningdb.UpdateCommentReplyParams) (videolearningdb.VideoCommentReply, error)
	DeleteCommentReply(ctx context.Context, params videolearningdb.DeleteCommentReplyParams) error

	// User activity
	GetUserCommentsAndReplies(ctx context.Context, params videolearningdb.GetUserCommentsAndRepliesParams) ([]videolearningdb.GetUserCommentsAndRepliesRow, error)
}
