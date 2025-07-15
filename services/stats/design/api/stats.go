package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("stats", func() {
	Description("Service for managing statistics in sustainable classrooms")

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
	Error("permission_denied", String, "Permission denied (only teachers)")
	Error("not_found", String, "Resource not found")
	Error("internal_error", String, "Internal server error")

	// --- Leaderboard Methods ---
	Method("GetCourseLeaderboard", func() {
		Description("Get course completion leaderboard with top performing users")
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
			// gRPC method for microservice communication
		})
	})

	Method("GetUserOverallStats", func() {
		Description("Get overall statistics for a user across all courses")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User unique identifier (optional, if not provided uses session user)")
			Required("session_token")
		})
		Result(UserOverallStats)
		HTTP(func() {
			GET("/users/stats")
			Cookie("session_token:session")
			Param("user_id")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		GRPC(func() {
			// gRPC method for microservice communication
		})
	})

	// --- Progress Tracking Methods ---
	Method("GetUserCourseProgress", func() {
		Description("Get detailed progress for a user in a specific course")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "course_id", Int64, "Course unique identifier")
			Field(3, "user_id", Int64, "User unique identifier (optional, if not provided uses session user)")
			Required("session_token", "course_id")
		})
		Result(UserCourseProgressStats)
		HTTP(func() {
			GET("/courses/{course_id}/progress")
			Cookie("session_token:session")
			Param("user_id")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		GRPC(func() {
			// gRPC method for microservice communication
		})
	})

	Method("GetUserCompletedArticles", func() {
		Description("Get list of articles completed by a user")
		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User unique identifier (optional, if not provided uses session user)")
			Field(3, "limit", Int64, "Maximum number of results (optional, default 20)", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})
			Required("session_token")
		})
		Result(UserCompletedArticles)
		HTTP(func() {
			GET("/users/completed-articles")
			Cookie("session_token:session")
			Param("user_id")
			Param("limit")
			Response(StatusOK)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("permission_denied", StatusForbidden)
			Response("not_found", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
			Response("service_unavailable", StatusServiceUnavailable)
		})
		GRPC(func() {
			// gRPC method for microservice communication
		})
	})
})