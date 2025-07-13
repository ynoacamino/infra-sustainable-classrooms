package mappers

import (
	"github.com/jackc/pgx/v5/pgtype"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
)

func timestampToMillis(timestamp pgtype.Timestamptz) int64 {
	if timestamp.Valid {
		return timestamp.Time.UnixMilli()
	}
	return 0
}

func PGTextToString(pgText pgtype.Text) *string {
	if pgText.Valid {
		return &pgText.String
	}
	return nil
}

func CourseDBToAPI(course *textdb.Course) *text.Course {
	if course == nil {
		return nil
	}

	apiCourse := &text.Course{
		ID:         course.ID,
		Title: 		course.Title,
		Description: course.Description,
		ImageURL: PGTextToString(course.ImageUrl),
		CreatedAt: timestampToMillis(course.CreatedAt),
		UpdatedAt: timestampToMillis(course.UpdatedAt),
	}

	return apiCourse
}

func SectionDBToAPI(section *textdb.Section) *text.Section {
	if section == nil {
		return nil
	}

	apiSection := &text.Section{
		ID: section.ID,
		CourseID: section.CourseID,
		Title: section.Title,
		Description: section.Description,
		Order: int64(section.Order),
		CreatedAt: timestampToMillis(section.CreatedAt),
		UpdatedAt: timestampToMillis(section.UpdatedAt),
	}
	
	return apiSection
}

func ArticleDBToAPI(article *textdb.Article) *text.Article {
	if article == nil {
		return nil
	}

	apiArticle := &text.Article{
		ID:        article.ID,
		SectionID: article.SectionID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: timestampToMillis(article.CreatedAt),
		UpdatedAt: timestampToMillis(article.UpdatedAt),
	}

	return apiArticle
}