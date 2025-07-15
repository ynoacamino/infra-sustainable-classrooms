package repositories

import (
	"context"

	genTextclient "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/grpc/text/client"
	gentext "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/ports"
	"google.golang.org/grpc"
)

type TextServiceRepository struct {
	client *gentext.Client
}

func NewTextServiceRepository(grpcConn *grpc.ClientConn) ports.TextServiceRepository {
	grpcClient := genTextclient.NewClient(grpcConn)

	client := &gentext.Client{
		// Course endpoints
		GetCourseEndpoint:   grpcClient.GetCourse(),
		ListCoursesEndpoint: grpcClient.ListCourses(),

		// Section endpoints
		GetSectionEndpoint:   grpcClient.GetSection(),
		ListSectionsEndpoint: grpcClient.ListSections(),

		// Article endpoints
		GetArticleEndpoint:   grpcClient.GetArticle(),
		ListArticlesEndpoint: grpcClient.ListArticles(),

		// Progress endpoints
		MarkArticleCompletedEndpoint:   grpcClient.MarkArticleCompleted(),
		UnmarkArticleCompletedEndpoint: grpcClient.UnmarkArticleCompleted(),
		CheckArticleCompletedEndpoint:  grpcClient.CheckArticleCompleted(),

		// Course content and progress endpoints
		GetCourseContentEndpoint:         grpcClient.GetCourseContent(),
		GetUserCourseProgressEndpoint:    grpcClient.GetUserCourseProgress(),
		GetCourseCompletionStatsEndpoint: grpcClient.GetCourseCompletionStats(),
		GetCourseLeaderboardEndpoint:     grpcClient.GetCourseLeaderboard(),
	}

	return &TextServiceRepository{
		client: client,
	}
}

// Implement the interface methods by delegating to the client

func (r *TextServiceRepository) GetArticle(ctx context.Context, payload *gentext.GetArticlePayload) (*gentext.Article, error) {
	return r.client.GetArticle(ctx, payload)
}

func (r *TextServiceRepository) ListArticlesBySection(ctx context.Context, payload *gentext.ListArticlesPayload) ([]*gentext.Article, error) {
	return r.client.ListArticles(ctx, payload)
}

func (r *TextServiceRepository) GetCourse(ctx context.Context, payload *gentext.GetCoursePayload) (*gentext.Course, error) {
	return r.client.GetCourse(ctx, payload)
}

func (r *TextServiceRepository) ListCourses(ctx context.Context, payload *gentext.ListCoursesPayload) ([]*gentext.Course, error) {
	return r.client.ListCourses(ctx, payload)
}

func (r *TextServiceRepository) GetSection(ctx context.Context, payload *gentext.GetSectionPayload) (*gentext.Section, error) {
	return r.client.GetSection(ctx, payload)
}

func (r *TextServiceRepository) ListSectionsByCourse(ctx context.Context, payload *gentext.ListSectionsPayload) ([]*gentext.Section, error) {
	return r.client.ListSections(ctx, payload)
}

func (r *TextServiceRepository) MarkArticleAsCompleted(ctx context.Context, payload *gentext.MarkArticleCompletedPayload) (*gentext.SimpleResponse, error) {
	return r.client.MarkArticleCompleted(ctx, payload)
}

func (r *TextServiceRepository) UnmarkArticleAsCompleted(ctx context.Context, payload *gentext.UnmarkArticleCompletedPayload) (*gentext.SimpleResponse, error) {
	return r.client.UnmarkArticleCompleted(ctx, payload)
}

func (r *TextServiceRepository) CheckArticleCompleted(ctx context.Context, payload *gentext.CheckArticleCompletedPayload) (*gentext.CheckArticleCompletedResult, error) {
	return r.client.CheckArticleCompleted(ctx, payload)
}

func (r *TextServiceRepository) GetUserCourseProgress(ctx context.Context, payload *gentext.GetUserCourseProgressPayload) (*gentext.UserCourseProgress, error) {
	return r.client.GetUserCourseProgress(ctx, payload)
}

func (r *TextServiceRepository) GetCourseCompletionStats(ctx context.Context, payload *gentext.GetCourseCompletionStatsPayload) (*gentext.CourseCompletionStats, error) {
	return r.client.GetCourseCompletionStats(ctx, payload)
}

func (r *TextServiceRepository) GetCourseContent(ctx context.Context, payload *gentext.GetCourseContentPayload) (*gentext.CourseContent, error) {
	return r.client.GetCourseContent(ctx, payload)
}

func (r *TextServiceRepository) GetCourseLeaderboard(ctx context.Context, payload *gentext.GetCourseLeaderboardPayload) (*gentext.CourseLeaderboard, error) {
	return r.client.GetCourseLeaderboard(ctx, payload)
}