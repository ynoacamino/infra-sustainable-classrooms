package design

import (
	. "goa.design/goa/v3/dsl"
)

var Course = Type("Course", func() {
	Description("Course details")

	Field(1, "id", Int64, "Unique identifier for the course", func() {
		Example(12345)
	})
	Field(2, "title", String, "Title of the course", func() {
		Example("Introduction to Go")
	})
	Field(3, "description", String, "Description of the course", func() {
		Example("Learn the basics of Go programming language.")
	})
	Field(4, "imageUrl", String, "URL of the course image", func() {
		Example("https://example.com/course-image.jpg")
	})

	Field(5, "created_at", Int64, "Timestamp when the course was created (milliseconds since epoch)")
	Field(6, "updated_at", Int64, "Timestamp when the course was last updated (milliseconds since epoch)")
	Required("id", "title", "description", "created_at", "updated_at")
})

var SimpleResponse = Type("SimpleResponse", func() {
	Description("Basic response with success status and message")

	Field(1, "success", Boolean, "Operation success status")
	Field(2, "message", String, "Response message")

	Required("success", "message")
})
// Section type for course sections
var Section = Type("Section", func() {
	Description("Section details within a course")
	Field(1, "id", Int64, "Unique identifier for the section", func() {
		Example(101)
	})
	Field(2, "course_id", Int64, "ID of the parent course", func() {
		Example(12345)
	})
	Field(3, "title", String, "Title of the section", func() {
		Example("Getting Started")
	})
	Field(4, "description", String, "Description of the section", func() {
		Example("Introduction to the course structure.")
	})
	Field(5, "order", Int64, "Order of the section in the course (autonumbered for frontend rendering)", func() {
		Example(1)
		Minimum(1)
	})
	Field(6, "created_at", Int64, "Timestamp when the section was created (milliseconds since epoch)")
	Field(7, "updated_at", Int64, "Timestamp when the section was last updated (milliseconds since epoch)")
	Required("id", "course_id", "description", "title", "order", "created_at", "updated_at")
})

// Article type for section articles
var Article = Type("Article", func() {
	Description("Article details within a section")
	Field(1, "id", Int64, "Unique identifier for the article", func() {
		Example(1001)
	})
	Field(2, "section_id", Int64, "ID of the parent section", func() {
		Example(101)
	})
	Field(3, "title", String, "Title of the article", func() {
		Example("What is Go?")
	})
	Field(4, "content", String, "Content of the article", func() {
		Example("Go is an open source programming language...")
	})
	Field(5, "created_at", Int64, "Timestamp when the article was created (milliseconds since epoch)")
	Field(6, "updated_at", Int64, "Timestamp when the article was last updated (milliseconds since epoch)")
	Required("id", "section_id", "title", "content", "created_at", "updated_at")
})

// SectionWithArticles type for section containing its articles
var SectionWithArticles = Type("SectionWithArticles", func() {
	Description("Section with all its articles")
	Field(1, "id", Int64, "Unique identifier for the section", func() {
		Example(101)
	})
	Field(2, "course_id", Int64, "ID of the parent course", func() {
		Example(1)
	})
	Field(3, "title", String, "Title of the section", func() {
		Example("Getting Started")
	})
	Field(4, "description", String, "Description of the section", func() {
		Example("Introduction to the course fundamentals.")
	})
	Field(5, "order", Int64, "Order of the section within the course")
	Field(6, "created_at", Int64, "Timestamp when the section was created (milliseconds since epoch)")
	Field(7, "updated_at", Int64, "Timestamp when the section was last updated (milliseconds since epoch)")
	Field(8, "articles", ArrayOf(Article), "List of articles in this section")
	Required("id", "course_id", "title", "description", "order", "created_at", "updated_at", "articles")
})

// CourseContent type for complete course structure
var CourseContent = Type("CourseContent", func() {
	Description("Complete course content with all sections and articles")
	Field(1, "course", Course, "Course information")
	Field(2, "sections", ArrayOf(SectionWithArticles), "List of sections with their articles")
	Field(3, "total_sections", Int64, "Total number of sections in the course")
	Field(4, "total_articles", Int64, "Total number of articles in the course")
	Required("course", "sections", "total_sections", "total_articles")
})

// ArticleWithProgress type for article with completion status
var ArticleWithProgress = Type("ArticleWithProgress", func() {
	Description("Article with completion progress information")
	Field(1, "id", Int64, "Unique identifier for the article")
	Field(2, "section_id", Int64, "ID of the parent section")
	Field(3, "title", String, "Title of the article")
	Field(4, "content", String, "Content of the article")
	Field(5, "created_at", Int64, "Timestamp when the article was created (milliseconds since epoch)")
	Field(6, "updated_at", Int64, "Timestamp when the article was last updated (milliseconds since epoch)")
	Field(7, "completed", Boolean, "Whether the article is completed by the user")
	Field(8, "completed_at", Int64, "Timestamp when completed (0 if not completed)")
	Required("id", "section_id", "title", "content", "created_at", "updated_at", "completed", "completed_at")
})

// SectionWithProgress type for section with completion progress
var SectionWithProgress = Type("SectionWithProgress", func() {
	Description("Section with progress information for a specific user")
	Field(1, "id", Int64, "Unique identifier for the section")
	Field(2, "course_id", Int64, "ID of the parent course")
	Field(3, "title", String, "Title of the section")
	Field(4, "description", String, "Description of the section")
	Field(5, "order", Int64, "Order of the section within the course")
	Field(6, "created_at", Int64, "Timestamp when the section was created (milliseconds since epoch)")
	Field(7, "updated_at", Int64, "Timestamp when the section was last updated (milliseconds since epoch)")
	Field(8, "articles", ArrayOf(ArticleWithProgress), "List of articles with their completion status")
	Field(9, "total_articles", Int64, "Total number of articles in this section")
	Field(10, "completed_articles", Int64, "Number of completed articles in this section")
	Field(11, "completion_percentage", Float64, "Completion percentage for this section (0-100)")
	Required("id", "course_id", "title", "description", "order", "created_at", "updated_at", "articles", "total_articles", "completed_articles", "completion_percentage")
})

// UserCourseProgress type for user's progress in a course
var UserCourseProgress = Type("UserCourseProgress", func() {
	Description("User's progress in a specific course with detailed completion information")
	Field(1, "course", Course, "Course information")
	Field(2, "user_id", Int64, "User unique identifier")
	Field(3, "sections", ArrayOf(SectionWithProgress), "List of sections with progress information")
	Field(4, "total_sections", Int64, "Total number of sections in the course")
	Field(5, "total_articles", Int64, "Total number of articles in the course")
	Field(6, "completed_articles", Int64, "Number of completed articles")
	Field(7, "completion_percentage", Float64, "Overall completion percentage for the course (0-100)")
	Field(8, "last_accessed", Int64, "Timestamp of last access to any article in the course")
	Required("course", "user_id", "sections", "total_sections", "total_articles", "completed_articles", "completion_percentage", "last_accessed")
})

var CourseCompletionStats = Type("CourseCompletionStats", func() {
	Description("Statistics for course completion")
	Field(1, "course_id", Int64, "Unique identifier for the course", func() {
		Example(12345)
	})
	Field(2, "total_articles", Int64, "Total number of articles in the course", func() {
		Example(25)
	})
	Field(3, "completed_articles", Int64, "Number of completed articles in the course", func() {
		Example(15)
	})
	Field(4, "completion_percentage", Float64, "Overall completion percentage for the course (0-100)", func() {
		Example(60.0)
	})
	Required("course_id", "total_articles", "completed_articles", "completion_percentage")
})

var CheckArticleCompletedResult = Type("CheckArticleCompletedResult", func() {
	Field(1, "completed", Boolean, "Whether the article is completed")
	Required("completed")
})

var CourseLeaderboard = Type("CourseLeaderboard", func() {
	Description("Course completion leaderboard")
	Field(1, "course_id", Int64, "Unique identifier for the course", func() {
		Example(12345)
	})
	Field(2, "course_title", String, "Title of the course", func() {
		Example("Introduction to Go")
	})
	Field(3, "entries", ArrayOf(LeaderboardEntry), "Leaderboard entries sorted by completion")
	Field(4, "total_participants", Int64, "Total number of participants in the course", func() {
		Example(25)
	})
	Field(5, "generated_at", Int64, "Timestamp when leaderboard was generated")
	Required("course_id", "course_title", "entries", "total_participants", "generated_at")
})

var LeaderboardEntry = Type("LeaderboardEntry", func() {
	Description("Single entry in course leaderboard")
	Field(1, "user_id", Int64, "Unique identifier for the user", func() {
		Example(67890)
	})
	Field(2, "username", String, "Display name of the user", func() {
		Example("John Doe")
	})
	Field(3, "completion_percentage", Float64, "Course completion percentage (0-100)", func() {
		Example(85.5)
	})
	Field(4, "completed_articles", Int64, "Number of completed articles", func() {
		Example(17)
	})
	Field(5, "total_articles", Int64, "Total number of articles in course", func() {
		Example(20)
	})
	Field(6, "rank", Int64, "Position in leaderboard (1-based)", func() {
		Example(3)
	})
	Field(7, "last_activity", Int64, "Timestamp of last activity in the course")
	Required("user_id", "username", "completion_percentage", "completed_articles", "total_articles", "rank")
})