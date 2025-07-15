package controllers

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/mappers"
)

// Helper function for timestamp conversion
func timestampToMillis(timestamp pgtype.Timestamptz) int64 {
	if timestamp.Valid {
		return timestamp.Time.UnixMilli()
	}
	return 0
}

func (s *textsrvc) CreateCourse(ctx context.Context, payload *text.CreateCoursePayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.courseRepo.CreateCourse(ctx, textdb.CreateCourseParams{
		Title: 	 payload.Title,
		Description: payload.Description,
		ImageUrl: pgtype.Text{
			String: *payload.ImageURL,
			Valid: payload.ImageURL != nil,
		},
	})
	if err != nil {
		return nil, text.InvalidInput("Failed to create course: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Course created successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) GetCourse(ctx context.Context, payload *text.GetCoursePayload) (res *text.Course, err error) {
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	course, err := s.courseRepo.GetCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.InvalidInput("Failed to get course: " + err.Error())
	}
	
	return mappers.CourseDBToAPI(&course), nil
}

func (s *textsrvc) ListCourses(ctx context.Context, payload *text.ListCoursesPayload) (res []*text.Course, err error) {
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	courses, err := s.courseRepo.ListCourses(ctx)
	if err != nil {
		return nil, text.InternalError("Failed to list courses: " + err.Error())
	}

	res = make([]*text.Course, len(courses))
	for i, course := range courses {
		res[i] = mappers.CourseDBToAPI(&course)
	}

	return res, nil
}

func (s *textsrvc) DeleteCourse(ctx context.Context, payload *text.DeleteCoursePayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.courseRepo.DeleteCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.InvalidInput("Failed to delete course: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Course deleted successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) UpdateCourse(ctx context.Context, payload *text.UpdateCoursePayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	updateParams := textdb.UpdateCourseParams{
		ID: payload.CourseID,
	}

	if payload.Title != nil {
		updateParams.Title = *payload.Title
	}

	if payload.Description != nil {
		updateParams.Description = *payload.Description
	}

	if payload.ImageURL != nil {
		updateParams.ImageUrl = pgtype.Text{
			String: *payload.ImageURL,
			Valid:  true,
		}
	}

	err = s.courseRepo.UpdateCourse(ctx, updateParams)
	if err != nil {
		return nil, text.InvalidInput("Failed to update course: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Course updated successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) CreateSection(ctx context.Context, payload *text.CreateSectionPayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	order, err := s.sectionRepo.GetNextOrderForCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.InternalError("Failed to get next order for course: " + err.Error())
	}

	err = s.sectionRepo.CreateSection(ctx, textdb.CreateSectionParams{
		CourseID:   payload.CourseID,
		Title:      payload.Title,
		Description: payload.Description,
		Order: order,
	})
	if err != nil {
		return nil, text.InvalidInput("Failed to create section: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Section created successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) GetSection(ctx context.Context, payload *text.GetSectionPayload) (res *text.Section, err error) {
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	section, err := s.sectionRepo.GetSection(ctx, payload.SectionID)
	if err != nil {
		return nil, text.NotFound("Section not found: " + err.Error())
	}

	return mappers.SectionDBToAPI(&section), nil
}

func (s *textsrvc) ListSections(ctx context.Context, payload *text.ListSectionsPayload) (res []*text.Section, err error) {
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	sections, err := s.sectionRepo.ListSectionsByCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.InternalError("Failed to list sections: " + err.Error())
	}

	res = make([]*text.Section, len(sections))

	for i, section := range sections {
		res[i] = mappers.SectionDBToAPI(&section)
	}
	
	return res, nil
}

func (s *textsrvc) UpdateSection(ctx context.Context, payload *text.UpdateSectionPayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.sectionRepo.UpdateSection(ctx, textdb.UpdateSectionParams{
		ID:          payload.SectionID,
		Title:       *payload.Title,
		Description: *payload.Description,
		Order:       int32(*payload.Order),
	})
	if err != nil {
		return nil, text.InvalidInput("Failed to update section: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Section updated successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) DeleteSection(ctx context.Context, payload *text.DeleteSectionPayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.sectionRepo.DeleteSection(ctx, payload.SectionID)
	if err != nil {
		return nil, text.InvalidInput("Failed to delete section: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Section deleted successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) CreateArticle(ctx context.Context, payload *text.CreateArticlePayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.articleRepo.CreateArticle(ctx, textdb.CreateArticleParams{
		SectionID: payload.SectionID,
		Title:     payload.Title,
		Content:   payload.Content,
	})
	if err != nil {
		return nil, text.InvalidInput("Failed to create article: " + err.Error())
	}
	
	return &text.SimpleResponse{
		Message: "Article created successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) GetArticle(ctx context.Context, payload *text.GetArticlePayload) (res *text.Article, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}
	
	article, err := s.articleRepo.GetArticle(ctx, payload.ArticleID)
	if err != nil {
		return nil, text.NotFound("Article not found: " + err.Error())
	}

	// Marcar automáticamente el artículo como completado cuando se accede a él
	// Solo para estudiantes (no profesores)
	if profileInfo.Role == "student" {
		_ = s.progressRepo.MarkArticleAsCompleted(ctx, textdb.MarkArticleAsCompletedParams{
			UserID:   profileInfo.UserID,
			ArticleID: payload.ArticleID,
		})
		// Ignoramos errores aquí para no afectar la obtención del artículo
	}

	return mappers.ArticleDBToAPI(&article), nil
}

func (s *textsrvc) ListArticles(ctx context.Context, payload *text.ListArticlesPayload) (res []*text.Article, err error) {
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	articles, err := s.articleRepo.ListArticlesBySection(ctx, payload.SectionID)
	if err != nil {
		return nil, text.InternalError("Failed to list articles: " + err.Error())
	}
	
	res = make([]*text.Article, len(articles))

	for i, article := range articles {
		res[i] = mappers.ArticleDBToAPI(&article)
	}

	return res, nil
}

func (s *textsrvc) UpdateArticle(ctx context.Context, payload *text.UpdateArticlePayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.articleRepo.UpdateArticle(ctx, textdb.UpdateArticleParams{
		ID:          payload.ArticleID,
		Title:       *payload.Title,
		Content:     *payload.Content,
	})
	if err != nil {
		return nil, text.InvalidInput("Failed to update article: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Article updated successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) DeleteArticle(ctx context.Context, payload *text.DeleteArticlePayload) (res *text.SimpleResponse, err error) {
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})	
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	if profileInfo.Role != "teacher" {
		return nil, text.PermissionDenied("Only teachers can create courses")
	}

	err = s.articleRepo.DeleteArticle(ctx, payload.ArticleID)
	if err != nil {
		return nil, text.InvalidInput("Failed to delete article: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Article deleted successfully",
		Success: true,
	}, nil
}

// --- Progress Methods Implementation ---

func (s *textsrvc) MarkArticleCompleted(ctx context.Context, payload *text.MarkArticleCompletedPayload) (res *text.SimpleResponse, err error) {
	// Obtener información del perfil del usuario
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Invalid session: " + err.Error())
	}

	// Verificar que el artículo existe
	_, err = s.articleRepo.GetArticle(ctx, payload.ArticleID)
	if err != nil {
		return nil, text.NotFound("Article not found: " + err.Error())
	}

	// Marcar como completado
	err = s.progressRepo.MarkArticleAsCompleted(ctx, textdb.MarkArticleAsCompletedParams{
		UserID:    profileInfo.UserID,
		ArticleID: payload.ArticleID,
	})
	if err != nil {
		return nil, text.InternalError("Failed to mark article as completed: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Article marked as completed successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) UnmarkArticleCompleted(ctx context.Context, payload *text.UnmarkArticleCompletedPayload) (res *text.SimpleResponse, err error) {
	// Obtener información del perfil del usuario
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Invalid session: " + err.Error())
	}

	// Verificar que el artículo existe
	_, err = s.articleRepo.GetArticle(ctx, payload.ArticleID)
	if err != nil {
		return nil, text.NotFound("Article not found: " + err.Error())
	}

	// Desmarcar como completado
	err = s.progressRepo.UnmarkArticleAsCompleted(ctx, textdb.UnmarkArticleAsCompletedParams{
		UserID:    profileInfo.UserID,
		ArticleID: payload.ArticleID,
	})
	if err != nil {
		return nil, text.InternalError("Failed to unmark article as completed: " + err.Error())
	}

	return &text.SimpleResponse{
		Message: "Article unmarked as completed successfully",
		Success: true,
	}, nil
}

func (s *textsrvc) CheckArticleCompleted(ctx context.Context, payload *text.CheckArticleCompletedPayload) (res *text.CheckArticleCompletedResult, err error) {
	// Obtener información del perfil del usuario
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Invalid session: " + err.Error())
	}

	// Verificar que el artículo existe
	_, err = s.articleRepo.GetArticle(ctx, payload.ArticleID)
	if err != nil {
		return nil, text.NotFound("Article not found: " + err.Error())
	}

	// Verificar si está completado
	completed, err := s.progressRepo.CheckArticleCompleted(ctx, textdb.CheckArticleCompletedParams{
		UserID:    profileInfo.UserID,
		ArticleID: payload.ArticleID,
	})
	if err != nil {
		return nil, text.InternalError("Failed to check article completion status: " + err.Error())
	}

	return &text.CheckArticleCompletedResult{
		Completed: completed,
	}, nil
}

// --- Course Content and Progress Methods Implementation ---

func (s *textsrvc) GetCourseContent(ctx context.Context, payload *text.GetCourseContentPayload) (res *text.CourseContent, err error) {
	// Validar sesión del usuario
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Invalid session: " + err.Error())
	}

	// Obtener información del curso
	course, err := s.courseRepo.GetCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.NotFound("Course not found: " + err.Error())
	}

	// Obtener todas las secciones del curso
	sections, err := s.sectionRepo.ListSectionsByCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.InternalError("Failed to get course sections: " + err.Error())
	}

	// Para cada sección, obtener sus artículos
	var sectionsWithArticles []*text.SectionWithArticles
	totalArticles := int64(0)

	for _, section := range sections {
		// Obtener artículos de la sección
		articles, err := s.articleRepo.ListArticlesBySection(ctx, section.ID)
		if err != nil {
			return nil, text.InternalError("Failed to get articles for section: " + err.Error())
		}

		// Convertir artículos a formato API
		var apiArticles []*text.Article
		for _, article := range articles {
			apiArticles = append(apiArticles, mappers.ArticleDBToAPI(&article))
		}

		// Crear sección con artículos
		sectionWithArticles := &text.SectionWithArticles{
			ID:          section.ID,
			CourseID:    section.CourseID,
			Title:       section.Title,
			Description: section.Description,
			Order:       int64(section.Order),
			CreatedAt:   timestampToMillis(section.CreatedAt),
			UpdatedAt:   timestampToMillis(section.UpdatedAt),
			Articles:    apiArticles,
		}

		sectionsWithArticles = append(sectionsWithArticles, sectionWithArticles)
		totalArticles += int64(len(articles))
	}

	// Crear respuesta completa
	courseContent := &text.CourseContent{
		Course:        mappers.CourseDBToAPI(&course),
		Sections:      sectionsWithArticles,
		TotalSections: int64(len(sections)),
		TotalArticles: totalArticles,
	}

	return courseContent, nil
}

func (s *textsrvc) GetUserCourseProgress(ctx context.Context, payload *text.GetUserCourseProgressPayload) (res *text.UserCourseProgress, err error) {
	// Validar sesión del usuario
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Invalid session: " + err.Error())
	}

	// Determinar el usuario del cual obtener el progreso
	targetUserID := profileInfo.UserID
	if payload.UserID != nil {
		// Solo permitir ver el progreso de otros usuarios si es profesor o admin
		if profileInfo.Role != "teacher" && profileInfo.Role != "admin" {
			return nil, text.PermissionDenied("Only teachers and admins can view other users' progress")
		}
		targetUserID = *payload.UserID
	}

	// Obtener información del curso
	course, err := s.courseRepo.GetCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.NotFound("Course not found: " + err.Error())
	}

	// Obtener todas las secciones del curso
	sections, err := s.sectionRepo.ListSectionsByCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.InternalError("Failed to get course sections: " + err.Error())
	}

	// Para cada sección, obtener sus artículos con progreso
	var sectionsWithProgress []*text.SectionWithProgress
	totalArticles := int64(0)
	totalCompletedArticles := int64(0)
	var lastAccessed int64 = 0

	for _, section := range sections {
		// Obtener artículos de la sección
		articles, err := s.articleRepo.ListArticlesBySection(ctx, section.ID)
		if err != nil {
			return nil, text.InternalError("Failed to get articles for section: " + err.Error())
		}

		// Para cada artículo, verificar si está completado
		var articlesWithProgress []*text.ArticleWithProgress
		sectionCompletedArticles := int64(0)

		for _, article := range articles {
			// Verificar si el artículo está completado
			completed, err := s.progressRepo.CheckArticleCompleted(ctx, textdb.CheckArticleCompletedParams{
				UserID:    targetUserID,
				ArticleID: article.ID,
			})
			if err != nil {
				// Si hay error, asumir que no está completado
				completed = false
			}

			var completedAt int64 = 0
			if completed {
				sectionCompletedArticles++
				totalCompletedArticles++
				// Aquí podrías obtener la fecha exacta de completación desde la base de datos
				// Por ahora usamos un timestamp placeholder
				completedAt = timestampToMillis(article.UpdatedAt)
				if completedAt > lastAccessed {
					lastAccessed = completedAt
				}
			}

			articleWithProgress := &text.ArticleWithProgress{
				ID:          article.ID,
				SectionID:   article.SectionID,
				Title:       article.Title,
				Content:     article.Content,
				CreatedAt:   timestampToMillis(article.CreatedAt),
				UpdatedAt:   timestampToMillis(article.UpdatedAt),
				Completed:   completed,
				CompletedAt: completedAt,
			}

			articlesWithProgress = append(articlesWithProgress, articleWithProgress)
		}

		// Calcular porcentaje de completación de la sección
		var sectionCompletionPercentage float64 = 0
		if len(articles) > 0 {
			sectionCompletionPercentage = float64(sectionCompletedArticles) / float64(len(articles)) * 100
		}

		// Crear sección con progreso
		sectionWithProgress := &text.SectionWithProgress{
			ID:                   section.ID,
			CourseID:            section.CourseID,
			Title:               section.Title,
			Description:         section.Description,
			Order:               int64(section.Order),
			CreatedAt:           timestampToMillis(section.CreatedAt),
			UpdatedAt:           timestampToMillis(section.UpdatedAt),
			Articles:            articlesWithProgress,
			TotalArticles:       int64(len(articles)),
			CompletedArticles:   sectionCompletedArticles,
			CompletionPercentage: sectionCompletionPercentage,
		}

		sectionsWithProgress = append(sectionsWithProgress, sectionWithProgress)
		totalArticles += int64(len(articles))
	}

	// Calcular porcentaje de completación general del curso
	var overallCompletionPercentage float64 = 0
	if totalArticles > 0 {
		overallCompletionPercentage = float64(totalCompletedArticles) / float64(totalArticles) * 100
	}

	// Crear respuesta completa
	userCourseProgress := &text.UserCourseProgress{
		Course:               mappers.CourseDBToAPI(&course),
		UserID:              targetUserID,
		Sections:            sectionsWithProgress,
		TotalSections:       int64(len(sections)),
		TotalArticles:       totalArticles,
		CompletedArticles:   totalCompletedArticles,
		CompletionPercentage: overallCompletionPercentage,
		LastAccessed:        lastAccessed,
	}

	return userCourseProgress, nil
}

func (s *textsrvc) GetCourseCompletionStats(ctx context.Context, payload *text.GetCourseCompletionStatsPayload) (res *text.CourseCompletionStats, err error) {
	// Validar sesión del usuario
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Invalid session: " + err.Error())
	}

	// Usar el ID del usuario autenticado
	targetUserID := profileInfo.UserID

	// Verificar que el curso existe
	_, err = s.courseRepo.GetCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.NotFound("Course not found: " + err.Error())
	}

	// Obtener estadísticas de completación del curso
	stats, err := s.progressRepo.GetCourseCompletionStats(ctx, textdb.GetCourseCompletionStatsParams{
		UserID:   targetUserID,
		CourseID: payload.CourseID,
	})
	if err != nil {
		return nil, text.InternalError("Failed to get course completion stats: " + err.Error())
	}

	// Crear respuesta
	courseStats := &text.CourseCompletionStats{
		CourseID:            payload.CourseID,
		TotalArticles:       stats.TotalArticles,
		CompletedArticles:   stats.CompletedArticles,
		CompletionPercentage: float64(stats.CompletionPercentage),
	}

	return courseStats, nil
}

func (s *textsrvc) GetCourseLeaderboard(ctx context.Context, payload *text.GetCourseLeaderboardPayload) (res *text.CourseLeaderboard, err error) {
	// Validate user session
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}

	// Verificar que el curso existe
	course, err := s.courseRepo.GetCourse(ctx, payload.CourseID)
	if err != nil {
		return nil, text.NotFound("Course not found: " + err.Error())
	}

	// Obtener el leaderboard del curso
	leaderboardData, err := s.progressRepo.GetCourseLeaderboard(ctx, textdb.GetCourseLeaderboardParams{
		CourseID: payload.CourseID,
		Limit:    int32(payload.Limit),
	})
	if err != nil {
		return nil, text.InternalError("Failed to get course leaderboard: " + err.Error())
	}

	// Obtener el total de artículos del curso para calcular porcentajes
	courseStats, err := s.progressRepo.GetCourseCompletionStats(ctx, textdb.GetCourseCompletionStatsParams{
		UserID:   0, // Usar 0 para obtener el total de artículos del curso
		CourseID: payload.CourseID,
	})
	if err != nil {
		return nil, text.InternalError("Failed to get course stats: " + err.Error())
	}

	// Convertir los datos a formato de respuesta
	var entries []*text.LeaderboardEntry
	for i, entry := range leaderboardData {
		// Obtener información del usuario desde el servicio de profiles
		userProfile, err := s.profilesServiceRepo.GetPublicProfileByID(ctx, &profiles.GetPublicProfileByIDPayload{
			UserID: entry.UserID,
		})
		
		var username string
		if err != nil {
			// Si no se puede obtener el perfil, usar un nombre por defecto
			username = "Unknown User"
		} else {
			username = userProfile.FirstName + " " + userProfile.LastName
		}

		// Calcular porcentaje de completación
		var completionPercentage float64 = 0
		if courseStats.TotalArticles > 0 {
			completionPercentage = float64(entry.CompletedCount) / float64(courseStats.TotalArticles) * 100
		}

		leaderboardEntry := &text.LeaderboardEntry{
			UserID:               entry.UserID,
			Username:            username,
			CompletionPercentage: completionPercentage,
			CompletedArticles:   entry.CompletedCount,
			TotalArticles:       courseStats.TotalArticles,
			Rank:                int64(i + 1), // Rank basado en posición (1-indexed)
			LastActivity:        nil, // Por ahora no tenemos esta información
		}
		entries = append(entries, leaderboardEntry)
	}

	// Obtener total de participantes en el curso
	totalParticipants, err := s.progressRepo.GetCourseParticipantCount(ctx, payload.CourseID)
	if err != nil {
		// Si falla, usar el número de entradas como fallback
		totalParticipants = int64(len(entries))
	}

	leaderboard := &text.CourseLeaderboard{
		CourseID:         course.ID,
		CourseTitle:      course.Title,
		Entries:          entries,
		TotalParticipants: totalParticipants,
		GeneratedAt:      mappers.TimestampToMillis(pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		}),
	}

	return leaderboard, nil
}