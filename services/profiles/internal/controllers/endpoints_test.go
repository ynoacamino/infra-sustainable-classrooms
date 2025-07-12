package controllers

import (
	"context"
	"testing"
	"time"

	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/repositories/mocks"
)

// setupTestService creates a service instance with mocked repositories for testing
func setupTestService(
	profileRepo *mocks.MockProfileRepository,
	studentProfileRepo *mocks.MockStudentProfileRepository,
	teacherProfileRepo *mocks.MockTeacherProfileRepository,
	authServiceRepo *mocks.MockAuthServiceRepository,
) *profilessrvc {
	// If mocks are nil, create empty ones
	if profileRepo == nil {
		profileRepo = &mocks.MockProfileRepository{}
	}
	if studentProfileRepo == nil {
		studentProfileRepo = &mocks.MockStudentProfileRepository{}
	}
	if teacherProfileRepo == nil {
		teacherProfileRepo = &mocks.MockTeacherProfileRepository{}
	}
	if authServiceRepo == nil {
		authServiceRepo = &mocks.MockAuthServiceRepository{}
	}

	return &profilessrvc{
		profileRepo:        profileRepo,
		studentProfileRepo: studentProfileRepo,
		teacherProfileRepo: teacherProfileRepo,
		authServiceRepo:    authServiceRepo,
	}
}

// createTestAuthUser creates a test auth user for use in tests
func createTestAuthUser() *auth.User {
	return &auth.User{
		ID:         1,
		Identifier: "test@example.com",
		CreatedAt:  time.Now().UnixMilli(),
		IsVerified: true,
		Metadata:   map[string]string{},
	}
}

// createTestValidationResponse creates a test validation response
func createTestValidationResponse() *auth.UserValidationResponse {
	return &auth.UserValidationResponse{
		Valid: true,
		User:  createTestAuthUser(),
		Session: &auth.Session{
			ID:        1,
			UserID:    1,
			CreatedAt: time.Now().UnixMilli(),
			ExpiresAt: time.Now().Add(24 * time.Hour).UnixMilli(),
			IsActive:  true,
			UserAgent: stringPtr("test-agent"),
			IPAddress: stringPtr("127.0.0.1"),
			DeviceID:  stringPtr("test-device"),
			Platform:  stringPtr("web"),
		},
	}
}

// createTestProfile creates a test profile
func createTestProfile(role string) profilesdb.Profile {
	return profilesdb.Profile{
		ID:        1,
		UserID:    1,
		Role:      role,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestStudentProfile creates a test student profile
func createTestStudentProfile() profilesdb.StudentProfile {
	return profilesdb.StudentProfile{
		ID:         1,
		ProfileID:  1,
		GradeLevel: "undergraduate",
		Major:      pgtype.Text{String: "Computer Science", Valid: true},
		CreatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestTeacherProfile creates a test teacher profile
func createTestTeacherProfile() profilesdb.TeacherProfile {
	return profilesdb.TeacherProfile{
		ID:        1,
		ProfileID: 1,
		Position:  "Professor",
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestCompleteStudentProfile creates a test complete student profile
func createTestCompleteStudentProfile() profilesdb.GetCompleteStudentProfileRow {
	return profilesdb.GetCompleteStudentProfileRow{
		UserID:     1,
		Role:       "student",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		Phone:      pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl:  pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:        pgtype.Text{String: "Test bio", Valid: true},
		IsActive:   true,
		CreatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		GradeLevel: "undergraduate",
		Major:      pgtype.Text{String: "Computer Science", Valid: true},
	}
}

// createTestCompleteTeacherProfile creates a test complete teacher profile
func createTestCompleteTeacherProfile() profilesdb.GetCompleteTeacherProfileRow {
	return profilesdb.GetCompleteTeacherProfileRow{
		UserID:    1,
		Role:      "teacher",
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		Phone:     pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Position:  "Professor",
	}
}

// createTestPublicProfile creates a test public profile
func createTestPublicProfile() profilesdb.GetPublicProfileByUserIdRow {
	return profilesdb.GetPublicProfileByUserIdRow{
		UserID:    1,
		Role:      "student",
		FirstName: "John",
		LastName:  "Doe",
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

func TestCreateStudentProfile_Success(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockStudentProfileRepo := &mocks.MockStudentProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, mockStudentProfileRepo, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("student")
	expectedStudentProfile := createTestStudentProfile()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, nil)
	mockProfileRepo.On("CreateProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateProfileParams")).Return(expectedProfile, nil)
	mockStudentProfileRepo.On("CreateStudentProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateStudentProfileParams")).Return(expectedStudentProfile, nil)

	// Act
	result, err := service.CreateStudentProfile(context.Background(), &profiles.CreateStudentProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		Phone:        stringPtr("+1234567890"),
		AvatarURL:    stringPtr("https://example.com/avatar.jpg"),
		Bio:          stringPtr("Test bio"),
		GradeLevel:   "undergraduate",
		Major:        stringPtr("Computer Science"),
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "john.doe@example.com", result.Email)
	assert.Equal(t, "undergraduate", result.GradeLevel)
	assert.Equal(t, "Computer Science", *result.Major)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockStudentProfileRepo.AssertExpectations(t)
}

func TestCreateStudentProfile_InvalidSession(t *testing.T) {
	// Arrange
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(nil, nil, nil, mockAuthServiceRepo)

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(nil, errors.New("invalid session"))

	// Act
	result, err := service.CreateStudentProfile(context.Background(), &profiles.CreateStudentProfilePayload{
		SessionToken: "invalid-session-token",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		GradeLevel:   "undergraduate",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
}

func TestCreateStudentProfile_ProfileAlreadyExists(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(true, nil)

	// Act
	result, err := service.CreateStudentProfile(context.Background(), &profiles.CreateStudentProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		GradeLevel:   "undergraduate",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateTeacherProfile_Success(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockTeacherProfileRepo := &mocks.MockTeacherProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, mockTeacherProfileRepo, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("teacher")
	expectedProfile.FirstName = "Jane"
	expectedProfile.LastName = "Smith"
	expectedProfile.Email = "jane.smith@example.com"
	expectedTeacherProfile := createTestTeacherProfile()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, nil)
	mockProfileRepo.On("CreateProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateProfileParams")).Return(expectedProfile, nil)
	mockTeacherProfileRepo.On("CreateTeacherProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateTeacherProfileParams")).Return(expectedTeacherProfile, nil)

	// Act
	result, err := service.CreateTeacherProfile(context.Background(), &profiles.CreateTeacherProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane.smith@example.com",
		Phone:        stringPtr("+1234567890"),
		AvatarURL:    stringPtr("https://example.com/avatar.jpg"),
		Bio:          stringPtr("Test bio"),
		Position:     "Professor",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Smith", result.LastName)
	assert.Equal(t, "jane.smith@example.com", result.Email)
	assert.Equal(t, "Professor", result.Position)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockTeacherProfileRepo.AssertExpectations(t)
}

func TestCreateTeacherProfile_InvalidSession(t *testing.T) {
	// Arrange
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(nil, nil, nil, mockAuthServiceRepo)

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(nil, errors.New("invalid session"))

	// Act
	result, err := service.CreateTeacherProfile(context.Background(), &profiles.CreateTeacherProfilePayload{
		SessionToken: "invalid-session-token",
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane.smith@example.com",
		Position:     "Professor",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_StudentSuccess(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockStudentProfileRepo := &mocks.MockStudentProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, mockStudentProfileRepo, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("student")
	expectedCompleteProfile := createTestCompleteStudentProfile()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(expectedProfile, nil)
	mockStudentProfileRepo.On("GetCompleteStudentProfile", mock.Anything, int64(1)).Return(expectedCompleteProfile, nil)

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid-session-token",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "student", result.Role)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "undergraduate", *result.GradeLevel)
	assert.Equal(t, "Computer Science", *result.Major)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockStudentProfileRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_TeacherSuccess(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockTeacherProfileRepo := &mocks.MockTeacherProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, mockTeacherProfileRepo, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("teacher")
	expectedCompleteProfile := createTestCompleteTeacherProfile()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(expectedProfile, nil)
	mockTeacherProfileRepo.On("GetCompleteTeacherProfile", mock.Anything, int64(1)).Return(expectedCompleteProfile, nil)

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid-session-token",
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "teacher", result.Role)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Smith", result.LastName)
	assert.Equal(t, "Professor", *result.Position)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockTeacherProfileRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_InvalidSession(t *testing.T) {
	// Arrange
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(nil, nil, nil, mockAuthServiceRepo)

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(nil, errors.New("invalid session"))

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "invalid-session-token",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_ProfileNotFound(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(profilesdb.Profile{}, errors.New("profile not found"))

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid-session-token",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestGetPublicProfileByID_Success(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, nil)

	expectedPublicProfile := createTestPublicProfile()

	mockProfileRepo.On("GetPublicProfileByUserId", mock.Anything, int64(1)).Return(expectedPublicProfile, nil)

	// Act
	result, err := service.GetPublicProfileByID(context.Background(), &profiles.GetPublicProfileByIDPayload{
		UserID: 1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "student", result.Role)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.True(t, result.IsActive)

	mockProfileRepo.AssertExpectations(t)
}

func TestGetPublicProfileByID_ProfileNotFound(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, nil)

	mockProfileRepo.On("GetPublicProfileByUserId", mock.Anything, int64(999)).Return(profilesdb.GetPublicProfileByUserIdRow{}, errors.New("profile not found"))

	// Act
	result, err := service.GetPublicProfileByID(context.Background(), &profiles.GetPublicProfileByIDPayload{
		UserID: 999,
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockProfileRepo.AssertExpectations(t)
}

func TestUpdateProfile_Success(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	existingProfile := createTestProfile("student")
	updatedProfile := createTestProfile("student")
	updatedProfile.FirstName = "Jane"
	updatedProfile.LastName = "Smith"

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(existingProfile, nil)
	mockProfileRepo.On("UpdateProfile", mock.Anything, mock.AnythingOfType("profilesdb.UpdateProfileParams")).Return(updatedProfile, nil)

	// Act
	result, err := service.UpdateProfile(context.Background(), &profiles.UpdateProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    stringPtr("Jane"),
		LastName:     stringPtr("Smith"),
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Smith", result.LastName)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestUpdateProfile_InvalidSession(t *testing.T) {
	// Arrange
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(nil, nil, nil, mockAuthServiceRepo)

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(nil, errors.New("invalid session"))

	// Act
	result, err := service.UpdateProfile(context.Background(), &profiles.UpdateProfilePayload{
		SessionToken: "invalid-session-token",
		FirstName:    stringPtr("Jane"),
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
}

func TestUpdateProfile_ProfileNotFound(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(profilesdb.Profile{}, errors.New("profile not found"))

	// Act
	result, err := service.UpdateProfile(context.Background(), &profiles.UpdateProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    stringPtr("Jane"),
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestValidateUserRole_Success(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, nil)

	expectedProfile := createTestProfile("student")

	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(expectedProfile, nil)

	// Act
	result, err := service.ValidateUserRole(context.Background(), &profiles.ValidateUserRolePayload{
		UserID: 1,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.UserID)
	assert.Equal(t, "student", result.Role)

	mockProfileRepo.AssertExpectations(t)
}

func TestValidateUserRole_ProfileNotFound(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, nil)

	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(999)).Return(profilesdb.Profile{}, errors.New("profile not found"))

	// Act
	result, err := service.ValidateUserRole(context.Background(), &profiles.ValidateUserRolePayload{
		UserID: 999,
	})

	// Assert
	assert.NoError(t, err) // ValidateUserRole doesn't return error for not found
	assert.NotNil(t, result)
	assert.Equal(t, int64(999), result.UserID)
	assert.Equal(t, "", result.Role) // Empty role when profile not found

	mockProfileRepo.AssertExpectations(t)
}

func TestCreateStudentProfile_CreateProfileFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, nil)
	mockProfileRepo.On("CreateProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateProfileParams")).Return(profilesdb.Profile{}, errors.New("database error"))

	// Act
	result, err := service.CreateStudentProfile(context.Background(), &profiles.CreateStudentProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		GradeLevel:   "undergraduate",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateStudentProfile_CreateStudentProfileFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockStudentProfileRepo := &mocks.MockStudentProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, mockStudentProfileRepo, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("student")

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, nil)
	mockProfileRepo.On("CreateProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateProfileParams")).Return(expectedProfile, nil)
	mockStudentProfileRepo.On("CreateStudentProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateStudentProfileParams")).Return(profilesdb.StudentProfile{}, errors.New("database error"))

	// Act
	result, err := service.CreateStudentProfile(context.Background(), &profiles.CreateStudentProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		GradeLevel:   "undergraduate",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockStudentProfileRepo.AssertExpectations(t)
}

func TestCreateTeacherProfile_CreateProfileFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, nil)
	mockProfileRepo.On("CreateProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateProfileParams")).Return(profilesdb.Profile{}, errors.New("database error"))

	// Act
	result, err := service.CreateTeacherProfile(context.Background(), &profiles.CreateTeacherProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane.smith@example.com",
		Position:     "Professor",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateTeacherProfile_CreateTeacherProfileFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockTeacherProfileRepo := &mocks.MockTeacherProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, mockTeacherProfileRepo, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("teacher")

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, nil)
	mockProfileRepo.On("CreateProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateProfileParams")).Return(expectedProfile, nil)
	mockTeacherProfileRepo.On("CreateTeacherProfile", mock.Anything, mock.AnythingOfType("profilesdb.CreateTeacherProfileParams")).Return(profilesdb.TeacherProfile{}, errors.New("database error"))

	// Act
	result, err := service.CreateTeacherProfile(context.Background(), &profiles.CreateTeacherProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane.smith@example.com",
		Position:     "Professor",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockTeacherProfileRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_InvalidRole(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("invalid-role")

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(expectedProfile, nil)

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid-session-token",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_StudentProfileNotFound(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockStudentProfileRepo := &mocks.MockStudentProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, mockStudentProfileRepo, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("student")

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(expectedProfile, nil)
	mockStudentProfileRepo.On("GetCompleteStudentProfile", mock.Anything, int64(1)).Return(profilesdb.GetCompleteStudentProfileRow{}, errors.New("student profile not found"))

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid-session-token",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockStudentProfileRepo.AssertExpectations(t)
}

func TestGetCompleteProfile_TeacherProfileNotFound(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockTeacherProfileRepo := &mocks.MockTeacherProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, mockTeacherProfileRepo, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	expectedProfile := createTestProfile("teacher")

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(expectedProfile, nil)
	mockTeacherProfileRepo.On("GetCompleteTeacherProfile", mock.Anything, int64(1)).Return(profilesdb.GetCompleteTeacherProfileRow{}, errors.New("teacher profile not found"))

	// Act
	result, err := service.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid-session-token",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
	mockTeacherProfileRepo.AssertExpectations(t)
}

func TestUpdateProfile_UpdateFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()
	existingProfile := createTestProfile("student")

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("GetProfileByUserId", mock.Anything, int64(1)).Return(existingProfile, nil)
	mockProfileRepo.On("UpdateProfile", mock.Anything, mock.AnythingOfType("profilesdb.UpdateProfileParams")).Return(profilesdb.Profile{}, errors.New("update failed"))

	// Act
	result, err := service.UpdateProfile(context.Background(), &profiles.UpdateProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    stringPtr("Jane"),
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateStudentProfile_CheckProfileExistsFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, errors.New("database error"))

	// Act
	result, err := service.CreateStudentProfile(context.Background(), &profiles.CreateStudentProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		GradeLevel:   "undergraduate",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}

func TestCreateTeacherProfile_CheckProfileExistsFails(t *testing.T) {
	// Arrange
	mockProfileRepo := &mocks.MockProfileRepository{}
	mockAuthServiceRepo := &mocks.MockAuthServiceRepository{}
	service := setupTestService(mockProfileRepo, nil, nil, mockAuthServiceRepo)

	expectedAuthResponse := createTestValidationResponse()

	mockAuthServiceRepo.On("ValidateUser", mock.Anything, mock.AnythingOfType("*auth.ValidateUserPayload")).Return(expectedAuthResponse, nil)
	mockProfileRepo.On("CheckProfileExists", mock.Anything, int64(1)).Return(false, errors.New("database error"))

	// Act
	result, err := service.CreateTeacherProfile(context.Background(), &profiles.CreateTeacherProfilePayload{
		SessionToken: "valid-session-token",
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane.smith@example.com",
		Position:     "Professor",
	})

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	mockAuthServiceRepo.AssertExpectations(t)
	mockProfileRepo.AssertExpectations(t)
}
