package design

import (
	. "goa.design/goa/v3/dsl"
)

// CodeLabService defines the codelab microservice
var _ = Service("codelab", func() {
	Description("Codelab microservice for coding exercises, tests, answers and attempts")

	// HTTP transport for client communication only
	HTTP(func() {
		Path("/api/codelab")
	})

	// gRPC transport for inter-service communication
	GRPC(func() {
		// gRPC service configuration for microservice communication
	})

	// Global error definitions for the service

	Error("invalid_input", String, "Invalid input parameters")
	Error("rate_limited", String, "Too many requests")
	Error("service_unavailable", String, "Service temporarily unavailable")
	Error("unauthorized", String, "Unauthorized access")
	Error("permission_denied", String, "Permission denied (only teachers)")
	Error("not_found", String, "Resource not found")
	Error("internal_error", String, "Internal server error")

	// ========================================
	// EXERCISE CRUD ENDPOINTS (for professors)
	// ========================================

	Method("CreateExercise", func() {
		Description("Create a new coding exercise (professors only)")

		Payload(CreateExercisePayload)

		Result(SimpleResponse)

		HTTP(func() {
			POST("/exercises")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
		})
	})

	Method("GetExercise", func() {
		Description("Get exercise by ID with solution (professors only)")

		Payload(func() {
			Field(1, "id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(2, "session_token", String, "Authentication session token")

			Required("session_token", "id")
		})

		Result(Exercise)

		HTTP(func() {
			GET("/exercises/{id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
		})
	})

	Method("ListExercises", func() {
		Description("List all exercises with solutions (professors only)")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Required("session_token")
		})

		Result(ArrayOf(Exercise))

		HTTP(func() {
			GET("/exercises")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
		})
	})

	Method("UpdateExercise", func() {
		Description("Update an exercise (professors only)")

		Payload(func() {
			Field(1, "id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(2, "exercise", UpdateExercisePayload, "Exercise data to update")
			Field(3, "session_token", String, "Authentication session token")
			Required("session_token", "id", "exercise")
		})

		Result(SimpleResponse)

		HTTP(func() {
			PUT("/exercises/{id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("DeleteExercise", func() {
		Description("Delete an exercise (professors only)")

		Payload(func() {
			Field(1, "id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(2, "session_token", String, "Authentication session token")
			Required("session_token", "id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			DELETE("/exercises/{id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// ========================================
	// TEST CRUD ENDPOINTS (for professors)
	// ========================================

	Method("CreateTest", func() {
		Description("Create a new test case for an exercise (professors only)")

		Payload(CreateTestPayload)

		Result(SimpleResponse)

		HTTP(func() {
			POST("/tests")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
		})
	})

	Method("GetTestsByExercise", func() {
		Description("Get all test cases for an exercise (professors only)")

		Payload(func() {
			Field(1, "exercise_id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(2, "session_token", String, "Authentication session token")
			Required("session_token", "exercise_id")
		})

		Result(ArrayOf(Test))

		HTTP(func() {
			GET("/exercises/{exercise_id}/tests")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("UpdateTest", func() {
		Description("Update a test case (professors only)")

		Payload(func() {
			Field(1, "id", Int64, "Test ID", func() {
				Example(1)
			})
			Field(2, "test", UpdateTestPayload, "Test data to update")
			Field(3, "session_token", String, "Authentication session token")
			Required("session_token", "id", "test")
		})

		Result(SimpleResponse)

		HTTP(func() {
			PUT("/tests/{id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
		})
	})

	Method("DeleteTest", func() {
		Description("Delete a test case (professors only)")

		Payload(func() {
			Field(1, "id", Int64, "Test ID", func() {
				Example(1)
			})
			Field(2, "session_token", String, "Authentication session token")

			Required("session_token", "id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			DELETE("/tests/{id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// ========================================
	// STUDENT ENDPOINTS (read exercises, submit attempts)
	// ========================================

	Method("GetExerciseForStudent", func() {
		Description("Get exercise by ID without solution (students)")

		Payload(func() {
			Field(1, "id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(2, "session_token", String, "Authentication session token")

			Required("session_token", "id")
		})

		Result(ExerciseForStudents)

		HTTP(func() {
			GET("/student/exercises/{id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("ListExercisesForStudents", func() {
		Description("List all exercises without solutions (students)")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Required("session_token")
		})

		Result(ArrayOf(ExerciseForStudentsListView))

		HTTP(func() {
			GET("/student/exercises")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("CreateAttempt", func() {
		Description("Submit a code attempt for an exercise (students)")

		Payload(CreateAttemptPayload)

		Result(SimpleResponse)

		HTTP(func() {
			POST("/attempts")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("invalid_input", StatusBadRequest)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
		})
	})

	Method("GetAttemptsByUserAndExercise", func() {
		Description("Get user's attempts for a specific exercise (students)")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID", func() {
				Example(123)
			})
			Field(2, "exercise_id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(3, "session_token", String, "Authentication session token")

			Required("user_id", "exercise_id", "session_token")
		})

		Result(ArrayOf(Attempt))

		HTTP(func() {
			GET("/student/users/{user_id}/exercises/{exercise_id}/attempts")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})

	// ========================================
	// ANSWER MANAGEMENT (internal/helper endpoints)
	// ========================================

	Method("GetAnswerByUserAndExercise", func() {
		Description("Get user's answer for a specific exercise")

		Payload(func() {
			Field(1, "user_id", Int64, "User ID", func() {
				Example(123)
			})
			Field(2, "exercise_id", Int64, "Exercise ID", func() {
				Example(1)
			})
			Field(3, "session_token", String, "Authentication session token")

			Required("user_id", "exercise_id", "session_token")
		})

		Result(Answer)

		HTTP(func() {
			GET("/answers/user/{user_id}/exercise/{exercise_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("unauthorized", StatusUnauthorized)
			Response("service_unavailable", StatusServiceUnavailable)
			Response("permission_denied", StatusForbidden)
			Response("invalid_input", StatusBadRequest)
		})
	})
})
