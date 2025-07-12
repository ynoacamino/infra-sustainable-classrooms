package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("knowledge", func() {
	Description("Knowledge microservice for simple test forms")

	HTTP(func() {
		Path("/knowledge")
	})

	GRPC(func() {})

	Error("unauthorized", String, "Unauthorized access")
	Error("test_not_found", String, "Test not found")
	Error("question_not_found", String, "Question not found")
	Error("submission_not_found", String, "Submission not found")
	Error("test_already_submitted", String, "Test already submitted by user")
	Error("invalid_input", String, "Invalid input")

	// === TEACHER METHODS ===
	Method("CreateTest", func() {
		Description("Create a new test form")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "title", String, "Test title")
			Required("session_token", "title")
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

	Method("GetMyTests", func() {
		Description("Get my created tests")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Required("session_token")
		})
		Result(TestsResponse)
		HTTP(func() {
			GET("/tests/my")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("UpdateTest", func() {
		Description("Update test title")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "title", String, "New title")
			Required("session_token", "test_id", "title")
		})
		Result(SimpleResponse)
		HTTP(func() {
			PUT("/tests/{test_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
		})
	})

	Method("DeleteTest", func() {
		Description("Delete a test")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Required("session_token", "test_id")
		})
		Result(SimpleResponse)
		HTTP(func() {
			DELETE("/tests/{test_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
		})
	})

	Method("GetTestQuestions", func() {
		Description("Get questions for a test")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Required("session_token", "test_id")
		})
		Result(QuestionsResponse)
		HTTP(func() {
			GET("/tests/{test_id}/questions")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
		})
	})

	Method("AddQuestion", func() {
		Description("Add a question to a test")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "question_text", String, "Question text")
			Field(4, "option_a", String, "Option A")
			Field(5, "option_b", String, "Option B")
			Field(6, "option_c", String, "Option C")
			Field(7, "option_d", String, "Option D")
			Field(8, "correct_answer", Int, "Correct answer (0=A, 1=B, 2=C, 3=D)")
			Required("session_token", "test_id", "question_text", "option_a", "option_b", "option_c", "option_d", "correct_answer")
		})
		Result(SimpleResponse)
		HTTP(func() {
			POST("/tests/{test_id}/questions")
			Cookie("session_token:session")
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("UpdateQuestion", func() {
		Description("Update a question")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "question_id", Int64, "Question ID")
			Field(4, "question_text", String, "Question text")
			Field(5, "option_a", String, "Option A")
			Field(6, "option_b", String, "Option B")
			Field(7, "option_c", String, "Option C")
			Field(8, "option_d", String, "Option D")
			Field(9, "correct_answer", Int, "Correct answer (0=A, 1=B, 2=C, 3=D)")
			Required("session_token", "test_id", "question_id", "question_text", "option_a", "option_b", "option_c", "option_d", "correct_answer")
		})
		Result(SimpleResponse)
		HTTP(func() {
			PUT("/tests/{test_id}/questions/{question_id}")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("question_not_found", StatusNotFound)
		})
	})

	Method("DeleteQuestion", func() {
		Description("Delete a question")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "question_id", Int64, "Question ID")
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
		})
	})

	// === STUDENT METHODS ===
	Method("GetAvailableTests", func() {
		Description("Get available tests for students")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Required("session_token")
		})
		Result(TestsResponse)
		HTTP(func() {
			GET("/tests/available")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("GetTestForm", func() {
		Description("Get test form for taking")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Required("session_token", "test_id")
		})
		Result(FormResponse)
		HTTP(func() {
			GET("/tests/{test_id}/form")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("test_already_submitted", StatusConflict)
		})
	})

	Method("SubmitTest", func() {
		Description("Submit test answers")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "test_id", Int64, "Test ID")
			Field(3, "answers", ArrayOf(Answer), "Answer submissions")
			Required("session_token", "test_id", "answers")
		})
		Result(SubmitResponse)
		HTTP(func() {
			POST("/tests/{test_id}/submit")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("test_not_found", StatusNotFound)
			Response("test_already_submitted", StatusConflict)
			Response("invalid_input", StatusBadRequest)
		})
	})

	Method("GetMySubmissions", func() {
		Description("Get my test submissions")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Required("session_token")
		})
		Result(SubmissionsResponse)
		HTTP(func() {
			GET("/submissions/my")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("GetSubmissionResult", func() {
		Description("Get detailed submission result")
		Payload(func() {
			Field(1, "session_token", String, "Session token")
			Field(2, "submission_id", Int64, "Submission ID")
			Required("session_token", "submission_id")
		})
		Result(SubmissionResult)
		HTTP(func() {
			GET("/submissions/{submission_id}/result")
			Cookie("session_token:session")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("submission_not_found", StatusNotFound)
		})
	})
})
