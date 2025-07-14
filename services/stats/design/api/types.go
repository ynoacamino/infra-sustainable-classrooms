package design

import (
	. "goa.design/goa/v3/dsl"
)

// SimpleResponse type for basic responses
var SimpleResponse = Type("SimpleResponse", func() {
	Description("Simple response with message and success status")
	Field(1, "message", String, "Response message", func() {
		Example("Operation completed successfully")
	})
	Field(2, "success", Boolean, "Whether the operation was successful", func() {
		Example(true)
	})
	Required("message", "success")
})

// ArticleProgress type for individual article progress
var ArticleProgress = Type("ArticleProgress", func() {
	Description("Progress information for a single article")
	Field(1, "article_id", Int64, "Article unique identifier", func() {
		Example(1001)
	})
	Field(2, "article_title", String, "Title of the article", func() {
		Example("Introduction to Go")
	})
	Field(3, "section_id", Int64, "Section unique identifier", func() {
		Example(101)
	})
	Field(4, "section_title", String, "Title of the section", func() {
		Example("Getting Started")
	})
	Field(5, "completed", Boolean, "Whether the article is completed", func() {
		Example(true)
	})
	Field(6, "completed_at", Int64, "Timestamp when completed (milliseconds since epoch, 0 if not completed)", func() {
		Example(1640995200000)
	})
	Required("article_id", "article_title", "section_id", "section_title", "completed", "completed_at")
})

// SectionProgress type for section-level progress
var SectionProgress = Type("SectionProgress", func() {
	Description("Progress information for a section")
	Field(1, "section_id", Int64, "Section unique identifier", func() {
		Example(101)
	})
	Field(2, "section_title", String, "Title of the section", func() {
		Example("Getting Started")
	})
	Field(3, "total_articles", Int64, "Total number of articles in the section", func() {
		Example(5)
	})
	Field(4, "completed_articles", Int64, "Number of completed articles", func() {
		Example(3)
	})
	Field(5, "completion_percentage", Float64, "Completion percentage for the section", func() {
		Example(60.0)
	})
	Field(6, "articles", ArrayOf(ArticleProgress), "List of articles with their progress")
	Required("section_id", "section_title", "total_articles", "completed_articles", "completion_percentage", "articles")
})

// CourseProgress type for complete course progress
var CourseProgress = Type("CourseProgress", func() {
	Description("Complete progress information for a course")
	Field(1, "course_id", Int64, "Course unique identifier", func() {
		Example(1)
	})
	Field(2, "course_title", String, "Title of the course", func() {
		Example("Introduction to Programming")
	})
	Field(3, "total_articles", Int64, "Total number of articles in the course", func() {
		Example(25)
	})
	Field(4, "completed_articles", Int64, "Number of completed articles", func() {
		Example(15)
	})
	Field(5, "completion_percentage", Float64, "Overall completion percentage for the course", func() {
		Example(60.0)
	})
	Field(6, "sections", ArrayOf(SectionProgress), "List of sections with their progress")
	Field(7, "last_accessed", Int64, "Timestamp of last access to any article in the course", func() {
		Example(1640995200000)
	})
	Required("course_id", "course_title", "total_articles", "completed_articles", "completion_percentage", "sections", "last_accessed")
})

// UserOverallStats type for user's overall statistics
var UserOverallStats = Type("UserOverallStats", func() {
	Description("User's overall learning statistics")
	Field(1, "user_id", Int64, "User unique identifier", func() {
		Example(123)
	})
	Field(2, "total_courses", Int64, "Total number of courses the user has access to", func() {
		Example(10)
	})
	Field(3, "courses_in_progress", Int64, "Number of courses currently in progress", func() {
		Example(3)
	})
	Field(4, "completed_courses", Int64, "Number of fully completed courses", func() {
		Example(2)
	})
	Field(5, "total_articles_completed", Int64, "Total number of articles completed across all courses", func() {
		Example(45)
	})
	Field(6, "overall_completion_percentage", Float64, "Overall completion percentage across all courses", func() {
		Example(72.5)
	})
	Field(7, "courses", ArrayOf(CourseProgress), "List of courses with their progress")
	Required("user_id", "total_courses", "courses_in_progress", "completed_courses", "total_articles_completed", "overall_completion_percentage", "courses")
})

var LeaderboardEntry = Type("LeaderboardEntry", func() {
	Field(1, "user_id", Int64, "User unique identifier")
	Field(2, "username", String, "Username")
	Field(3, "completion_percentage", Float64, "Course completion percentage")
	Field(4, "completed_articles", Int64, "Number of completed articles")
	Field(5, "last_activity", Int64, "Last activity timestamp")
	Required("user_id", "username", "completion_percentage", "completed_articles", "last_activity")
})