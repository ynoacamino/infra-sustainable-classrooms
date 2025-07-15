package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("text", func() {
	Description("Course management service for text-based learning, with sections and articles. Only teachers can create, update, or delete content. All endpoints require session_token via cookie.")

	HTTP(func() {
		Path("/text")
	})

	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	Error("invalid_input", String, "Invalid input parameters")
	Error("rate_limited", String, "Too many requests")
	Error("service_unavailable", String, "Service temporarily unavailable")
	Error("unauthorized", String, "Unauthorized access")
	Error("permission_denied", String, "Permission denied (only teachers)")
	Error("not_found", String, "Resource not found")
	Error("internal_error", String, "Internal server error")

	// --- Courses ---
	Method("CreateCourse", func() {
		Description("Create a new text-based course (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "title", String, "Course title", func() {
				Example("Introduction to Go")
				MinLength(3)
				MaxLength(150)
			})
			Field(3, "description", String, "Course description", func() {
				Example("Learn the basics of Go programming language.")
				MinLength(10)
				MaxLength(300)
			})
			Field(4, "imageUrl", String, "Course image URL", func() {
				Example("https://example.com/course-image.jpg")
				Format("uri")
				MinLength(5)
				MaxLength(500)
			})
			Required("session_token", "title", "description")
		})
		Result(SimpleResponse)
		HTTP(func() {
			POST("/courses")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("GetCourse", func() {
		Description("Retrieve course details by ID")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier", func() {
				Example(12345)
			})
			Required("session_token", "course_id")
		})
		Result(Course)
		HTTP(func() {
			GET("/courses/{course_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("ListCourses", func() {
		Description("List all available courses")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Required("session_token")
		})
		Result(ArrayOf(Course))
		HTTP(func() {
			GET("/courses")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("DeleteCourse", func() {
		Description("Delete a course by ID (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier", func() {
				Example(12345)
			})
			Required("session_token", "course_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			DELETE("/courses/{course_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("UpdateCourse", func() {
		Description("Update course details by ID (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier", func() {
				Example(12345)
			})
			Field(3, "title", String, "Course title", func() {
				Example("Advanced Go Programming")
				MinLength(3)
				MaxLength(150)
			})
			Field(4, "description", String, "Course description", func() {
				Example("Deep dive into Go programming language features.")
				MinLength(10)
				MaxLength(300)
			})
			Field(5, "imageUrl", String, "Course image URL", func() {
				Example("https://example.com/updated-course-image.jpg")
				Format("uri")
				MinLength(5)
				MaxLength(500)
			})
			Required("session_token", "course_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			PATCH("/courses/{course_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// --- Sections ---
	Method("CreateSection", func() {
		Description("Create a new section in a course (teachers only)")
	   Payload(func() {
		   Field(1, "session_token", String, "Authentication session token")
		   Field(2, "course_id", Int64, "Course unique identifier")
		   Field(3, "title", String, "Section title", func() {
			   Example("Getting Started")
			   MinLength(3)
			   MaxLength(100)
		   })
		   Field(4, "description", String, "Section description", func() {
			   Example("Introduction to the course structure.")
			   MinLength(5)
			   MaxLength(200)
		   })
		   Field(5, "order", Int64, "Order of the section in the course (optional, if not set it will be auto-numbered)", func() {
			   Example(1)
			   Minimum(1)
		   })
		   Required("session_token", "description", "course_id", "title")
	   })
		Result(SimpleResponse)
		HTTP(func() {
			POST("/courses/{course_id}/sections")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("GetSection", func() {
		Description("Get section details by ID")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "section_id", Int64, "Section unique identifier")
			Required("session_token", "section_id")
		})
		Result(Section)
		HTTP(func() {
			GET("/sections/{section_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("ListSections", func() {
		Description("List all sections for a course")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier")
			Required("session_token", "course_id")
		})
		Result(ArrayOf(Section))
		HTTP(func() {
			GET("/courses/{course_id}/sections")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("UpdateSection", func() {
		Description("Update section details (teachers only)")
	   Payload(func() {
		   Field(1, "session_token", String, "Authentication session token")
		   Field(2, "section_id", Int64, "Section unique identifier")
		   Field(3, "title", String, "Section title", func() {
			   Example("Updated Section Title.")
			   MinLength(3)
			   MaxLength(100)
		   })
		   Field(4, "description", String, "Section description", func() {
			   Example("Updated section description.")
			   MinLength(5)
			   MaxLength(200)
		   })
		   Field(5, "order", Int64, "Order of the section in the course (optional, if set will update the order)", func() {
			   Example(2)
			   Minimum(1)
		   })
		   Required("session_token", "section_id")
	   })
		Result(SimpleResponse)
		HTTP(func() {
			PATCH("/sections/{section_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("DeleteSection", func() {
		Description("Delete a section (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "section_id", Int64, "Section unique identifier")
			Required("session_token", "section_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			DELETE("/sections/{section_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// --- Articles ---
	Method("CreateArticle", func() {
		Description("Create a new article in a section (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "section_id", Int64, "Section unique identifier")
			Field(3, "title", String, "Article title", func() {
				Example("What is Go?")
				MinLength(3)
				MaxLength(100)
			})
			Field(4, "content", String, "Article content", func() {
				Example("Go is an open source programming language...")
				MinLength(10)
			})
			Required("session_token", "section_id", "title", "content")
		})
		Result(SimpleResponse)
		HTTP(func() {
			POST("/sections/{section_id}/articles")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("GetArticle", func() {
		Description("Get article details by ID. Automatically marks the article as completed when viewed by a student.")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "article_id", Int64, "Article unique identifier")
			Required("session_token", "article_id")
		})
		Result(Article)
		HTTP(func() {
			GET("/articles/{article_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("ListArticles", func() {
		Description("List all articles for a section")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "section_id", Int64, "Section unique identifier")
			Required("session_token", "section_id")
		})
		Result(ArrayOf(Article))
		HTTP(func() {
			GET("/sections/{section_id}/articles")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("UpdateArticle", func() {
		Description("Update article details (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "article_id", Int64, "Article unique identifier")
			Field(3, "title", String, "Article title", func() {
				Example("Updated Article Title")
				MinLength(3)
				MaxLength(100)
			})
			Field(4, "content", String, "Article content", func() {
				Example("Updated article content...")
				MinLength(10)
			})
			Required("session_token", "article_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			PATCH("/articles/{article_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("DeleteArticle", func() {
		Description("Delete an article (teachers only)")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "article_id", Int64, "Article unique identifier")
			Required("session_token", "article_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			DELETE("/articles/{article_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	// --- Article Progress Methods ---
	Method("MarkArticleCompleted", func() {
		Description("Mark an article as completed by a user")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "article_id", Int64, "Article unique identifier")
			Required("session_token", "article_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			POST("/articles/{article_id}/complete")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		
		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("UnmarkArticleCompleted", func() {
		Description("Unmark an article as completed by a user")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "article_id", Int64, "Article unique identifier")
			Required("session_token", "article_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			DELETE("/articles/{article_id}/complete")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		
		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("CheckArticleCompleted", func() {
		Description("Check if an article is completed by a user")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "article_id", Int64, "Article unique identifier")
			Required("session_token", "article_id")
		})
		Result(CheckArticleCompletedResult)
		HTTP(func() {
			GET("/articles/{article_id}/completed")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		
		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	// --- Course Content and Progress Methods ---
	Method("GetCourseContent", func() {
		Description("Get all sections and articles for a course with complete structure")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier")
			Required("session_token", "course_id")
		})
		Result(CourseContent)
		HTTP(func() {
			GET("/courses/{course_id}/content")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		
		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("GetUserCourseProgress", func() {
		Description("Get sections and articles completed by a user for a specific course")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier")
			Field(3, "user_id", Int64, "User unique identifier (optional, if not provided uses session user)")
			Required("session_token", "course_id")
		})
		
		Result(UserCourseProgress)

		HTTP(func() {
			GET("/courses/{course_id}/progress")
			Cookie("session_token:session")
			Param("user_id")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		
		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("GetCourseCompletionStats", func() {
		Description("Get course completion statistics for a user")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier")
			Required("session_token", "course_id")
		})

		Result(CourseCompletionStats)

		HTTP(func() {
			GET("/courses/{course_id}/completion")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})

	Method("GetCourseLeaderboard", func() {
		Description("Get course leaderboard with users who completed the most articles")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier")
			Field(3, "limit", Int64, "Maximum number of results (optional, default 10)", func() {
				Default(10)
				Minimum(1)
				Maximum(100)
			})
			Required("session_token", "course_id")
		})

		Result(CourseLeaderboard)

		HTTP(func() {
			GET("/courses/{course_id}/leaderboard")
			Cookie("session_token:session")
			Param("limit")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("unauthorized", CodePermissionDenied)
			Response("not_found", CodeNotFound)
		})
	})
})