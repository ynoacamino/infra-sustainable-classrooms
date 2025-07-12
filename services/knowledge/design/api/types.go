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
var TestResponse = Type("TestResponse", func() {
	Description("Response for test creation and management")

	Field(1, "test_id", Int64, "Unique test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "category_id", Int32, "Test category ID")
	Field(5, "category_name", String, "Test category name")
	Field(6, "difficulty_level", String, "Difficulty level")
	Field(7, "duration_minutes", Int, "Duration in minutes")
	Field(8, "passing_score", Float64, "Minimum passing score")
	Field(9, "is_active", Boolean, "Whether test is active")
	Field(4, "created_at", Int64, "Creation timestamp (Unix)")
	Field(11, "expires_at", String, "Expiration timestamp")
	Field(12, "instructions", String, "Special instructions for the test")
	Field(13, "created_by", Int64, "Creator user ID")
	Field(14, "updated_at", String, "Last update timestamp")
	Field(15, "total_questions", Int, "Total number of questions in test")

	Required("test_id", "title", "category_id", "duration_minutes", "is_active", "created_at", "created_by", "total_questions")
})

var TestQuestion = Type("TestQuestion", func() {
	Description("Question specific to a test")

	Field(1, "question_id", Int64, "Unique question identifier")
	Field(2, "test_id", Int64, "Test this question belongs to")
	Field(3, "question_text", String, "Question text")
	Field(4, "options", ArrayOf(String), "Multiple choice options")
	Field(5, "correct_answer", Int, "Index of correct answer (0-based)")
	Field(6, "explanation", String, "Explanation for correct answer")
	Field(7, "points", Int, "Points awarded for correct answer")
	Field(8, "question_order", Int, "Question order in test")
	Field(9, "created_at", Int64, "Creation timestamp (Unix)")

	Required("question_id", "test_id", "question_text", "options", "correct_answer", "points", "question_order")
})

var QuestionResponse = Type("QuestionResponse", func() {
	Description("Response for question creation and management")

	Field(1, "question_id", Int64, "Question identifier")
	Field(2, "test_id", Int64, "Test this question belongs to")
	Field(3, "question_text", String, "Question text")
	Field(4, "options", ArrayOf(String), "Multiple choice options")
	Field(5, "correct_answer", Int, "Index of correct answer")
	Field(6, "explanation", String, "Explanation for correct answer")
	Field(7, "points", Int, "Points awarded")
	Field(8, "question_order", Int, "Question order in test")
	Field(9, "created_at", Int64, "Creation timestamp (Unix)")
	Field(10, "updated_at", Int64, "Last update timestamp (Unix)")

	Required("question_id", "test_id", "question_text", "options", "correct_answer", "points", "question_order", "created_at")
})

var TestQuestionsResponse = Type("TestQuestionsResponse", func() {
	Description("Response containing all questions for a test")

	Field(1, "test_id", Int64, "Test identifier")
	Field(2, "test_title", String, "Test title")
	Field(3, "questions", ArrayOf(TestQuestion), "List of questions")
	Field(4, "total_questions", Int, "Total number of questions")

	Required("test_id", "test_title", "questions", "total_questions")
})

var MyTestsResponse = Type("MyTestsResponse", func() {
	Description("Response containing user's created tests")

	Field(1, "tests", ArrayOf(TestResponse), "List of tests")
	Field(2, "total_tests", Int, "Total number of tests")
	Field(3, "page", Int, "Current page number")
	Field(4, "limit", Int, "Items per page")
	Field(5, "total_pages", Int, "Total number of pages")

	Required("tests", "total_tests", "page", "limit", "total_pages")
})

var TestFormResponse = Type("TestFormResponse", func() {
	Description("Test form with questions for student to take")

	Field(1, "test_id", Int64, "Test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "duration_minutes", Int, "Duration in minutes")
	Field(5, "total_questions", Int, "Total number of questions")
	Field(6, "questions", ArrayOf(TestQuestion), "Test questions")
	Field(7, "instructions", String, "Test instructions")
	Field(8, "passing_score", Float64, "Minimum passing score")
	Field(9, "expires_at", Int64, "Test expiration timestamp (Unix))")

	Required("test_id", "title", "duration_minutes", "total_questions", "questions")
})

var AnswerSubmission = Type("AnswerSubmission", func() {
	Description("Student answer submission for a question")

	Field(1, "question_id", Int64, "Question identifier")
	Field(2, "selected_answer", Int, "Selected option index (0-based)")

	Required("question_id", "selected_answer")
})

var TestSubmissionResponse = Type("TestSubmissionResponse", func() {
	Description("Response after test submission")

	Field(1, "submission_id", Int64, "Unique submission identifier")
	Field(2, "test_id", Int64, "Test identifier")
	Field(3, "user_id", Int64, "User identifier")
	Field(4, "score", Float64, "Test score (0-100)")
	Field(5, "total_points", Int, "Total points earned")
	Field(6, "max_points", Int, "Maximum possible points")
	Field(7, "correct_answers", Int, "Number of correct answers")
	Field(8, "total_questions", Int, "Total number of questions")
	Field(9, "time_taken", Int, "Time taken in seconds")
	Field(10, "passed", Boolean, "Whether test was passed")
	Field(11, "submitted_at", Int64, "Submissionubmission timestamp (Unix)")

	Required("submission_id", "test_id", "user_id", "score", "total_points", "max_points",
		"correct_answers", "total_questions", "passed", "submitted_at")
})

var QuestionResult = Type("QuestionResult", func() {
	Description("Detailed result for a specific question")

	Field(1, "question_id", Int64, "Question identifier")
	Field(2, "question_text", String, "Question text")
	Field(3, "options", ArrayOf(String), "Multiple choice options")
	Field(4, "correct_answer", Int, "Index of correct answer")
	Field(5, "student_answer", Int, "Student's selected answer")
	Field(6, "is_correct", Boolean, "Whether answer was correct")
	Field(7, "points_earned", Int, "Points earned for this question")
	Field(8, "max_points", Int, "Maximum points for this question")
	Field(9, "explanation", String, "Explanation for correct answer")

	Required("question_id", "question_text", "options", "correct_answer",
		"student_answer", "is_correct", "points_earned", "max_points")
})

var TestResultResponse = Type("TestResultResponse", func() {
	Description("Detailed test results with question-by-question breakdown")

	Field(1, "submission_id", Int64, "Submission identifier")
	Field(2, "test_id", Int64, "Test identifier")
	Field(3, "test_title", String, "Test title")
	Field(4, "user_id", Int64, "User identifier")
	Field(5, "score", Float64, "Overall score (0-100)")
	Field(6, "total_points", Int, "Total points earned")
	Field(7, "max_points", Int, "Maximum possible points")
	Field(8, "correct_answers", Int, "Number of correct answers")
	Field(9, "total_questions", Int, "Total number of questions")
	Field(10, "time_taken", Int, "Time taken in seconds")
	Field(11, "passed", Boolean, "Whether test was passed")
	Field(12, "question_results", ArrayOf(QuestionResult), "Detailed results per question")
	Field(13, "submitted_at", Int64, "Submissionubmission timestamp (Unix) (Unix)")
	Field(15, "category_id", Int32, "Test category ID")
	Field(16, "category_name", String, "Test category name")

	Required("submission_id", "test_id", "test_title", "user_id", "score",
		"total_points", "max_points", "correct_answers", "total_questions",
		"passed", "question_results", "submitted_at")
})

// === PARTICIPANT AND USER TYPES ===
var Participant = Type("Participant", func() {
	Description("Test participant information")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "username", String, "Username")
	Field(3, "email", String, "User email")
	Field(4, "submission_id", Int64, "Submission identifier (if submitted)")
	Field(5, "score", Float64, "Test score (if completed)")
	Field(6, "status", String, "Status (not_started, in_progress, completed, expired)")
	Field(7, "started_at", Int64, "Test start timestamp (Unix)")
	Field(8, "submitted_at", Int64, "Submission timestamp (Unix)")
	Field(9, "time_taken", Int, "Time taken in seconds")
	Field(10, "passed", Boolean, "Whether test was passed")

	Required("user_id", "username", "status")
})

var ParticipantsResponse = Type("ParticipantsResponse", func() {
	Description("Response containing list of test participants")

	Field(1, "test_id", Int64, "Test identifier")
	Field(2, "test_title", String, "Test title")
	Field(3, "participants", ArrayOf(Participant), "List of participants")
	Field(4, "total_participants", Int, "Total number of participants")
	Field(5, "completed_count", Int, "Number of completed submissions")
	Field(6, "passed_count", Int, "Number of passed submissions")
	Field(7, "average_score", Float64, "Average score of completed tests")
	Field(8, "page", Int, "Current page number")
	Field(9, "limit", Int, "Items per page")
	Field(10, "total_pages", Int, "Total number of pages")

	Required("test_id", "test_title", "participants", "total_participants",
		"completed_count", "passed_count", "page", "limit", "total_pages")
})

var UserSubmission = Type("UserSubmission", func() {
	Description("User's test submission summary")

	Field(1, "submission_id", Int64, "Submission identifier")
	Field(2, "test_id", Int64, "Test identifier")
	Field(3, "test_title", String, "Test title")
	Field(4, "category_id", Int32, "Test category ID")
	Field(5, "category_name", String, "Test category name")
	Field(6, "score", Float64, "Test score")
	Field(7, "passed", Boolean, "Whether test was passed")
	Field(8, "time_taken", Int, "Time taken in seconds")
	Field(9, "submitted_at", Int64, "Submission timestamp (Unix)Unix)")
	Field(10, "difficulty_level", String, "Test difficulty level")

	Required("submission_id", "test_id", "test_title", "category_id", "category_name",
		"score", "passed", "submitted_at")
})

var UserSubmissionsResponse = Type("UserSubmissionsResponse", func() {
	Description("Response containing user's test submissions")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "submissions", ArrayOf(UserSubmission), "List of submissions")
	Field(3, "total_submissions", Int, "Total number of submissions")
	Field(4, "average_score", Float64, "Average score across all tests")
	Field(5, "tests_passed", Int, "Number of tests passed")
	Field(6, "tests_failed", Int, "Number of tests failed")
	Field(7, "page", Int, "Current page number")
	Field(8, "limit", Int, "Items per page")
	Field(9, "total_pages", Int, "Total number of pages")

	Required("user_id", "submissions", "total_submissions", "average_score",
		"tests_passed", "tests_failed", "page", "limit", "total_pages")
})

// === QUESTION BANK TYPES (REMOVED - SIMPLIFIED) ===
// We removed the complex question bank system
// Now questions belong directly to tests

// === AVAILABLE TESTS TYPES ===
var AvailableTest = Type("AvailableTest", func() {
	Description("Available test information for students")

	Field(1, "test_id", Int64, "Test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "category_id", Int32, "Test category ID")
	Field(5, "category_name", String, "Test category name")
	Field(6, "difficulty_level", String, "Difficulty level")
	Field(7, "duration_minutes", Int, "Duration in minutes")
	Field(8, "total_questions", Int, "Total number of questions")
	Field(9, "passing_score", Float64, "Minimum passing score")
	Field(11, "last_attempt", Int64, "Last attempt timestamp (Unix) (Unix)")
	Field(12, "expires_at", Int64, "Test expiration timestamp (Unix) (Unix)")

	Required("test_id", "title", "category_id", "difficulty_level", "duration_minutes",
		"total_questions", "passing_score")
})

var AvailableTestsResponse = Type("AvailableTestsResponse", func() {
	Description("Response containing available tests for a student")

	Field(1, "tests", ArrayOf(AvailableTest), "List of available tests")
	Field(2, "total_tests", Int, "Total number of available tests")
	Field(3, "page", Int, "Current page number")
	Field(4, "limit", Int, "Items per page")
	Field(5, "total_pages", Int, "Total number of pages")
	Field(6, "filters", MapOf(String, Any), "Applied filters")

	Required("tests", "total_tests", "page", "limit", "total_pages")
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
	Field(4, "category_id", Int32, "Test category ID")
	Field(5, "category_name", String, "Test category name")
	Field(6, "difficulty_level", String, "Difficulty level")
	Field(7, "duration_minutes", Int, "Duration in minutes")
	Field(8, "total_questions", Int, "Total number of questions")
	Field(9, "passing_score", Float64, "Minimum passing score")
	Field(11, "expires_at", Int64, "Test expiration timestamp (Unix)")
	Field(12, "created_by", String, "Creator name")
	Field(13, "instructions", String, "Special instructions for the test")

	Required("test_id", "title", "category_id", "difficulty_level", "duration_minutes",
		"total_questions", "passing_score")
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

// === TEST CATEGORY TYPES ===
var TestCategory = Type("TestCategory", func() {
	Description("Test category information")

	Field(1, "id", Int32, "Category identifier")
	Field(2, "name", String, "Category name")
	Field(3, "description", String, "Category description")
	Field(4, "created_at", Int64, "Creation timestamp (Unix)")
	Field(5, "test_count", Int, "Number of tests in this category")
	Field(6, "active_test_count", Int, "Number of active tests in this category")

	Required("id", "name", "created_at")
})

var TestCategoriesResponse = Type("TestCategoriesResponse", func() {
	Description("Response containing test categories")

	Field(1, "categories", ArrayOf(TestCategory), "List of categories")
	Field(2, "total_categories", Int, "Total number of categories")

	Required("categories", "total_categories")
})
