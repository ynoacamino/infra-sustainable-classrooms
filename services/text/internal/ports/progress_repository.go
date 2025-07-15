package ports

import (
	"context"

	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type ProgressRepository interface {
	MarkArticleAsCompleted(ctx context.Context, arg textdb.MarkArticleAsCompletedParams) error
	
	UnmarkArticleAsCompleted(ctx context.Context, arg textdb.UnmarkArticleAsCompletedParams) error
	
	CheckArticleCompleted(ctx context.Context, arg textdb.CheckArticleCompletedParams) (bool, error)
	
	GetUserCompletedArticles(ctx context.Context, userID int64) ([]textdb.GetUserCompletedArticlesRow, error)
	
	GetUserProgressForCourse(ctx context.Context, arg textdb.GetUserProgressForCourseParams) ([]textdb.GetUserProgressForCourseRow, error)
	
	GetCourseCompletionStats(ctx context.Context, arg textdb.GetCourseCompletionStatsParams) (textdb.GetCourseCompletionStatsRow, error)
	
	GetCourseLeaderboard(ctx context.Context, arg textdb.GetCourseLeaderboardParams) ([]textdb.GetCourseLeaderboardRow, error)
	
	GetCourseParticipantCount(ctx context.Context, courseID int64) (int64, error)
}
