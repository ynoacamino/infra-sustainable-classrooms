package ports

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
)

type TextServiceRepository interface {
	GetArticle(ctx context.Context, payload *text.GetArticlePayload) (*text.Article, error)
	ListArticlesBySection(ctx context.Context, payload *text.ListArticlesPayload) ([]*text.Article, error)

	GetCourse(ctx context.Context, payload *text.GetCoursePayload) (*text.Course, error)
	ListCourses(ctx context.Context, payload *text.ListCoursesPayload) ([]*text.Course, error)

	GetSection(ctx context.Context, payload *text.GetSectionPayload) (*text.Section, error)
	ListSectionsByCourse(ctx context.Context, payload *text.ListSectionsPayload) ([]*text.Section, error)
	
	MarkArticleAsCompleted(ctx context.Context, payload *text.MarkArticleCompletedPayload) (*text.SimpleResponse, error)
	
	UnmarkArticleAsCompleted(ctx context.Context, payload *text.UnmarkArticleCompletedPayload) (*text.SimpleResponse, error)
	
	CheckArticleCompleted(ctx context.Context, payload *text.CheckArticleCompletedPayload) (*text.CheckArticleCompletedResult, error)
	
	GetUserCourseProgress(ctx context.Context, payload *text.GetUserCourseProgressPayload) (*text.UserCourseProgress, error)
	
	GetCourseCompletionStats(ctx context.Context, payload *text.GetCourseCompletionStatsPayload) (*text.CourseCompletionStats, error)

	GetCourseLeaderboard(ctx context.Context, payload *text.GetCourseLeaderboardPayload) (*text.CourseLeaderboard, error)

	GetCourseContent(ctx context.Context, payload *text.GetCourseContentPayload) (*text.CourseContent, error)
}