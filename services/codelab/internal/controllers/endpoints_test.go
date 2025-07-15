package controllers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/codelab"
	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/internal/repositories/mocks"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

// setupTestService creates a service instance with mocked repositories for testing
func setupTestService(
	profilesServiceRepo *mocks.MockProfilesServiceRepository,
	exercisesRepo *mocks.MockExercisesRepository,
	testsRepo *mocks.MockTestsRepository,
	answersRepo *mocks.MockAnswersRepository,
	attemptsRepo *mocks.MockAttemptsRepository,
) *codelabsvrc {
	// If mocks are nil, create empty ones
	if profilesServiceRepo == nil {
		profilesServiceRepo = &mocks.MockProfilesServiceRepository{}
	}
	if exercisesRepo == nil {
		exercisesRepo = &mocks.MockExercisesRepository{}
	}
	if testsRepo == nil {
		testsRepo = &mocks.MockTestsRepository{}
	}
	if answersRepo == nil {
		answersRepo = &mocks.MockAnswersRepository{}
	}
	if attemptsRepo == nil {
		attemptsRepo = &mocks.MockAttemptsRepository{}
	}

	return &codelabsvrc{
		profilesServiceRepo: profilesServiceRepo,
		exercisesRepo:       exercisesRepo,
		testsRepo:           testsRepo,
		answersRepo:         answersRepo,
		attemptsRepo:        attemptsRepo,
	}
}

// createTestTeacherProfile creates a test teacher profile response
func createTestTeacherProfile() *profiles.CompleteProfileResponse {
	return &profiles.CompleteProfileResponse{
		UserID:    1,
		Role:      "teacher",
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@example.com",
		Phone:     stringPtr("+1234567890"),
		AvatarURL: stringPtr("https://example.com/avatar.jpg"),
		Bio:       stringPtr("Test teacher bio"),
		IsActive:  true,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: timePtr(time.Now().UnixMilli()),
		Position:  stringPtr("Professor"),
	}
}

// createTestStudentProfile creates a test student profile response
func createTestStudentProfile() *profiles.CompleteProfileResponse {
	return &profiles.CompleteProfileResponse{
		UserID:     1,
		Role:       "student",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		Phone:      stringPtr("+1234567890"),
		AvatarURL:  stringPtr("https://example.com/avatar.jpg"),
		Bio:        stringPtr("Test student bio"),
		IsActive:   true,
		CreatedAt:  time.Now().UnixMilli(),
		UpdatedAt:  timePtr(time.Now().UnixMilli()),
		GradeLevel: stringPtr("undergraduate"),
		Major:      stringPtr("Computer Science"),
	}
}

// createTestExercise creates a test exercise
func createTestExercise() codelabdb.Exercise {
	now := time.Now()
	return codelabdb.Exercise{
		ID:          1,
		Title:       "Test Exercise",
		Description: "This is a test exercise",
		InitialCode: "function solution(input) {\n  // Your code here\n}",
		Solution:    "function solution(input) {\n  return input * 2;\n}",
		Difficulty:  "easy",
		CreatedBy:   1,
		CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
	}
}

// createTestExerciseToResolve creates a test exercise for students (without solution)
func createTestExerciseToResolve() codelabdb.GetExerciseToResolveByIdRow {
	now := time.Now()
	return codelabdb.GetExerciseToResolveByIdRow{
		ID:          1,
		Title:       "Test Exercise",
		Description: "This is a test exercise",
		InitialCode: "function solution(input) {\n  // Your code here\n}",
		Difficulty:  "easy",
		CreatedBy:   1,
		CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
	}
}

// createTestAnswer creates a test answer
func createTestAnswer() codelabdb.Answer {
	now := time.Now()
	return codelabdb.Answer{
		ID:         1,
		ExerciseID: 1,
		UserID:     1,
		Completed:  false,
		CreatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
	}
}

// createTestAttempt creates a test attempt
func createTestAttempt(success bool) codelabdb.Attempt {
	now := time.Now()
	return codelabdb.Attempt{
		ID:        1,
		AnswerID:  1,
		Code:      "function solution(input) { return input * 2; }",
		Success:   success,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

// Helper function to create time pointers
func timePtr(t int64) *int64 {
	return &t
}

// ========================================
// EXERCISE TESTS
// ========================================

func TestCreateExercise_Success(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("CreateExercise", mock.Anything, mock.AnythingOfType("codelabdb.CreateExerciseParams")).Return(nil)

	// Act
	result, err := service.CreateExercise(context.Background(), &codelab.CreateExercisePayload{
		SessionToken: "valid-session-token",
		Title:        "Test Exercise",
		Description:  "This is a test exercise",
		InitialCode:  "function solution(input) {\n  // Your code here\n}",
		Solution:     "function solution(input) {\n  return input * 2;\n}",
		Difficulty:   "easy",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Exercise created successfully", result.Message)
	assert.True(t, result.Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

func TestCreateExercise_Unauthorized(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(nil, errors.New("unauthorized"))

	// Act
	result, err := service.CreateExercise(context.Background(), &codelab.CreateExercisePayload{
		SessionToken: "invalid-session-token",
		Title:        "Test Exercise",
		Description:  "This is a test exercise",
		InitialCode:  "function solution(input) {\n  // Your code here\n}",
		Solution:     "function solution(input) {\n  return input * 2;\n}",
		Difficulty:   "easy",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Unauthorized access")
	mockProfilesServiceRepo.AssertExpectations(t)
}

func TestCreateExercise_PermissionDenied_Student(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	expectedProfile := createTestStudentProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)

	// Act
	result, err := service.CreateExercise(context.Background(), &codelab.CreateExercisePayload{
		SessionToken: "valid-session-token",
		Title:        "Test Exercise",
		Description:  "This is a test exercise",
		InitialCode:  "function solution(input) {\n  // Your code here\n}",
		Solution:     "function solution(input) {\n  return input * 2;\n}",
		Difficulty:   "easy",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Only teachers can create exercises")
	mockProfilesServiceRepo.AssertExpectations(t)
}

func TestCreateExercise_DatabaseError(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("CreateExercise", mock.Anything, mock.AnythingOfType("codelabdb.CreateExerciseParams")).Return(errors.New("database error"))

	// Act
	result, err := service.CreateExercise(context.Background(), &codelab.CreateExercisePayload{
		SessionToken: "valid-session-token",
		Title:        "Test Exercise",
		Description:  "This is a test exercise",
		InitialCode:  "function solution(input) {\n  // Your code here\n}",
		Solution:     "function solution(input) {\n  return input * 2;\n}",
		Difficulty:   "easy",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Failed to create exercise")
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

func TestGetExercise_Success(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()
	expectedExercise := createTestExercise()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseById", mock.Anything, int64(1)).Return(expectedExercise, nil)

	// Act
	result, err := service.GetExercise(context.Background(), &codelab.GetExercisePayload{
		SessionToken: "valid-session-token",
		ID:           1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "Test Exercise", result.Title)
	assert.Equal(t, "This is a test exercise", result.Description)
	assert.Equal(t, "function solution(input) {\n  // Your code here\n}", result.InitialCode)
	assert.Equal(t, "function solution(input) {\n  return input * 2;\n}", result.Solution)
	assert.Equal(t, "easy", result.Difficulty)
	assert.Equal(t, int64(1), result.CreatedBy)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

func TestGetExercise_PermissionDenied_Student(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	expectedProfile := createTestStudentProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)

	// Act
	result, err := service.GetExercise(context.Background(), &codelab.GetExercisePayload{
		SessionToken: "valid-session-token",
		ID:           1,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Only teachers can view exercise with solution")
	mockProfilesServiceRepo.AssertExpectations(t)
}

func TestGetExercise_NotFound(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseById", mock.Anything, int64(999)).Return(codelabdb.Exercise{}, errors.New("not found"))

	// Act
	result, err := service.GetExercise(context.Background(), &codelab.GetExercisePayload{
		SessionToken: "valid-session-token",
		ID:           999,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Exercise not found")
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

// ========================================
// CREATE ATTEMPT TESTS (MOST CRITICAL)
// ========================================

func TestCreateAttempt_Success_AllTestsPass(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockAnswersRepo := &mocks.MockAnswersRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	mockAttemptsRepo := &mocks.MockAttemptsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, mockAnswersRepo, mockAttemptsRepo)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()
	expectedAnswer := createTestAnswer()
	expectedHiddenTests := []codelabdb.Test{
		{
			ID:         1,
			Input:      "5",
			Output:     "10",
			Public:     false,
			ExerciseID: 1,
		},
		{
			ID:         2,
			Input:      "3",
			Output:     "6",
			Public:     false,
			ExerciseID: 1,
		},
	}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockAnswersRepo.On("CheckIfAnswerExists", mock.Anything, mock.AnythingOfType("codelabdb.CheckIfAnswerExistsParams")).Return(int32(1), nil)
	mockAnswersRepo.On("GetAnswerByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAnswerByUserAndExerciseParams")).Return(expectedAnswer, nil)
	mockTestsRepo.On("GetHiddenTestsByExercise", mock.Anything, int64(1)).Return(expectedHiddenTests, nil)
	mockAttemptsRepo.On("CreateAttempt", mock.Anything, mock.AnythingOfType("codelabdb.CreateAttemptParams")).Return(nil)
	mockAnswersRepo.On("UpdateAnswerCompleted", mock.Anything, mock.AnythingOfType("codelabdb.UpdateAnswerCompletedParams")).Return(nil)

	// Valid JavaScript code that should pass the tests
	validCode := `function solution(input) { return parseInt(input) * 2; }`

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         validCode,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "All tests passed successfully", result.Message)
	assert.True(t, result.Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockAnswersRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
	mockAttemptsRepo.AssertExpectations(t)
}

func TestCreateAttempt_CodeExecutionError(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockAnswersRepo := &mocks.MockAnswersRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, mockAnswersRepo, nil)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()
	expectedAnswer := createTestAnswer()
	expectedHiddenTests := []codelabdb.Test{}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockAnswersRepo.On("CheckIfAnswerExists", mock.Anything, mock.AnythingOfType("codelabdb.CheckIfAnswerExistsParams")).Return(int32(1), nil)
	mockAnswersRepo.On("GetAnswerByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAnswerByUserAndExerciseParams")).Return(expectedAnswer, nil)
	mockTestsRepo.On("GetHiddenTestsByExercise", mock.Anything, int64(1)).Return(expectedHiddenTests, nil)

	// Invalid JavaScript code
	invalidCode := `function solution(input) { return input * 2; // missing closing brace`

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         invalidCode,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Code execution failed")
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockAnswersRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
}

func TestCreateAttempt_MissingFunction(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockAnswersRepo := &mocks.MockAnswersRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, mockAnswersRepo, nil)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()
	expectedAnswer := createTestAnswer()
	expectedHiddenTests := []codelabdb.Test{}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockAnswersRepo.On("CheckIfAnswerExists", mock.Anything, mock.AnythingOfType("codelabdb.CheckIfAnswerExistsParams")).Return(int32(1), nil)
	mockAnswersRepo.On("GetAnswerByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAnswerByUserAndExerciseParams")).Return(expectedAnswer, nil)
	mockTestsRepo.On("GetHiddenTestsByExercise", mock.Anything, int64(1)).Return(expectedHiddenTests, nil)

	// Code without required 'solution' function
	codeWithoutFunction := `var x = 42;`

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         codeWithoutFunction,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Code must define a 'solution' function")
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockAnswersRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
}

func TestCreateAttempt_TestFailure_RuntimeError(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockAnswersRepo := &mocks.MockAnswersRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	mockAttemptsRepo := &mocks.MockAttemptsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, mockAnswersRepo, mockAttemptsRepo)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()
	expectedAnswer := createTestAnswer()
	expectedHiddenTests := []codelabdb.Test{
		{
			ID:         1,
			Input:      "5",
			Output:     "10",
			Public:     false,
			ExerciseID: 1,
		},
	}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockAnswersRepo.On("CheckIfAnswerExists", mock.Anything, mock.AnythingOfType("codelabdb.CheckIfAnswerExistsParams")).Return(int32(1), nil)
	mockAnswersRepo.On("GetAnswerByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAnswerByUserAndExerciseParams")).Return(expectedAnswer, nil)
	mockTestsRepo.On("GetHiddenTestsByExercise", mock.Anything, int64(1)).Return(expectedHiddenTests, nil)
	mockAttemptsRepo.On("CreateAttempt", mock.Anything, mock.AnythingOfType("codelabdb.CreateAttemptParams")).Return(nil)

	// Code that causes runtime error
	errorCode := `function solution(input) { throw new Error("Runtime error"); }`

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         errorCode,
	})

	// Assert
	assert.NoError(t, err) // No error, but test failure
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Test failed:")
	assert.False(t, result.Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockAnswersRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
	mockAttemptsRepo.AssertExpectations(t)
}

func TestCreateAttempt_TestFailure_WrongOutput(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockAnswersRepo := &mocks.MockAnswersRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	mockAttemptsRepo := &mocks.MockAttemptsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, mockAnswersRepo, mockAttemptsRepo)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()
	expectedAnswer := createTestAnswer()
	expectedHiddenTests := []codelabdb.Test{
		{
			ID:         1,
			Input:      "5",
			Output:     "10",
			Public:     false,
			ExerciseID: 1,
		},
	}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockAnswersRepo.On("CheckIfAnswerExists", mock.Anything, mock.AnythingOfType("codelabdb.CheckIfAnswerExistsParams")).Return(int32(1), nil)
	mockAnswersRepo.On("GetAnswerByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAnswerByUserAndExerciseParams")).Return(expectedAnswer, nil)
	mockTestsRepo.On("GetHiddenTestsByExercise", mock.Anything, int64(1)).Return(expectedHiddenTests, nil)
	mockAttemptsRepo.On("CreateAttempt", mock.Anything, mock.AnythingOfType("codelabdb.CreateAttemptParams")).Return(nil)

	// Code that returns wrong output
	wrongCode := `function solution(input) { return parseInt(input) * 3; }`

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         wrongCode,
	})

	// Assert
	assert.NoError(t, err) // No error, but test failure
	assert.NotNil(t, result)
	assert.Contains(t, result.Message, "Test failed: expected")
	assert.False(t, result.Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockAnswersRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
	mockAttemptsRepo.AssertExpectations(t)
}

func TestCreateAttempt_PermissionDenied_Teacher(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         "function solution(input) { return input * 2; }",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Only students can create attempts")
	mockProfilesServiceRepo.AssertExpectations(t)
}

func TestCreateAttempt_ExerciseNotFound(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestStudentProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(999)).Return(codelabdb.GetExerciseToResolveByIdRow{}, errors.New("not found"))

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   999,
		Code:         "function solution(input) { return input * 2; }",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Exercise not found")
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

func TestCreateAttempt_CreateAnswerWhenNotExists(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockAnswersRepo := &mocks.MockAnswersRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	mockAttemptsRepo := &mocks.MockAttemptsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, mockAnswersRepo, mockAttemptsRepo)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()
	expectedAnswer := createTestAnswer()
	expectedHiddenTests := []codelabdb.Test{
		{
			ID:         1,
			Input:      "5",
			Output:     "10",
			Public:     false,
			ExerciseID: 1,
		},
	}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockAnswersRepo.On("CheckIfAnswerExists", mock.Anything, mock.AnythingOfType("codelabdb.CheckIfAnswerExistsParams")).Return(int32(0), errors.New("not found"))
	mockAnswersRepo.On("CreateAnswer", mock.Anything, mock.AnythingOfType("codelabdb.CreateAnswerParams")).Return(nil)
	mockAnswersRepo.On("GetAnswerByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAnswerByUserAndExerciseParams")).Return(expectedAnswer, nil)
	mockTestsRepo.On("GetHiddenTestsByExercise", mock.Anything, int64(1)).Return(expectedHiddenTests, nil)
	mockAttemptsRepo.On("CreateAttempt", mock.Anything, mock.AnythingOfType("codelabdb.CreateAttemptParams")).Return(nil)
	mockAnswersRepo.On("UpdateAnswerCompleted", mock.Anything, mock.AnythingOfType("codelabdb.UpdateAnswerCompletedParams")).Return(nil)

	// Valid JavaScript code that should pass the tests
	validCode := `function solution(input) { return parseInt(input) * 2; }`

	// Act
	result, err := service.CreateAttempt(context.Background(), &codelab.CreateAttemptPayload{
		SessionToken: "valid-session-token",
		ExerciseID:   1,
		Code:         validCode,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "All tests passed successfully", result.Message)
	assert.True(t, result.Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockAnswersRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
	mockAttemptsRepo.AssertExpectations(t)
}

// ========================================
// TEST TESTS
// ========================================

func TestCreateTest_Success(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	mockTestsRepo := &mocks.MockTestsRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, mockTestsRepo, nil, nil)

	expectedProfile := createTestTeacherProfile()
	expectedExercise := createTestExercise()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseById", mock.Anything, int64(1)).Return(expectedExercise, nil)
	mockTestsRepo.On("CreateTest", mock.Anything, mock.AnythingOfType("codelabdb.CreateTestParams")).Return(nil)

	// Act
	result, err := service.CreateTest(context.Background(), &codelab.CreateTestPayload{
		SessionToken: "valid-session-token",
		Input:        "5",
		Output:       "10",
		Public:       true,
		ExerciseID:   1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test created successfully", result.Message)
	assert.True(t, result.Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
	mockTestsRepo.AssertExpectations(t)
}

func TestCreateTest_PermissionDenied_Student(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	expectedProfile := createTestStudentProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)

	// Act
	result, err := service.CreateTest(context.Background(), &codelab.CreateTestPayload{
		SessionToken: "valid-session-token",
		Input:        "5",
		Output:       "10",
		Public:       true,
		ExerciseID:   1,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Only teachers can create tests")
	mockProfilesServiceRepo.AssertExpectations(t)
}

func TestCreateTest_ExerciseNotFound(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseById", mock.Anything, int64(999)).Return(codelabdb.Exercise{}, errors.New("not found"))

	// Act
	result, err := service.CreateTest(context.Background(), &codelab.CreateTestPayload{
		SessionToken: "valid-session-token",
		Input:        "5",
		Output:       "10",
		Public:       true,
		ExerciseID:   999,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Exercise not found")
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

// ========================================
// STUDENT ACCESS TESTS
// ========================================

func TestGetExerciseForStudent_Success(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockExercisesRepo := &mocks.MockExercisesRepository{}
	service := setupTestService(mockProfilesServiceRepo, mockExercisesRepo, nil, nil, nil)

	expectedProfile := createTestStudentProfile()
	expectedExercise := createTestExerciseToResolve()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockExercisesRepo.On("GetExerciseToResolveById", mock.Anything, int64(1)).Return(expectedExercise, nil)

	// Act
	result, err := service.GetExerciseForStudent(context.Background(), &codelab.GetExerciseForStudentPayload{
		SessionToken: "valid-session-token",
		ID:           1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "Test Exercise", result.Title)
	assert.Equal(t, "This is a test exercise", result.Description)
	assert.Equal(t, "function solution(input) {\n  // Your code here\n}", result.InitialCode)
	assert.Equal(t, "easy", result.Difficulty)
	assert.Equal(t, int64(1), result.CreatedBy)
	// Note: No solution field in student response
	mockProfilesServiceRepo.AssertExpectations(t)
	mockExercisesRepo.AssertExpectations(t)
}

func TestGetAttemptsByUserAndExercise_Success(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	mockAttemptsRepo := &mocks.MockAttemptsRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, mockAttemptsRepo)

	expectedProfile := createTestStudentProfile()
	expectedAttempts := []codelabdb.Attempt{
		createTestAttempt(false),
		createTestAttempt(true),
	}

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	mockAttemptsRepo.On("GetAttemptsByUserAndExercise", mock.Anything, mock.AnythingOfType("codelabdb.GetAttemptsByUserAndExerciseParams")).Return(expectedAttempts, nil)

	// Act
	result, err := service.GetAttemptsByUserAndExercise(context.Background(), &codelab.GetAttemptsByUserAndExercisePayload{
		SessionToken: "valid-session-token",
		UserID:       1, // Same as profile UserID
		ExerciseID:   1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, int64(1), result[0].AnswerID)
	assert.Equal(t, "function solution(input) { return input * 2; }", result[0].Code)
	assert.False(t, result[0].Success)
	assert.True(t, result[1].Success)
	mockProfilesServiceRepo.AssertExpectations(t)
	mockAttemptsRepo.AssertExpectations(t)
}

func TestGetAttemptsByUserAndExercise_PermissionDenied_DifferentUser(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	expectedProfile := createTestStudentProfile() // UserID = 1

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)

	// Act
	result, err := service.GetAttemptsByUserAndExercise(context.Background(), &codelab.GetAttemptsByUserAndExercisePayload{
		SessionToken: "valid-session-token",
		UserID:       2, // Different from profile UserID
		ExerciseID:   1,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Students can only view their own attempts")
	mockProfilesServiceRepo.AssertExpectations(t)
}

func TestGetAttemptsByUserAndExercise_PermissionDenied_Teacher(t *testing.T) {
	// Arrange
	mockProfilesServiceRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockProfilesServiceRepo, nil, nil, nil, nil)

	expectedProfile := createTestTeacherProfile()

	mockProfilesServiceRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)

	// Act
	result, err := service.GetAttemptsByUserAndExercise(context.Background(), &codelab.GetAttemptsByUserAndExercisePayload{
		SessionToken: "valid-session-token",
		UserID:       1,
		ExerciseID:   1,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Only students can access this endpoint")
	mockProfilesServiceRepo.AssertExpectations(t)
}
