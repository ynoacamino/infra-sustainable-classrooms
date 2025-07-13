package controllers

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/mappers"
)

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
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, text.Unauthorized("Unauthorized: " + err.Error())
	}
	
	article, err := s.articleRepo.GetArticle(ctx, payload.ArticleID)
	if err != nil {
		return nil, text.NotFound("Article not found: " + err.Error())
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