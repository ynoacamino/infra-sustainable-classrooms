package repositories

import (
	"context"

	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
)

type ProgressRepository struct {
	*textdb.Queries
}

func NewProgressRepository(db textdb.DBTX) ports.ProgressRepository {
	return &ProgressRepository{
		Queries: textdb.New(db),
	}
}

func (r *ProgressRepository) MarkArticleAsCompleted(ctx context.Context, userID int64, articleID int64) error {
	return r.Queries.MarkArticleAsCompleted(ctx, textdb.MarkArticleAsCompletedParams{
		UserID:    userID,
		ArticleID: articleID,
	})
}

func (r *ProgressRepository) UnmarkArticleAsCompleted(ctx context.Context, userID int64, articleID int64) error {
	return r.Queries.UnmarkArticleAsCompleted(ctx, textdb.UnmarkArticleAsCompletedParams{
		UserID:    userID,
		ArticleID: articleID,
	})
}

func (r *ProgressRepository) IsArticleCompleted(ctx context.Context, userID int64, articleID int64) (bool, error) {
	result, err := r.Queries.IsArticleCompleted(ctx, textdb.IsArticleCompletedParams{
		UserID:    userID,
		ArticleID: articleID,
	})
	if err != nil {
		return false, err
	}
	return result, nil
}

func (r *ProgressRepository) GetUserCompletedArticles(ctx context.Context, userID int64) ([]textdb.GetUserCompletedArticlesRow, error) {
	return r.Queries.GetUserCompletedArticles(ctx, userID)
}

func (r *ProgressRepository) GetUserProgressForCourse(ctx context.Context, userID int64, courseID int64) ([]textdb.GetUserProgressForCourseRow, error) {
	return r.Queries.GetUserProgressForCourse(ctx, textdb.GetUserProgressForCourseParams{
		UserID:   userID,
		CourseID: courseID,
	})
}

func (r *ProgressRepository) GetCourseCompletionStats(ctx context.Context, userID int64, courseID int64) (textdb.GetCourseCompletionStatsRow, error) {
	return r.Queries.GetCourseCompletionStats(ctx, textdb.GetCourseCompletionStatsParams{
		UserID:   userID,
		CourseID: courseID,
	})
}
