package ports

import (
	"context"

	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type ProgressRepository interface {
	// Marcar un artículo como completado por un usuario
	MarkArticleAsCompleted(ctx context.Context, userID int64, articleID int64) error
	
	// Desmarcar un artículo como completado por un usuario
	UnmarkArticleAsCompleted(ctx context.Context, userID int64, articleID int64) error
	
	// Verificar si un artículo está completado por un usuario
	IsArticleCompleted(ctx context.Context, userID int64, articleID int64) (bool, error)
	
	// Obtener todos los artículos completados por un usuario
	GetUserCompletedArticles(ctx context.Context, userID int64) ([]textdb.GetUserCompletedArticlesRow, error)
	
	// Obtener el progreso de un usuario en un curso específico
	GetUserProgressForCourse(ctx context.Context, userID int64, courseID int64) ([]textdb.GetUserProgressForCourseRow, error)
	
	// Obtener estadísticas de completación de un curso para un usuario
	GetCourseCompletionStats(ctx context.Context, userID int64, courseID int64) (textdb.GetCourseCompletionStatsRow, error)
}
