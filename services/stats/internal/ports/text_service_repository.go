package ports

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

// TextServiceRepository interface for communicating with the text service via gRPC
type TextServiceRepository interface {
	// Obtener el progreso de un usuario en un curso específico
	GetUserProgressForCourse(ctx context.Context, userID int64, courseID int64) (*UserProgressForCourseResponse, error)
	
	// Obtener todos los artículos completados por un usuario
	GetUserCompletedArticles(ctx context.Context, userID int64) (*UserCompletedArticlesResponse, error)
	
	// Obtener estadísticas de completación de un curso para un usuario
	GetCourseCompletionStats(ctx context.Context, userID int64, courseID int64) (*CourseCompletionStatsResponse, error)
}

// ProfilesServiceRepository interface for communicating with the profiles service via gRPC
type ProfilesServiceRepository interface {
	GetCompleteProfile(ctx context.Context, payload *profiles.GetCompleteProfilePayload) (*profiles.CompleteProfileResponse, error)
}

// Response types for text service communication
type UserProgressForCourseResponse struct {
	Articles []ArticleProgressData `json:"articles"`
}

type ArticleProgressData struct {
	ArticleID     int64  `json:"article_id"`
	ArticleTitle  string `json:"article_title"`
	SectionID     int64  `json:"section_id"`
	SectionTitle  string `json:"section_title"`
	Completed     bool   `json:"completed"`
	CompletedAt   *int64 `json:"completed_at"`
}

type UserCompletedArticlesResponse struct {
	Articles []CompletedArticleData `json:"articles"`
}

type CompletedArticleData struct {
	ArticleID   int64  `json:"article_id"`
	Title       string `json:"title"`
	SectionID   int64  `json:"section_id"`
	CompletedAt int64  `json:"completed_at"`
}

type CourseCompletionStatsResponse struct {
	TotalArticles      int64   `json:"total_articles"`
	CompletedArticles  int64   `json:"completed_articles"`
	CompletionPercentage float64 `json:"completion_percentage"`
}
