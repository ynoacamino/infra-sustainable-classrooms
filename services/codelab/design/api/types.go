package design

import (
	. "goa.design/goa/v3/dsl"
)

var SimpleResponse = Type("SimpleResponse", func() {
	Description("Basic response with success status and message")

	Field(1, "success", Boolean, "Operation success status")
	Field(2, "message", String, "Response message")

	Required("success", "message")
})

// Exercise represents a coding exercise
var Exercise = Type("Exercise", func() {
	Description("A coding exercise with initial code, solution and difficulty level")

	Field(1, "id", Int64, "Exercise ID", func() {
		Example(1)
	})
	Field(2, "title", String, "Exercise title", func() {
		Example("Sum Two Numbers")
		MaxLength(200)
	})
	Field(3, "description", String, "Exercise description", func() {
		Example("Write a function that returns the sum of two numbers")
	})
	Field(4, "initial_code", String, "Initial code template", func() {
		Example("def sum_two_numbers(a, b):\n    # Write your code here\n    pass")
	})
	Field(5, "solution", String, "Exercise solution", func() {
		Example("def sum_two_numbers(a, b):\n    return a + b")
	})
	Field(6, "difficulty", String, "Exercise difficulty level", func() {
		Example("easy")
		Enum("easy", "medium", "hard")
	})
	Field(7, "created_by", Int64, "ID of user who created the exercise", func() {
		Example(123)
	})
	Field(8, "created_at", Int64, "Creation timestamp in miliseconds", func() {
		Example(1672531200000)
	})
	Field(9, "updated_at", Int64, "Last update timestamp", func() {
		Example(1672531200000)
	})

	Required("id", "title", "description", "initial_code", "solution", "difficulty", "created_by", "created_at", "updated_at")
})

// ExerciseForStudents represents an exercise without the solution (for students)
var ExerciseForStudents = Type("ExerciseForStudents", func() {
	Description("A coding exercise without solution for students")

	Field(1, "id", Int64, "Exercise ID", func() {
		Example(1)
	})
	Field(2, "title", String, "Exercise title", func() {
		Example("Sum Two Numbers")
		MaxLength(200)
	})
	Field(3, "description", String, "Exercise description", func() {
		Example("Write a function that returns the sum of two numbers")
	})
	Field(4, "initial_code", String, "Initial code template", func() {
		Example("def sum_two_numbers(a, b):\n    # Write your code here\n    pass")
	})
	Field(5, "difficulty", String, "Exercise difficulty level", func() {
		Example("easy")
		Enum("easy", "medium", "hard")
	})
	Field(6, "tests", ArrayOf(Test), "Associated tests for the exercise", func() {
		Description("List of public tests for the exercise")
	})
	Field(7, "attempts", ArrayOf(Attempt), "Associated attempts for the exercise", func() {
		Description("List of attempts made by students for this exercise")
	})
	Field(8, "answer", Answer, "Student's answer/participation in the exercise", func() {
		Description("Student's answer/participation in the exercise")
	})
	Field(9, "created_by", Int64, "ID of user who created the exercise", func() {
		Example(123)
	})
	Field(10, "created_at", Int64, "Creation timestamp", func() {
		Example(1672531200000)
	})
	Field(11, "updated_at", Int64, "Last update timestamp", func() {
		Example(1672531200000)
	})

	Required("id", "title", "description", "initial_code", "difficulty", "tests", "attempts", "answer", "created_by", "created_at", "updated_at")
})

var ExerciseForStudentsListView = Type("ExerciseForStudentsListView", func() {
	Description("View for listing exercises available to students")

	Field(1, "id", Int64, "Exercise ID", func() {
		Example(1)
	})
	Field(2, "title", String, "Exercise title", func() {
		Example("Sum Two Numbers")
		MaxLength(200)
	})
	Field(3, "description", String, "Exercise description", func() {
		Example("Write a function that returns the sum of two numbers")
	})
	Field(4, "difficulty", String, "Exercise difficulty level", func() {
		Example("easy")
		Enum("easy", "medium", "hard")
	})
	Field(5, "completed", Boolean, "Whether the exercise is completed by the student", func() {
		Example(false)
	})
	Field(6, "created_by", Int64, "ID of user who created the exercise", func() {
		Example(123)
	})
	Field(7, "created_at", Int64, "Creation timestamp", func() {
		Example(1672531200000)
	})
	Field(8, "updated_at", Int64, "Last update timestamp", func() {
		Example(1672531200000)
	})

	Required("id", "title", "description", "difficulty", "created_by", "created_at", "updated_at")
})

// Test represents a test case for an exercise
var Test = Type("Test", func() {
	Description("A test case with input and expected output")

	Field(1, "id", Int64, "Test ID", func() {
		Example(1)
	})
	Field(2, "input", String, "Test input", func() {
		Example("5, 3")
	})
	Field(3, "output", String, "Expected output", func() {
		Example("8")
	})
	Field(4, "public", Boolean, "Whether test is visible to students", func() {
		Example(true)
	})
	Field(5, "exercise_id", Int64, "Associated exercise ID", func() {
		Example(1)
	})
	Field(6, "created_at", Int64, "Creation timestamp", func() {
		Example(1672531200000)
	})
	Field(7, "updated_at", Int64, "Last update timestamp", func() {
		Example(1672531200000)
	})

	Required("id", "input", "output", "public", "exercise_id", "created_at", "updated_at")
})

// Answer represents a student's answer/participation in an exercise
var Answer = Type("Answer", func() {
	Description("A student's answer/participation in an exercise")

	Field(1, "id", Int64, "Answer ID", func() {
		Example(1)
	})
	Field(2, "exercise_id", Int64, "Associated exercise ID", func() {
		Example(1)
	})
	Field(3, "user_id", Int64, "Student user ID", func() {
		Example(123)
	})
	Field(4, "completed", Boolean, "Whether the exercise is completed", func() {
		Example(false)
	})
	Field(5, "created_at", Int64, "Creation timestamp", func() {
		Example(1672531200000)
	})
	Field(6, "updated_at", Int64, "Last update timestamp", func() {
		Example(1672531200000)
	})

	Required("id", "exercise_id", "user_id", "completed", "created_at", "updated_at")
})

// Attempt represents a code submission attempt
var Attempt = Type("Attempt", func() {
	Description("A code submission attempt for an answer")

	Field(1, "id", Int64, "Attempt ID", func() {
		Example(1)
	})
	Field(2, "answer_id", Int64, "Associated answer ID", func() {
		Example(1)
	})
	Field(3, "code", String, "Submitted code", func() {
		Example("def sum_two_numbers(a, b):\n    return a + b")
	})
	Field(4, "success", Boolean, "Whether the attempt was successful", func() {
		Example(true)
	})
	Field(5, "created_at", Int64, "Creation timestamp", func() {
		Example(1672531200000)
	})

	Required("id", "answer_id", "code", "success", "created_at")
})

// CreateExercisePayload for creating a new exercise
var CreateExercisePayload = Type("CreateExercisePayload", func() {
	Description("Payload for creating a new exercise")

	Field(1, "title", String, "Exercise title", func() {
		Example("Sum Two Numbers")
		MaxLength(200)
	})
	Field(2, "description", String, "Exercise description", func() {
		Example("Write a function that returns the sum of two numbers")
	})
	Field(3, "initial_code", String, "Initial code template", func() {
		Example("def sum_two_numbers(a, b):\n    # Write your code here\n    pass")
	})
	Field(4, "solution", String, "Exercise solution", func() {
		Example("def sum_two_numbers(a, b):\n    return a + b")
	})
	Field(5, "difficulty", String, "Exercise difficulty level", func() {
		Example("easy")
		Enum("easy", "medium", "hard")
	})
	Field(6, "created_by", Int64, "ID of user creating the exercise", func() {
		Example(123)
	})
	Field(7, "session_token", String, "Authentication session token")

	Required("session_token", "title", "description", "initial_code", "solution", "difficulty", "created_by")
})

// UpdateExercisePayload for updating an exercise
var UpdateExercisePayload = Type("UpdateExercisePayload", func() {
	Description("Payload for updating an exercise")

	Field(1, "title", String, "Exercise title", func() {
		Example("Sum Two Numbers")
		MaxLength(200)
	})
	Field(2, "description", String, "Exercise description", func() {
		Example("Write a function that returns the sum of two numbers")
	})
	Field(3, "initial_code", String, "Initial code template", func() {
		Example("def sum_two_numbers(a, b):\n    # Write your code here\n    pass")
	})
	Field(4, "solution", String, "Exercise solution", func() {
		Example("def sum_two_numbers(a, b):\n    return a + b")
	})
	Field(5, "difficulty", String, "Exercise difficulty level", func() {
		Example("easy")
		Enum("easy", "medium", "hard")
	})

	Required("title", "description", "initial_code", "solution", "difficulty")
})

// CreateTestPayload for creating a new test
var CreateTestPayload = Type("CreateTestPayload", func() {
	Description("Payload for creating a new test")

	Field(1, "input", String, "Test input", func() {
		Example("5, 3")
	})
	Field(2, "output", String, "Expected output", func() {
		Example("8")
	})
	Field(3, "public", Boolean, "Whether test is visible to students", func() {
		Example(true)
	})
	Field(4, "exercise_id", Int64, "Associated exercise ID", func() {
		Example(1)
	})
	Field(5, "session_token", String, "Authentication session token")

	Required("session_token", "input", "output", "public", "exercise_id")
})

// UpdateTestPayload for updating a test
var UpdateTestPayload = Type("UpdateTestPayload", func() {
	Description("Payload for updating a test")

	Field(1, "input", String, "Test input", func() {
		Example("5, 3")
	})
	Field(2, "output", String, "Expected output", func() {
		Example("8")
	})
	Field(3, "public", Boolean, "Whether test is visible to students", func() {
		Example(true)
	})

	Required("input", "output", "public")
})

// CreateAttemptPayload for creating a new attempt
var CreateAttemptPayload = Type("CreateAttemptPayload", func() {
	Description("Payload for creating a new attempt")

	Field(1, "exercise_id", Int64, "Associated exercise ID", func() {
		Example(1)
	})
	Field(2, "code", String, "Submitted code", func() {
		Example("def sum_two_numbers(a, b):\n    return a + b")
	})
	Field(3, "success", Boolean, "Whether the attempt was successful", func() {
		Example(true)
	})
	Field(4, "session_token", String, "Authentication session token")

	Required("session_token", "exercise_id", "code", "success")
})
