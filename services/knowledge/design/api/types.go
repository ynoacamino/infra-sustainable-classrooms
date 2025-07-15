package design

import (
	. "goa.design/goa/v3/dsl"
)

// === BASIC RESPONSE TYPES ===
var SimpleResponse = Type("SimpleResponse", func() {
	Description("Basic response with success status and message")
	Field(1, "success", Boolean, "Operation success status")
	Field(2, "message", String, "Response message")
	Required("success", "message")
})

// === TEST RELATED TYPES ===
var Test = Type("Test", func() {
	Description("Test/Form information")
	Field(1, "id", Int64, "Test ID")
	Field(2, "title", String, "Test title")
	Field(3, "created_by", Int64, "Creator user ID")
	Field(4, "created_at", Int64, "Creation timestamp")
	Field(5, "question_count", Int, "Number of questions")
	Required("id", "title", "created_by", "created_at")
})

var Question = Type("Question", func() {
	Description("Question information")
	Field(1, "id", Int64, "Question ID")
	Field(2, "test_id", Int64, "Test ID")
	Field(3, "question_text", String, "Question text")
	Field(4, "option_a", String, "Option A")
	Field(5, "option_b", String, "Option B")
	Field(6, "option_c", String, "Option C")
	Field(7, "option_d", String, "Option D")
	Field(8, "correct_answer", Int, "Correct answer (0=A, 1=B, 2=C, 3=D)")
	Field(9, "question_order", Int, "Question order")
	Required("id", "test_id", "question_text", "option_a", "option_b", "option_c", "option_d", "correct_answer", "question_order")
})

var QuestionForm = Type("QuestionForm", func() {
	Description("Question for form taking (without correct answer)")
	Field(1, "id", Int64, "Question ID")
	Field(2, "question_text", String, "Question text")
	Field(3, "option_a", String, "Option A")
	Field(4, "option_b", String, "Option B")
	Field(5, "option_c", String, "Option C")
	Field(6, "option_d", String, "Option D")
	Field(7, "question_order", Int, "Question order")
	Required("id", "question_text", "option_a", "option_b", "option_c", "option_d", "question_order")
})

var Answer = Type("Answer", func() {
	Description("Answer submission")
	Field(1, "question_id", Int64, "Question ID")
	Field(2, "selected_answer", Int, "Selected answer (0=A, 1=B, 2=C, 3=D)")
	Required("question_id", "selected_answer")
})

var Submission = Type("Submission", func() {
	Description("Test submission")
	Field(1, "id", Int64, "Submission ID")
	Field(2, "test_id", Int64, "Test ID")
	Field(3, "test_title", String, "Test title")
	Field(4, "score", Float64, "Score percentage")
	Field(5, "submitted_at", Int64, "Submission timestamp")
	Required("id", "test_id", "test_title", "score", "submitted_at")
})

var SubmissionResult = Type("SubmissionResult", func() {
	Description("Detailed submission result")
	Field(1, "submission", Submission, "Submission info")
	Field(2, "questions", ArrayOf(QuestionResult), "Question results")
	Required("submission", "questions")
})

var QuestionResult = Type("QuestionResult", func() {
	Description("Question result with user answer")
	Field(1, "question", Question, "Question info")
	Field(2, "selected_answer", Int, "User selected answer")
	Field(3, "is_correct", Boolean, "Whether answer was correct")
	Required("question", "selected_answer", "is_correct")
})

// === RESPONSE TYPES ===
var TestResponse = Type("TestResponse", func() {
	Description("Single test response")
	Field(1, "test", Test, "Test information")
	Required("test")
})

var TestsResponse = Type("TestsResponse", func() {
	Description("List of tests")
	Field(1, "tests", ArrayOf(Test), "Tests")
	Required("tests")
})

var QuestionResponse = Type("QuestionResponse", func() {
	Description("Single question response")
	Field(1, "question", Question, "Question information")
	Required("question")
})

var QuestionsResponse = Type("QuestionsResponse", func() {
	Description("List of questions")
	Field(1, "questions", ArrayOf(Question), "Questions")
	Required("questions")
})

var FormResponse = Type("FormResponse", func() {
	Description("Form for taking test")
	Field(1, "test", Test, "Test info")
	Field(2, "questions", ArrayOf(QuestionForm), "Questions")
	Required("test", "questions")
})

var SubmissionResponse = Type("SubmissionResponse", func() {
	Description("Single submission response")
	Field(1, "submission", Submission, "Submission information")
	Required("submission")
})

var SubmissionsResponse = Type("SubmissionsResponse", func() {
	Description("List of submissions")
	Field(1, "submissions", ArrayOf(Submission), "Submissions")
	Required("submissions")
})

var SubmitResponse = Type("SubmitResponse", func() {
	Description("Submit test response")
	Field(1, "success", Boolean, "Success status")
	Field(2, "message", String, "Response message")
	Field(3, "submission_id", Int64, "Submission ID")
	Field(4, "score", Float64, "Score percentage")
	Required("success", "message", "submission_id", "score")
})

// === INTER-SERVICE COMMUNICATION TYPES ===
var UserAccessResponse = Type("UserAccessResponse", func() {
	Description("User access validation response")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "username", String, "Username")
	Field(3, "email", String, "User email")
	Field(4, "permissions", ArrayOf(String), "User permissions")
	Field(5, "roles", ArrayOf(String), "User roles")
	Field(6, "is_active", Boolean, "Whether user is active")
	Field(7, "last_login", String, "Last login timestamp")
	Field(8, "session_valid", Boolean, "Whether session is valid")

	Required("user_id", "username", "permissions", "roles", "is_active", "session_valid")
})

// === NEW TYPES FOR ENHANCED FUNCTIONALITY ===
var TestPreviewResponse = Type("TestPreviewResponse", func() {
	Description("Test preview information before taking it")

	Field(1, "test_id", Int64, "Test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "difficulty_level", String, "Difficulty level")
	Field(5, "duration_minutes", Int32, "Duration in minutes")
	Field(6, "total_questions", Int, "Total number of questions")
	Field(7, "expires_at", Int64, "Test expiration timestamp (Unix)")
	Field(8, "created_by", Int64, "Creator name")

	Required("test_id", "title", "difficulty_level", "duration_minutes",
		"total_questions")
})

var BulkQuestionInput = Type("BulkQuestionInput", func() {
	Description("Input for bulk question creation")

	Field(1, "question_text", String, "Question text")
	Field(2, "options", ArrayOf(String), "Multiple choice options")
	Field(3, "correct_answer", Int, "Index of correct answer (0-based)")
	Field(4, "explanation", String, "Explanation for the correct answer")
	Field(5, "points", Int, "Points awarded for correct answer", func() {
		Default(1)
		Minimum(1)
	})
	Field(6, "order", Int, "Question order in test", func() {
		Default(1)
		Minimum(1)
	})

	Required("question_text", "options", "correct_answer")
})

var BulkQuestionResponse = Type("BulkQuestionResponse", func() {
	Description("Response for bulk question creation")

	Field(1, "test_id", Int64, "Test identifier")
	Field(2, "questions_added", Int, "Number of questions successfully added")
	Field(3, "questions_failed", Int, "Number of questions that failed to add")
	Field(4, "question_ids", ArrayOf(Int64), "List of created question IDs")
	Field(5, "errors", ArrayOf(String), "List of errors for failed questions")
	Field(6, "total_questions", Int, "Total questions in test after bulk add")

	Required("test_id", "questions_added", "questions_failed", "question_ids", "total_questions")
})
