package design

import (
	. "goa.design/goa/v3/dsl"
)

// --- Basic Types ---

var LeaderboardEntry = Type("LeaderboardEntry", func() {
	Description("A single entry in the leaderboard")
	Field(1, "user_id", Int64, "User unique identifier", func() {
		Example(12345)
	})
	Field(2, "username", String, "Username for display", func() {
		Example("john_doe")
	})
	Field(3, "completion_percentage", Float64, "Completion percentage for the course (0-100)", func() {
		Example(85.5)
	})
	Field(4, "completed_articles", Int64, "Number of completed articles", func() {
		Example(17)
	})
	Field(5, "total_articles", Int64, "Total number of articles in the course", func() {
		Example(20)
	})
	Field(6, "rank", Int64, "User's rank in the leaderboard", func() {
		Example(3)
	})
	Field(7, "last_activity", Int64, "Timestamp of last activity (milliseconds)", func() {
		Example(1672531200000)
	})
	Required("user_id", "username", "completion_percentage", "completed_articles", "total_articles", "rank")
})

var CourseLeaderboard = Type("CourseLeaderboard", func() {
	Description("Course completion leaderboard")
	Field(1, "course_id", Int64, "Course unique identifier", func() {
		Example(12345)
	})
	Field(2, "course_title", String, "Course title", func() {
		Example("Introduction to Go Programming")
	})
	Field(3, "entries", ArrayOf(LeaderboardEntry), "Leaderboard entries sorted by rank", func() {
		Example([]map[string]interface{}{
			{
				"user_id":              12345,
				"username":            "john_doe", 
				"completion_percentage": 85.5,
				"completed_articles":   17,
				"total_articles":       20,
				"rank":                 1,
				"last_activity":        1672531200000,
			},
		})
	})
	Field(4, "total_participants", Int64, "Total number of participants in the course", func() {
		Example(150)
	})
	Field(5, "generated_at", Int64, "Timestamp when leaderboard was generated (milliseconds)", func() {
		Example(1672531200000)
	})
	Required("course_id", "course_title", "entries", "total_participants", "generated_at")
})

var CourseStatsEntry = Type("CourseStatsEntry", func() {
	Description("Statistics for a single course")
	Field(1, "course_id", Int64, "Course unique identifier", func() {
		Example(12345)
	})
	Field(2, "course_title", String, "Course title", func() {
		Example("Introduction to Go Programming")
	})
	Field(3, "completion_percentage", Float64, "Completion percentage for this course (0-100)", func() {
		Example(75.0)
	})
	Field(4, "completed_articles", Int64, "Number of completed articles in this course", func() {
		Example(15)
	})
	Field(5, "total_articles", Int64, "Total number of articles in this course", func() {
		Example(20)
	})
	Field(6, "last_accessed", Int64, "Timestamp of last access to this course (milliseconds)", func() {
		Example(1672531200000)
	})
	Required("course_id", "course_title", "completion_percentage", "completed_articles", "total_articles")
})

var UserOverallStats = Type("UserOverallStats", func() {
	Description("Overall statistics for a user across all courses")
	Field(1, "user_id", Int64, "User unique identifier", func() {
		Example(12345)
	})
	Field(2, "username", String, "Username for display", func() {
		Example("john_doe")
	})
	Field(3, "total_courses_enrolled", Int64, "Total number of courses the user is enrolled in", func() {
		Example(5)
	})
	Field(4, "total_courses_completed", Int64, "Total number of courses completed (100%)", func() {
		Example(2)
	})
	Field(5, "total_articles_completed", Int64, "Total number of articles completed across all courses", func() {
		Example(85)
	})
	Field(6, "overall_completion_percentage", Float64, "Overall completion percentage across all courses (0-100)", func() {
		Example(68.5)
	})
	Field(7, "courses", ArrayOf(CourseStatsEntry), "Detailed statistics for each course", func() {
		Example([]map[string]interface{}{
			{
				"course_id":              12345,
				"course_title":          "Introduction to Go Programming",
				"completion_percentage": 75.0,
				"completed_articles":    15,
				"total_articles":        20,
				"last_accessed":         1672531200000,
			},
		})
	})
	Field(8, "last_activity", Int64, "Timestamp of last activity (milliseconds)", func() {
		Example(1672531200000)
	})
	Required("user_id", "username", "total_courses_enrolled", "total_courses_completed", "total_articles_completed", "overall_completion_percentage", "courses")
})

var SectionProgressEntry = Type("SectionProgressEntry", func() {
	Description("Progress statistics for a course section")
	Field(1, "section_id", Int64, "Section unique identifier", func() {
		Example(123)
	})
	Field(2, "section_title", String, "Section title", func() {
		Example("Getting Started")
	})
	Field(3, "section_order", Int64, "Order of the section in the course", func() {
		Example(1)
	})
	Field(4, "completion_percentage", Float64, "Completion percentage for this section (0-100)", func() {
		Example(80.0)
	})
	Field(5, "completed_articles", Int64, "Number of completed articles in this section", func() {
		Example(4)
	})
	Field(6, "total_articles", Int64, "Total number of articles in this section", func() {
		Example(5)
	})
	Field(7, "last_accessed", Int64, "Timestamp of last access to this section (milliseconds)", func() {
		Example(1672531200000)
	})
	Required("section_id", "section_title", "section_order", "completion_percentage", "completed_articles", "total_articles")
})

var UserCourseProgressStats = Type("UserCourseProgressStats", func() {
	Description("Detailed progress statistics for a user in a specific course")
	Field(1, "user_id", Int64, "User unique identifier", func() {
		Example(12345)
	})
	Field(2, "course_id", Int64, "Course unique identifier", func() {
		Example(12345)
	})
	Field(3, "course_title", String, "Course title", func() {
		Example("Introduction to Go Programming")
	})
	Field(4, "overall_completion_percentage", Float64, "Overall completion percentage for the course (0-100)", func() {
		Example(75.0)
	})
	Field(5, "completed_articles", Int64, "Total number of completed articles in the course", func() {
		Example(15)
	})
	Field(6, "total_articles", Int64, "Total number of articles in the course", func() {
		Example(20)
	})
	Field(7, "sections", ArrayOf(SectionProgressEntry), "Progress details for each section", func() {
		Example([]map[string]interface{}{
			{
				"section_id":            123,
				"section_title":         "Getting Started",
				"section_order":         1,
				"completion_percentage": 80.0,
				"completed_articles":    4,
				"total_articles":        5,
				"last_accessed":         1672531200000,
			},
		})
	})
	Field(8, "enrollment_date", Int64, "Timestamp when user enrolled in the course (milliseconds)", func() {
		Example(1672531200000)
	})
	Field(9, "last_accessed", Int64, "Timestamp of last access to the course (milliseconds)", func() {
		Example(1672531200000)
	})
	Field(10, "estimated_completion_time", Int64, "Estimated time to complete remaining content (minutes)", func() {
		Example(120)
	})
	Required("user_id", "course_id", "course_title", "overall_completion_percentage", "completed_articles", "total_articles", "sections")
})

var CompletedArticleEntry = Type("CompletedArticleEntry", func() {
	Description("Information about a completed article")
	Field(1, "article_id", Int64, "Article unique identifier", func() {
		Example(456)
	})
	Field(2, "article_title", String, "Article title", func() {
		Example("Variables and Data Types")
	})
	Field(3, "section_id", Int64, "Section unique identifier", func() {
		Example(123)
	})
	Field(4, "section_title", String, "Section title", func() {
		Example("Basic Concepts")
	})
	Field(5, "course_id", Int64, "Course unique identifier", func() {
		Example(12345)
	})
	Field(6, "course_title", String, "Course title", func() {
		Example("Introduction to Go Programming")
	})
	Field(7, "completed_at", Int64, "Timestamp when article was completed (milliseconds)", func() {
		Example(1672531200000)
	})
	Field(8, "reading_time_estimate", Int64, "Estimated reading time in minutes", func() {
		Example(10)
	})
	Required("article_id", "article_title", "section_id", "section_title", "course_id", "course_title", "completed_at")
})

var UserCompletedArticles = Type("UserCompletedArticles", func() {
	Description("List of articles completed by a user")
	Field(1, "user_id", Int64, "User unique identifier", func() {
		Example(12345)
	})
	Field(2, "total_completed", Int64, "Total number of completed articles", func() {
		Example(85)
	})
	Field(3, "articles", ArrayOf(CompletedArticleEntry), "List of completed articles", func() {
		Example([]map[string]interface{}{
			{
				"article_id":             456,
				"article_title":          "Variables and Data Types",
				"section_id":             123,
				"section_title":          "Basic Concepts",
				"course_id":              12345,
				"course_title":           "Introduction to Go Programming",
				"completed_at":           1672531200000,
				"reading_time_estimate":  10,
			},
		})
	})
	Field(4, "last_completed_at", Int64, "Timestamp of most recently completed article (milliseconds)", func() {
		Example(1672531200000)
	})
	Required("user_id", "total_completed", "articles")
})