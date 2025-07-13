package controllers

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	profilesgen "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
	text "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/repositories/mocks"
)

func setupTestService(
	articleRepo *mocks.MockArticleRepository,
	sectionRepo *mocks.MockSectionRepository,
	courseRepo *mocks.MockCourseRepository,
	profilesServiceRepo *mocks.MockProfilesServiceRepository,
) *textsrvc {
	if articleRepo == nil {
		articleRepo = &mocks.MockArticleRepository{}
	}
	if sectionRepo == nil {
		sectionRepo = &mocks.MockSectionRepository{}
	}
	if courseRepo == nil {
		courseRepo = &mocks.MockCourseRepository{}
	}
	if profilesServiceRepo == nil {
		profilesServiceRepo = &mocks.MockProfilesServiceRepository{}
	}

	return &textsrvc{
		articleRepo:         articleRepo,
		sectionRepo:         sectionRepo,
		courseRepo:          courseRepo,
		profilesServiceRepo: profilesServiceRepo,
	}
}

// createTestCourse creates a test course for use in tests
func createTestCourse() textdb.Course {
	return textdb.Course{
		ID:          1,
		Title:       "Test Course",
		Description: "Test course description",
		ImageUrl:    pgtype.Text{String: "https://example.com/image.jpg", Valid: true},
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestSection creates a test section for use in tests
func createTestSection() textdb.Section {
	return textdb.Section{
		ID:          1,
		CourseID:    1,
		Title:       "Test Section",
		Description: "Test section description",
		Order:       1,
		CreatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestArticle creates a test article for use in tests
func createTestArticle() textdb.Article {
	return textdb.Article{
		ID:        1,
		SectionID: 1,
		Title:     "Test Article",
		Content:   "Test article content",
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestValidationResponse creates a test teacher validation response
func createTestTeacherValidationResponse() *profilesgen.CompleteProfileResponse {
	return &profilesgen.CompleteProfileResponse{
		UserID:    1,
		Role:      "teacher",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		IsActive:  true,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: &[]int64{time.Now().UnixMilli()}[0],
	}
}

// createTestStudentValidationResponse creates a test student validation response
func createTestStudentValidationResponse() *profilesgen.CompleteProfileResponse {
	return &profilesgen.CompleteProfileResponse{
		UserID:    2,
		Role:      "student",
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		IsActive:  true,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: &[]int64{time.Now().UnixMilli()}[0],
	}
}

// Helper function to get pointer to string
func stringPtr(s string) *string {
	return &s
}

// Test Course endpoints

func TestCreateCourse_Success_Teacher(t *testing.T) {
	// Arrange
	mockCourseRepo := &mocks.MockCourseRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, mockCourseRepo, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockCourseRepo.On("CreateCourse", mock.Anything, mock.AnythingOfType("textdb.CreateCourseParams")).Return(nil)

	// Act
	result, err := service.CreateCourse(context.Background(), &text.CreateCoursePayload{
		SessionToken: "valid-teacher-session",
		Title:        "Test Course",
		Description:  "Test course description",
		ImageURL:     stringPtr("https://example.com/image.jpg"),
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Course created successfully", result.Message)

	mockCourseRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestCreateCourse_Unauthorized_Student(t *testing.T) {
	// Arrange
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, nil, mockProfilesRepo)

	studentValidation := createTestStudentValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(studentValidation, nil)

	// Act
	result, err := service.CreateCourse(context.Background(), &text.CreateCoursePayload{
		SessionToken: "valid-student-session",
		Title:        "Test Course",
		Description:  "Test course description",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockProfilesRepo.AssertExpectations(t)
}

func TestGetCourse_Success(t *testing.T) {
	// Arrange
	mockCourseRepo := &mocks.MockCourseRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, mockCourseRepo, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()
	expectedCourse := createTestCourse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockCourseRepo.On("GetCourse", mock.Anything, int64(1)).Return(expectedCourse, nil)

	// Act
	result, err := service.GetCourse(context.Background(), &text.GetCoursePayload{
		SessionToken: "valid-session",
		CourseID:     1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "Test Course", result.Title)
	assert.Equal(t, "Test course description", result.Description)

	mockCourseRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestGetCourse_NotFound(t *testing.T) {
	// Arrange
	mockCourseRepo := &mocks.MockCourseRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, mockCourseRepo, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockCourseRepo.On("GetCourse", mock.Anything, int64(999)).Return(textdb.Course{}, assert.AnError)

	// Act
	result, err := service.GetCourse(context.Background(), &text.GetCoursePayload{
		SessionToken: "valid-session",
		CourseID:     999,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockCourseRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestListCourses_Success(t *testing.T) {
	// Arrange
	mockCourseRepo := &mocks.MockCourseRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, mockCourseRepo, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()
	expectedCourses := []textdb.Course{createTestCourse()}

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockCourseRepo.On("ListCourses", mock.Anything).Return(expectedCourses, nil)

	// Act
	result, err := service.ListCourses(context.Background(), &text.ListCoursesPayload{
		SessionToken: "valid-session",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "Test Course", result[0].Title)

	mockCourseRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

// Test Section endpoints

func TestCreateSection_Success_Teacher(t *testing.T) {
	// Arrange
	mockSectionRepo := &mocks.MockSectionRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, mockSectionRepo, nil, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockSectionRepo.On("GetNextOrderForCourse", mock.Anything, int64(1)).Return(int32(1), nil)
	mockSectionRepo.On("CreateSection", mock.Anything, mock.AnythingOfType("textdb.CreateSectionParams")).Return(nil)

	// Act
	result, err := service.CreateSection(context.Background(), &text.CreateSectionPayload{
		SessionToken: "valid-teacher-session",
		CourseID:     1,
		Title:        "Test Section",
		Description:  "Test section description",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Section created successfully", result.Message)

	mockSectionRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestGetSection_Success(t *testing.T) {
	// Arrange
	mockSectionRepo := &mocks.MockSectionRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, mockSectionRepo, nil, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()
	expectedSection := createTestSection()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockSectionRepo.On("GetSection", mock.Anything, int64(1)).Return(expectedSection, nil)

	// Act
	result, err := service.GetSection(context.Background(), &text.GetSectionPayload{
		SessionToken: "valid-session",
		SectionID:    1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "Test Section", result.Title)
	assert.Equal(t, int64(1), result.CourseID)

	mockSectionRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

// Test Article endpoints

func TestCreateArticle_Success_Teacher(t *testing.T) {
	// Arrange
	mockArticleRepo := &mocks.MockArticleRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockArticleRepo, nil, nil, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockArticleRepo.On("CreateArticle", mock.Anything, mock.AnythingOfType("textdb.CreateArticleParams")).Return(nil)

	// Act
	result, err := service.CreateArticle(context.Background(), &text.CreateArticlePayload{
		SessionToken: "valid-teacher-session",
		SectionID:    1,
		Title:        "Test Article",
		Content:      "Test article content",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Article created successfully", result.Message)

	mockArticleRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestGetArticle_Success(t *testing.T) {
	// Arrange
	mockArticleRepo := &mocks.MockArticleRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(mockArticleRepo, nil, nil, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()
	expectedArticle := createTestArticle()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockArticleRepo.On("GetArticle", mock.Anything, int64(1)).Return(expectedArticle, nil)

	// Act
	result, err := service.GetArticle(context.Background(), &text.GetArticlePayload{
		SessionToken: "valid-session",
		ArticleID:    1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "Test Article", result.Title)
	assert.Equal(t, "Test article content", result.Content)
	assert.Equal(t, int64(1), result.SectionID)

	mockArticleRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestDeleteCourse_Success_Teacher(t *testing.T) {
	// Arrange
	mockCourseRepo := &mocks.MockCourseRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, mockCourseRepo, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockCourseRepo.On("DeleteCourse", mock.Anything, int64(1)).Return(nil)

	// Act
	result, err := service.DeleteCourse(context.Background(), &text.DeleteCoursePayload{
		SessionToken: "valid-teacher-session",
		CourseID:     1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Course deleted successfully", result.Message)

	mockCourseRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestUpdateCourse_Success_Teacher(t *testing.T) {
	// Arrange
	mockCourseRepo := &mocks.MockCourseRepository{}
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, mockCourseRepo, mockProfilesRepo)

	teacherValidation := createTestTeacherValidationResponse()

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(teacherValidation, nil)
	mockCourseRepo.On("UpdateCourse", mock.Anything, mock.AnythingOfType("textdb.UpdateCourseParams")).Return(nil)

	// Act
	result, err := service.UpdateCourse(context.Background(), &text.UpdateCoursePayload{
		SessionToken: "valid-teacher-session",
		CourseID:     1,
		Title:        stringPtr("Updated Course"),
		Description:  stringPtr("Updated description"),
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Course updated successfully", result.Message)

	mockCourseRepo.AssertExpectations(t)
	mockProfilesRepo.AssertExpectations(t)
}

func TestValidateUser_InvalidSession(t *testing.T) {
	// Arrange
	mockProfilesRepo := &mocks.MockProfilesServiceRepository{}
	service := setupTestService(nil, nil, nil, mockProfilesRepo)

	mockProfilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(nil, assert.AnError)

	// Act
	result, err := service.GetCourse(context.Background(), &text.GetCoursePayload{
		SessionToken: "invalid-session",
		CourseID:     1,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockProfilesRepo.AssertExpectations(t)
}