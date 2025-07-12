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