package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("stats", func() {
	Description("Statistics service for tracking user progress in courses, sections, and articles. Provides comprehensive analytics and progress tracking.")

	HTTP(func() {
		Path("/stats")
	})

	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	Error("invalid_input", String, "Invalid input parameters")
	Error("rate_limited", String, "Too many requests")
	Error("service_unavailable", String, "Service temporarily unavailable")
	Error("unauthorized", String, "Unauthorized access")
	Error("not_found", String, "Resource not found")
	Error("internal_error", String, "Internal server error")

	// --- User Progress Statistics ---
	Method("GetUserCourseProgress", func() {
		Description("Get detailed progress information for a user in a specific course")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User unique identifier", func() {
				Example(123)
			})
			Field(3, "course_id", Int64, "Course unique identifier", func() {
				Example(1)
			})
			Required("session_token", "user_id", "course_id")
		})
		Result(CourseProgress)
		HTTP(func() {
			GET("/users/{user_id}/courses/{course_id}/progress")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		GRPC(func() {
			// gRPC method for microservice communication
		})
	})

	Method("GetUserOverallStats", func() {
		Description("Get overall learning statistics for a user across all courses")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User unique identifier", func() {
				Example(123)
			})
			Required("session_token", "user_id")
		})
		Result(UserOverallStats)
		HTTP(func() {
			GET("/users/{user_id}/stats")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		GRPC(func() {
			// gRPC method for microservice communication
		})
	})

	Method("GetUserCompletedArticles", func() {
		Description("Get list of all completed articles for a user")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User unique identifier", func() {
				Example(123)
			})
			Field(3, "limit", Int64, "Maximum number of articles to return (optional, default 50)", func() {
				Example(20)
				Minimum(1)
				Maximum(100)
			})
			Required("session_token", "user_id")
		})
		Result(ArrayOf(ArticleProgress))
		HTTP(func() {
			GET("/users/{user_id}/completed-articles")
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
			// gRPC method for microservice communication
		})
	})

	Method("GetCourseLeaderboard", func() {
		Description("Get leaderboard for a specific course showing top performing users")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier", func() {
				Example(1)
			})
			Field(3, "limit", Int64, "Maximum number of users to return (optional, default 10)", func() {
				Example(20)
				Minimum(1)
				Maximum(50)
			})
			Required("session_token", "course_id")
		})
		Result(ArrayOf(LeaderboardEntry))
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
			// gRPC method for microservice communication
		})
	})
})
