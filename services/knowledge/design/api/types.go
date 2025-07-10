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

	Field(1, "test_id", String, "Unique test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "category", String, "Test category")
	Field(5, "difficulty_level", String, "Difficulty level")
	Field(6, "duration_minutes", Int, "Duration in minutes")
	Field(7, "passing_score", Float64, "Minimum passing score")
	Field(8, "is_active", Boolean, "Whether test is active")
	Field(9, "created_at", String, "Creation timestamp")
	Field(10, "expires_at", String, "Expiration timestamp")
	Field(11, "created_by", Int64, "Creator user ID")
	Field(12, "total_questions", Int, "Total number of questions in test")

	Required("test_id", "title", "category", "duration_minutes", "is_active", "created_at", "created_by", "total_questions")
})

var TestQuestion = Type("TestQuestion", func() {
	Description("Question specific to a test")

	Field(1, "question_id", String, "Unique question identifier")
	Field(2, "test_id", String, "Test this question belongs to")
	Field(3, "question_text", String, "Question text")
	Field(4, "options", ArrayOf(String), "Multiple choice options")
	Field(5, "correct_answer", Int, "Index of correct answer (0-based)")
	Field(6, "explanation", String, "Explanation for correct answer")
	Field(7, "points", Int, "Points awarded for correct answer")
	Field(8, "order", Int, "Question order in test")
	Field(9, "created_at", String, "Creation timestamp")

	Required("question_id", "test_id", "question_text", "options", "correct_answer", "points", "order")
})

var QuestionResponse = Type("QuestionResponse", func() {
	Description("Response for question creation and management")

	Field(1, "question_id", String, "Question identifier")
	Field(2, "test_id", String, "Test this question belongs to")
	Field(3, "question_text", String, "Question text")
	Field(4, "options", ArrayOf(String), "Multiple choice options")
	Field(5, "correct_answer", Int, "Index of correct answer")
	Field(6, "explanation", String, "Explanation for correct answer")
	Field(7, "points", Int, "Points awarded")
	Field(8, "order", Int, "Question order in test")
	Field(9, "created_at", String, "Creation timestamp")
	Field(10, "updated_at", String, "Last update timestamp")

	Required("question_id", "test_id", "question_text", "options", "correct_answer", "points", "order", "created_at")
})

var TestQuestionsResponse = Type("TestQuestionsResponse", func() {
	Description("Response containing all questions for a test")

	Field(1, "test_id", String, "Test identifier")
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

	Field(1, "test_id", String, "Test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "duration_minutes", Int, "Duration in minutes")
	Field(5, "total_questions", Int, "Total number of questions")
	Field(6, "questions", ArrayOf(TestQuestion), "Test questions")
	Field(7, "instructions", String, "Test instructions")
	Field(8, "passing_score", Float64, "Minimum passing score")
	Field(9, "expires_at", String, "Test expiration timestamp")

	Required("test_id", "title", "duration_minutes", "total_questions", "questions")
})

var AnswerSubmission = Type("AnswerSubmission", func() {
	Description("Student answer submission for a question")

	Field(1, "question_id", String, "Question identifier")
	Field(2, "selected_option", Int, "Selected option index (0-based)")

	Required("question_id", "selected_option")
})

var TestSubmissionResponse = Type("TestSubmissionResponse", func() {
	Description("Response after test submission")

	Field(1, "submission_id", String, "Unique submission identifier")
	Field(2, "test_id", String, "Test identifier")
	Field(3, "user_id", Int64, "User identifier")
	Field(4, "score", Float64, "Test score (0-100)")
	Field(5, "total_points", Int, "Total points earned")
	Field(6, "max_points", Int, "Maximum possible points")
	Field(7, "correct_answers", Int, "Number of correct answers")
	Field(8, "total_questions", Int, "Total number of questions")
	Field(9, "time_taken", Int, "Time taken in seconds")
	Field(10, "passed", Boolean, "Whether test was passed")
	Field(11, "submitted_at", String, "Submission timestamp")
	Field(12, "graded_at", String, "Grading timestamp")

	Required("submission_id", "test_id", "user_id", "score", "total_points", "max_points",
		"correct_answers", "total_questions", "passed", "submitted_at", "graded_at")
})

var QuestionResult = Type("QuestionResult", func() {
	Description("Detailed result for a specific question")

	Field(1, "question_id", String, "Question identifier")
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

	Field(1, "submission_id", String, "Submission identifier")
	Field(2, "test_id", String, "Test identifier")
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
	Field(13, "submitted_at", String, "Submission timestamp")
	Field(14, "graded_at", String, "Grading timestamp")
	Field(15, "category", String, "Test category")

	Required("submission_id", "test_id", "test_title", "user_id", "score",
		"total_points", "max_points", "correct_answers", "total_questions",
		"passed", "question_results", "submitted_at", "graded_at")
})

// === PARTICIPANT AND USER TYPES ===
var Participant = Type("Participant", func() {
	Description("Test participant information")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "username", String, "Username")
	Field(3, "email", String, "User email")
	Field(4, "submission_id", String, "Submission identifier (if submitted)")
	Field(5, "score", Float64, "Test score (if completed)")
	Field(6, "status", String, "Status (not_started, in_progress, completed, expired)")
	Field(7, "started_at", String, "Test start timestamp")
	Field(8, "submitted_at", String, "Submission timestamp")
	Field(9, "time_taken", Int, "Time taken in seconds")
	Field(10, "passed", Boolean, "Whether test was passed")

	Required("user_id", "username", "status")
})

var ParticipantsResponse = Type("ParticipantsResponse", func() {
	Description("Response containing list of test participants")

	Field(1, "test_id", String, "Test identifier")
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

	Field(1, "submission_id", String, "Submission identifier")
	Field(2, "test_id", String, "Test identifier")
	Field(3, "test_title", String, "Test title")
	Field(4, "category", String, "Test category")
	Field(5, "score", Float64, "Test score")
	Field(6, "passed", Boolean, "Whether test was passed")
	Field(7, "time_taken", Int, "Time taken in seconds")
	Field(8, "submitted_at", String, "Submission timestamp")
	Field(9, "graded_at", String, "Grading timestamp")
	Field(10, "difficulty_level", String, "Test difficulty level")

	Required("submission_id", "test_id", "test_title", "category",
		"score", "passed", "submitted_at", "graded_at")
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

// === PROGRESS TRACKING TYPES ===
var CategoryProgress = Type("CategoryProgress", func() {
	Description("Progress in a specific category")

	Field(1, "category", String, "Category name")
	Field(2, "tests_taken", Int, "Number of tests taken")
	Field(3, "tests_passed", Int, "Number of tests passed")
	Field(4, "average_score", Float64, "Average score in category")
	Field(5, "best_score", Float64, "Best score in category")
	Field(7, "last_activity", String, "Last activity timestamp")
	Field(8, "improvement_rate", Float64, "Score improvement rate")

	Required("category", "tests_taken", "tests_passed", "average_score", "best_score")
})

var StudentProgressResponse = Type("StudentProgressResponse", func() {
	Description("Comprehensive student progress report")

	Field(1, "user_id", Int64, "User identifier")
	Field(2, "overall_score", Float64, "Overall average score")
	Field(3, "total_tests", Int, "Total number of tests taken")
	Field(4, "tests_passed", Int, "Number of tests passed")
	Field(5, "tests_failed", Int, "Number of tests failed")
	Field(7, "category_progress", ArrayOf(CategoryProgress), "Progress by category")
	Field(8, "recent_activity", ArrayOf(UserSubmission), "Recent test submissions")
	Field(9, "achievements", ArrayOf(String), "Achieved milestones")
	Field(10, "strengths", ArrayOf(String), "Strong categories")
	Field(11, "areas_for_improvement", ArrayOf(String), "Categories needing improvement")
	Field(12, "generated_at", String, "Report generation timestamp")

	Required("user_id", "overall_score", "total_tests", "tests_passed",
		"tests_failed", "category_progress", "generated_at")
})

// === ANALYTICS TYPES ===
var QuestionAnalytics = Type("QuestionAnalytics", func() {
	Description("Analytics for a specific question")

	Field(1, "question_id", String, "Question identifier")
	Field(2, "question_text", String, "Question text")
	Field(3, "total_attempts", Int, "Total number of attempts")
	Field(4, "correct_attempts", Int, "Number of correct attempts")
	Field(5, "success_rate", Float64, "Success rate percentage")
	Field(6, "average_time", Float64, "Average time spent on question")
	Field(7, "difficulty_index", Float64, "Calculated difficulty index")
	Field(8, "option_distribution", ArrayOf(Int), "Distribution of selected options")

	Required("question_id", "question_text", "total_attempts", "correct_attempts",
		"success_rate", "average_time", "difficulty_index")
})

var TestAnalyticsResponse = Type("TestAnalyticsResponse", func() {
	Description("Comprehensive test analytics")

	Field(1, "test_id", String, "Test identifier")
	Field(2, "test_title", String, "Test title")
	Field(3, "total_attempts", Int, "Total number of attempts")
	Field(4, "completed_attempts", Int, "Number of completed attempts")
	Field(5, "average_score", Float64, "Average score")
	Field(6, "median_score", Float64, "Median score")
	Field(7, "highest_score", Float64, "Highest score")
	Field(8, "lowest_score", Float64, "Lowest score")
	Field(9, "pass_rate", Float64, "Pass rate percentage")
	Field(10, "average_time", Float64, "Average completion time")
	Field(11, "score_distribution", ArrayOf(Int), "Score distribution by ranges")
	Field(12, "question_analytics", ArrayOf(QuestionAnalytics), "Per-question analytics")
	Field(13, "generated_at", String, "Analytics generation timestamp")

	Required("test_id", "test_title", "total_attempts", "completed_attempts",
		"average_score", "pass_rate", "question_analytics", "generated_at")
})

// === AVAILABLE TESTS TYPES ===
var AvailableTest = Type("AvailableTest", func() {
	Description("Available test information for students")

	Field(1, "test_id", String, "Test identifier")
	Field(2, "title", String, "Test title")
	Field(3, "description", String, "Test description")
	Field(4, "category", String, "Test category")
	Field(5, "difficulty_level", String, "Difficulty level")
	Field(6, "duration_minutes", Int, "Duration in minutes")
	Field(7, "total_questions", Int, "Total number of questions")
	Field(8, "passing_score", Float64, "Minimum passing score")
	Field(9, "attempts_allowed", Int, "Number of attempts allowed")
	Field(10, "user_attempts", Int, "Number of attempts by current user")
	Field(11, "best_score", Float64, "User's best score (if any)")
	Field(12, "last_attempt", String, "Last attempt timestamp")
	Field(13, "expires_at", String, "Test expiration timestamp")
	Field(14, "estimated_difficulty", String, "AI-estimated difficulty based on performance")

	Required("test_id", "title", "category", "difficulty_level", "duration_minutes",
		"total_questions", "passing_score", "attempts_allowed", "user_attempts")
})

var AvailableTestsResponse = Type("AvailableTestsResponse", func() {
	Description("Response containing available tests for a student")

	Field(1, "tests", ArrayOf(AvailableTest), "List of available tests")
	Field(2, "total_tests", Int, "Total number of available tests")
	Field(3, "recommended_tests", ArrayOf(String), "Recommended test IDs based on progress")
	Field(4, "page", Int, "Current page number")
	Field(5, "limit", Int, "Items per page")
	Field(6, "total_pages", Int, "Total number of pages")
	Field(7, "filters", MapOf(String, Any), "Applied filters")

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
