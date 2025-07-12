package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("knowledge", func() {
	Description("Knowledge microservice for managing MCQ tests, validations, and grading")

	HTTP(func() {
		Path("/knowledge")
	})

	GRPC(func() {

	})

	Error("invalid_input", String, "Invalid input parameters")
	Error("unauthorized", String, "Unauthorized access")
	Error("test_not_found", String, "Test not found")
	Error("question_not_found", String, "Question not found")
	Error("submission_not_found", String, "Submission not found")
	Error("test_already_submitted", String, "Test already submitted by user")
	Error("test_expired", String, "Test has expired")
	Error("insufficient_permissions", String, "Insufficient permissions")

	Method("CreateTest", func() {
		Description("Create a new MCQ test form")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "title", String, "Test title")
			Field(3, "description", String, "Test description")
			Field(4, "category_id", Int32, "Test category ID")
			Field(5, "difficulty_level", String, "Difficulty level (easy, medium, hard)")
			Field(6, "duration_minutes", Int, "Test duration in minutes")
			Field(7, "passing_score", Float64, "Minimum passing score (0-100)")
			Field(8, "is_active", Boolean, "Whether test is active")
			Field(9, "expires_at", String, "Test expiration date (ISO format)")

			Required("session_token", "title", "duration_minutes", "category_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/tests")
			Cookie("session_token:session")

			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("GetMyTestsCreated", func() {
		Description("Get all tests created by the current teacher")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "page", Int, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Field(3, "limit", Int, "Items per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token")
		})

		Result(MyTestsResponse)

		HTTP(func() {
			GET("/my-tests")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("GetAvailableTests", func() {
		Description("Get available tests for students to take")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(4, "page", Int, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Field(5, "limit", Int, "Items per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token")
		})

		Result(AvailableTestsResponse)

		HTTP(func() {
			GET("/tests/available")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("UpdateTest", func() {
		Description("Update test information")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID to update")
			Field(3, "title", String, "Updated title")
			Field(4, "description", String, "Updated description")
			Field(5, "category_id", Int32, "Updated category ID")
			Field(6, "duration_minutes", Int, "Updated duration")
			Field(7, "passing_score", Float64, "Updated passing score")
			Field(8, "is_active", Boolean, "Updated active status")
			Field(9, "expires_at", String, "Updated expiration date")

			Required("session_token", "test_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			PUT("/tests/{test_id}")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("DeleteTest", func() {
		Description("Delete a test and all its questions")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID to delete")

			Required("session_token", "test_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			DELETE("/tests/{test_id}")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("GetTestQuestions", func() {
		Description("Get all questions for a specific test (for editing)")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID")

			Required("session_token", "test_id")
		})

		Result(TestQuestionsResponse)

		HTTP(func() {
			GET("/tests/{test_id}/questions")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("AddQuestionToTest", func() {
		Description("Add a new question directly to a specific test")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID to add question to")
			Field(3, "question_text", String, "Question text")
			Field(4, "options", ArrayOf(String), "Multiple choice options")
			Field(5, "correct_answer", Int, "Index of correct answer (0-based)")
			Field(6, "explanation", String, "Explanation for the correct answer")
			Field(7, "points", Int, "Points awarded for correct answer", func() {
				Default(1)
				Minimum(1)
			})
			Field(8, "question_order", Int, "Question order in test", func() {
				Default(1)
				Minimum(1)
			})

			Required("session_token", "test_id", "question_text", "options", "correct_answer")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/tests/{test_id}/questions")
			Cookie("session_token:session")

			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("UpdateTestQuestion", func() {
		Description("Update a specific question in a test")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "question_id", Int64, "Question ID to update")
			Field(4, "question_text", String, "Updated question text")
			Field(5, "options", ArrayOf(String), "Updated multiple choice options")
			Field(6, "correct_answer", Int, "Updated index of correct answer (0-based)")
			Field(7, "explanation", String, "Updated explanation")
			Field(8, "points", Int, "Updated points")
			Field(9, "question_order", Int, "Updated question order")

			Required("session_token", "test_id", "question_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			PUT("/tests/{test_id}/questions/{question_id}")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("question_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("DeleteTestQuestion", func() {
		Description("Delete a specific question from a test")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "question_id", Int64, "Question ID to delete")

			Required("session_token", "test_id", "question_id")
		})

		Result(SimpleResponse)

		HTTP(func() {
			DELETE("/tests/{test_id}/questions/{question_id}")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("question_not_found", StatusNotFound)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("GetTestForm", func() {
		Description("Get MCQ test form for student to take")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID")

			Required("session_token", "test_id")
		})

		Result(TestFormResponse)

		HTTP(func() {
			GET("/tests/{test_id}/form")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("test_expired", StatusGone)
		})
	})

	Method("SubmitTest", func() {
		Description("Submit MCQ test answers for validation and grading")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "answers", ArrayOf(AnswerSubmission), "Student answers")

			Required("session_token", "test_id", "answers")
		})

		Result(TestSubmissionResponse)

		HTTP(func() {
			POST("/tests/{test_id}/submit")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("test_already_submitted", StatusConflict)
			Response("test_expired", StatusGone)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("GetTestResults", func() {
		Description("Get test results and detailed grading")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "submission_id", Int64, "Submission ID")

			Required("session_token", "submission_id")
		})

		Result(TestResultResponse)

		HTTP(func() {
			GET("/submissions/{submission_id}/results")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("submission_not_found", StatusNotFound)
		})
	})

	Method("GetTestParticipants", func() {
		Description("Get list of participants for a specific test")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "page", Int, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Field(4, "limit", Int, "Items per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token", "test_id")
		})

		Result(ParticipantsResponse)

		HTTP(func() {
			GET("/tests/{test_id}/participants")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	Method("GetUserSubmissions", func() {
		Description("Get all test submissions for a user")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "user_id", Int64, "User ID (optional, defaults to authenticated user)")
			Field(3, "page", Int, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Field(4, "limit", Int, "Items per page", func() {
				Default(20)
				Minimum(1)
				Maximum(100)
			})

			Required("session_token")
		})

		Result(UserSubmissionsResponse)

		HTTP(func() {
			GET("/users/{user_id}/submissions")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("GetTestPreview", func() {
		Description("Get test preview information before taking it")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID to preview")

			Required("session_token", "test_id")
		})

		Result(TestPreviewResponse)

		HTTP(func() {
			GET("/tests/{test_id}/preview")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
		})
	})

	// === BULK OPERATIONS ===
	Method("BulkAddQuestions", func() {
		Description("Add multiple questions to a test at once")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "test_id", Int64, "Test ID to add questions to")
			Field(3, "questions", ArrayOf(BulkQuestionInput), "Array of questions to add")

			Required("session_token", "test_id", "questions")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/tests/{test_id}/questions/bulk")
			Cookie("session_token:session")

			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
			Response("insufficient_permissions", StatusForbidden)
		})
	})

	// === CATEGORY MANAGEMENT ===
	Method("GetCategories", func() {
		Description("Get all test categories")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")

			Required("session_token")
		})

		Result(TestCategoriesResponse)

		HTTP(func() {
			GET("/categories")
			Cookie("session_token:session")

			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("CreateCategory", func() {
		Description("Create a new test category (teacher only)")

		Payload(func() {
			Field(1, "session_token", String, "Authentication session token")
			Field(2, "name", String, "Category name")
			Field(3, "description", String, "Category description")

			Required("session_token", "name")
		})

		Result(SimpleResponse)

		HTTP(func() {
			POST("/categories")
			Cookie("session_token:session")

			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("invalid_input", StatusBadRequest)
			Response("insufficient_permissions", StatusForbidden)
		})
	})
})
