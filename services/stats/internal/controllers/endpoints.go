package controllers

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/gen/stats"
)

// GetUserCourseProgress implements getting detailed progress for a user in a specific course
func (s *statssrvc) GetUserCourseProgress(ctx context.Context, payload *stats.GetUserCourseProgressPayload) (res *stats.CourseProgress, err error) {
	// Validar sesión del usuario
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Obtener progreso del curso desde el servicio text
	progressData, err := s.textServiceRepo.GetUserProgressForCourse(ctx, payload.UserID, payload.CourseID)
	if err != nil {
		return nil, stats.InternalError("Failed to get course progress: " + err.Error())
	}

	// Obtener estadísticas de completación del curso
	completionStats, err := s.textServiceRepo.GetCourseCompletionStats(ctx, payload.UserID, payload.CourseID)
	if err != nil {
		return nil, stats.InternalError("Failed to get completion stats: " + err.Error())
	}

	// Convertir los datos del progreso de artículos
	var articlesProgress []*stats.ArticleProgress
	var sectionsMap = make(map[int64]*stats.SectionProgress)

	for _, article := range progressData.Articles {
		articleProgress := &stats.ArticleProgress{
			ArticleID:     article.ArticleID,
			ArticleTitle:  article.ArticleTitle,
			SectionID:     article.SectionID,
			SectionTitle:  article.SectionTitle,
			Completed:     article.Completed,
			CompletedAt:   0,
		}

		if article.CompletedAt != nil {
			articleProgress.CompletedAt = *article.CompletedAt
		}

		articlesProgress = append(articlesProgress, articleProgress)

		// Calcular progreso por sección
		if _, exists := sectionsMap[article.SectionID]; !exists {
			sectionsMap[article.SectionID] = &stats.SectionProgress{
				SectionID:         article.SectionID,
				SectionTitle:      article.SectionTitle,
				TotalArticles:     0,
				CompletedArticles: 0,
				CompletionPercentage: 0,
			}
		}

		sectionsMap[article.SectionID].TotalArticles++
		if article.Completed {
			sectionsMap[article.SectionID].CompletedArticles++
		}
	}

	// Calcular porcentajes de secciones
	var sectionsProgress []*stats.SectionProgress
	for _, section := range sectionsMap {
		if section.TotalArticles > 0 {
			section.CompletionPercentage = float64(section.CompletedArticles) / float64(section.TotalArticles) * 100
		}
		sectionsProgress = append(sectionsProgress, section)
	}

	// Crear respuesta final
	courseProgress := &stats.CourseProgress{
		CourseID:             payload.CourseID,
		CourseTitle:          "", // Se podría obtener del servicio text
		TotalArticles:        completionStats.TotalArticles,
		CompletedArticles:    completionStats.CompletedArticles,
		CompletionPercentage: completionStats.CompletionPercentage,
		Sections:            sectionsProgress,
		LastAccessed:        0, // Se podría calcular desde las fechas de completación
	}

	return courseProgress, nil
}

// GetUserOverallStats implements getting overall statistics for a user across all courses
func (s *statssrvc) GetUserOverallStats(ctx context.Context, payload *stats.GetUserOverallStatsPayload) (res *stats.UserOverallStats, err error) {
	// Validar sesión del usuario
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Por ahora retornamos estadísticas básicas
	// En una implementación real, aquí haríamos múltiples llamadas al servicio text
	// para obtener estadísticas de todos los cursos del usuario
	overallStats := &stats.UserOverallStats{
		UserID:                     payload.UserID,
		TotalCourses:              0,
		CoursesInProgress:         0,
		CompletedCourses:          0,
		TotalArticlesCompleted:    0,
		OverallCompletionPercentage: 0,
		Courses:                   []*stats.CourseProgress{},
	}

	return overallStats, nil
}

// GetUserCompletedArticles implements getting list of completed articles for a user
func (s *statssrvc) GetUserCompletedArticles(ctx context.Context, payload *stats.GetUserCompletedArticlesPayload) (res []*stats.ArticleProgress, err error) {
	// Validar sesión del usuario
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Obtener artículos completados desde el servicio text
	completedData, err := s.textServiceRepo.GetUserCompletedArticles(ctx, payload.UserID)
	if err != nil {
		return nil, stats.InternalError("Failed to get completed articles: " + err.Error())
	}

	// Convertir a formato de respuesta
	var result []*stats.ArticleProgress
	for _, article := range completedData.Articles {
		articleProgress := &stats.ArticleProgress{
			ArticleID:     article.ArticleID,
			ArticleTitle:  article.Title,
			SectionID:     article.SectionID,
			SectionTitle:  "", // Se podría obtener del servicio text si es necesario
			Completed:     true,
			CompletedAt:   article.CompletedAt,
		}
		result = append(result, articleProgress)
	}

	return result, nil
}

// GetCourseLeaderboard implements getting leaderboard for a specific course
func (s *statssrvc) GetCourseLeaderboard(ctx context.Context, payload *stats.GetCourseLeaderboardPayload) (res []*stats.LeaderboardEntry, err error) {
	// Validar sesión del usuario
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Por ahora retornamos un leaderboard vacío
	// En una implementación real, aquí obtendríamos datos de múltiples usuarios
	// y calcularíamos su ranking basado en completación de artículos
	var leaderboard []*stats.LeaderboardEntry

	return leaderboard, nil
}
