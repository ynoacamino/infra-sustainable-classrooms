package controllers

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

func TestValidateOwnership_Success(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil)

	// Act
	err := service.validateOwnership(1, 1)

	// Assert
	assert.NoError(t, err)
}

func TestValidateOwnership_PermissionDenied(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil)

	// Act
	err := service.validateOwnership(1, 2)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Permission denied")
}

func TestMapToStudentProfileResponse(t *testing.T) {
	// Arrange
	now := time.Now()
	profile := profilesdb.Profile{
		ID:        1,
		UserID:    123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: now, Valid: true},
	}

	studentProfile := profilesdb.StudentProfile{
		ID:         1,
		ProfileID:  1,
		GradeLevel: "undergraduate",
		Major:      pgtype.Text{String: "Computer Science", Valid: true},
	}

	// Act
	result := mapToStudentProfileResponse(profile, studentProfile)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "john.doe@example.com", result.Email)
	assert.Equal(t, "+1234567890", *result.Phone)
	assert.Equal(t, "https://example.com/avatar.jpg", *result.AvatarURL)
	assert.Equal(t, "Test bio", *result.Bio)
	assert.Equal(t, "undergraduate", result.GradeLevel)
	assert.Equal(t, "Computer Science", *result.Major)
	assert.True(t, result.IsActive)
	assert.Equal(t, now.UnixMilli(), result.CreatedAt)
	assert.Equal(t, now.UnixMilli(), *result.UpdatedAt)
}

func TestMapToStudentProfileResponse_NilOptionalFields(t *testing.T) {
	// Arrange
	now := time.Now()
	profile := profilesdb.Profile{
		ID:        1,
		UserID:    123,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     pgtype.Text{Valid: false},
		AvatarUrl: pgtype.Text{Valid: false},
		Bio:       pgtype.Text{Valid: false},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Valid: false},
	}

	studentProfile := profilesdb.StudentProfile{
		ID:         1,
		ProfileID:  1,
		GradeLevel: "undergraduate",
		Major:      pgtype.Text{Valid: false},
	}

	// Act
	result := mapToStudentProfileResponse(profile, studentProfile)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "john.doe@example.com", result.Email)
	assert.Nil(t, result.Phone)
	assert.Nil(t, result.AvatarURL)
	assert.Nil(t, result.Bio)
	assert.Equal(t, "undergraduate", result.GradeLevel)
	assert.Nil(t, result.Major)
	assert.True(t, result.IsActive)
	assert.Equal(t, now.UnixMilli(), result.CreatedAt)
	assert.Nil(t, result.UpdatedAt)
}

func TestMapToTeacherProfileResponse(t *testing.T) {
	// Arrange
	now := time.Now()
	profile := profilesdb.Profile{
		ID:        1,
		UserID:    123,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		Phone:     pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: now, Valid: true},
	}

	teacherProfile := profilesdb.TeacherProfile{
		ID:        1,
		ProfileID: 1,
		Position:  "Professor",
	}

	// Act
	result := mapToTeacherProfileResponse(profile, teacherProfile)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Smith", result.LastName)
	assert.Equal(t, "jane.smith@example.com", result.Email)
	assert.Equal(t, "+1234567890", *result.Phone)
	assert.Equal(t, "https://example.com/avatar.jpg", *result.AvatarURL)
	assert.Equal(t, "Test bio", *result.Bio)
	assert.Equal(t, "Professor", result.Position)
	assert.True(t, result.IsActive)
	assert.Equal(t, now.UnixMilli(), result.CreatedAt)
	assert.Equal(t, now.UnixMilli(), *result.UpdatedAt)
}

func TestMapToProfileResponse(t *testing.T) {
	// Arrange
	now := time.Now()
	profile := profilesdb.Profile{
		ID:        1,
		UserID:    123,
		Role:      "student",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: now, Valid: true},
	}

	// Act
	result := mapToProfileResponse(profile)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "student", result.Role)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "john.doe@example.com", result.Email)
	assert.Equal(t, "+1234567890", *result.Phone)
	assert.Equal(t, "https://example.com/avatar.jpg", *result.AvatarURL)
	assert.Equal(t, "Test bio", *result.Bio)
	assert.True(t, result.IsActive)
	assert.Equal(t, now.UnixMilli(), result.CreatedAt)
	assert.Equal(t, now.UnixMilli(), *result.UpdatedAt)
}

func TestMapToCompleteProfileResponse_Student(t *testing.T) {
	// Arrange
	now := time.Now()
	completeProfile := profilesdb.GetCompleteStudentProfileRow{
		UserID:     123,
		Role:       "student",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		Phone:      pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl:  pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:        pgtype.Text{String: "Test bio", Valid: true},
		IsActive:   true,
		CreatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		GradeLevel: "undergraduate",
		Major:      pgtype.Text{String: "Computer Science", Valid: true},
	}

	// Act
	result := mapToCompleteProfileResponse(completeProfile, "student")

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "student", result.Role)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "john.doe@example.com", result.Email)
	assert.Equal(t, "+1234567890", *result.Phone)
	assert.Equal(t, "https://example.com/avatar.jpg", *result.AvatarURL)
	assert.Equal(t, "Test bio", *result.Bio)
	assert.True(t, result.IsActive)
	assert.Equal(t, now.UnixMilli(), result.CreatedAt)
	assert.Equal(t, now.UnixMilli(), *result.UpdatedAt)
	assert.Equal(t, "undergraduate", *result.GradeLevel)
	assert.Equal(t, "Computer Science", *result.Major)
	assert.Nil(t, result.Position)
}

func TestMapToCompleteProfileResponse_Teacher(t *testing.T) {
	// Arrange
	now := time.Now()
	completeProfile := profilesdb.GetCompleteTeacherProfileRow{
		UserID:    123,
		Role:      "teacher",
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		Phone:     pgtype.Text{String: "+1234567890", Valid: true},
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		Position:  "Professor",
	}

	// Act
	result := mapToCompleteProfileResponse(completeProfile, "teacher")

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "teacher", result.Role)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Smith", result.LastName)
	assert.Equal(t, "jane.smith@example.com", result.Email)
	assert.Equal(t, "+1234567890", *result.Phone)
	assert.Equal(t, "https://example.com/avatar.jpg", *result.AvatarURL)
	assert.Equal(t, "Test bio", *result.Bio)
	assert.True(t, result.IsActive)
	assert.Equal(t, now.UnixMilli(), result.CreatedAt)
	assert.Equal(t, now.UnixMilli(), *result.UpdatedAt)
	assert.Equal(t, "Professor", *result.Position)
	assert.Nil(t, result.GradeLevel)
	assert.Nil(t, result.Major)
}

func TestMapToCompleteProfileResponse_InvalidRole(t *testing.T) {
	// Arrange
	completeProfile := "invalid-profile-type"

	// Act
	result := mapToCompleteProfileResponse(completeProfile, "invalid")

	// Assert
	assert.Nil(t, result)
}

func TestMapToCompleteProfileResponse_InvalidStudentCast(t *testing.T) {
	// Arrange
	completeProfile := "not-a-student-profile"

	// Act
	result := mapToCompleteProfileResponse(completeProfile, "student")

	// Assert
	assert.Nil(t, result)
}

func TestMapToCompleteProfileResponse_InvalidTeacherCast(t *testing.T) {
	// Arrange
	completeProfile := "not-a-teacher-profile"

	// Act
	result := mapToCompleteProfileResponse(completeProfile, "teacher")

	// Assert
	assert.Nil(t, result)
}

func TestMapToPublicProfileResponse(t *testing.T) {
	// Arrange
	publicProfile := profilesdb.GetPublicProfileByUserIdRow{
		UserID:    123,
		Role:      "student",
		FirstName: "John",
		LastName:  "Doe",
		AvatarUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
		Bio:       pgtype.Text{String: "Test bio", Valid: true},
		IsActive:  true,
	}

	// Act
	result := mapToPublicProfileResponse(publicProfile)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "student", result.Role)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Equal(t, "https://example.com/avatar.jpg", *result.AvatarURL)
	assert.Equal(t, "Test bio", *result.Bio)
	assert.True(t, result.IsActive)
}

func TestMapToPublicProfileResponse_NilOptionalFields(t *testing.T) {
	// Arrange
	publicProfile := profilesdb.GetPublicProfileByUserIdRow{
		UserID:    123,
		Role:      "student",
		FirstName: "John",
		LastName:  "Doe",
		AvatarUrl: pgtype.Text{Valid: false},
		Bio:       pgtype.Text{Valid: false},
		IsActive:  true,
	}

	// Act
	result := mapToPublicProfileResponse(publicProfile)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, int64(123), result.UserID)
	assert.Equal(t, "student", result.Role)
	assert.Equal(t, "John", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	assert.Nil(t, result.AvatarURL)
	assert.Nil(t, result.Bio)
	assert.True(t, result.IsActive)
}

func TestGetStringPtr_ValidText(t *testing.T) {
	// Arrange
	text := pgtype.Text{String: "test string", Valid: true}

	// Act
	result := getStringPtr(text)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "test string", *result)
}

func TestGetStringPtr_InvalidText(t *testing.T) {
	// Arrange
	text := pgtype.Text{String: "", Valid: false}

	// Act
	result := getStringPtr(text)

	// Assert
	assert.Nil(t, result)
}

func TestTimestampToMillis_ValidTimestamp(t *testing.T) {
	// Arrange
	now := time.Now()
	timestamp := pgtype.Timestamptz{Time: now, Valid: true}

	// Act
	result := timestampToMillis(timestamp)

	// Assert
	assert.Equal(t, now.UnixMilli(), result)
}

func TestTimestampToMillis_InvalidTimestamp(t *testing.T) {
	// Arrange
	timestamp := pgtype.Timestamptz{Valid: false}

	// Act
	result := timestampToMillis(timestamp)

	// Assert
	assert.Equal(t, int64(0), result)
}

func TestTimestampToMillisPtr_ValidTimestamp(t *testing.T) {
	// Arrange
	now := time.Now()
	timestamp := pgtype.Timestamptz{Time: now, Valid: true}

	// Act
	result := timestampToMillisPtr(timestamp)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, now.UnixMilli(), *result)
}

func TestTimestampToMillisPtr_InvalidTimestamp(t *testing.T) {
	// Arrange
	timestamp := pgtype.Timestamptz{Valid: false}

	// Act
	result := timestampToMillisPtr(timestamp)

	// Assert
	assert.Nil(t, result)
}

func TestStringToPgText_ValidString(t *testing.T) {
	// Arrange
	str := "test string"

	// Act
	result := stringToPgText(&str)

	// Assert
	assert.True(t, result.Valid)
	assert.Equal(t, "test string", result.String)
}

func TestStringToPgText_NilString(t *testing.T) {
	// Arrange
	var str *string = nil

	// Act
	result := stringToPgText(str)

	// Assert
	assert.False(t, result.Valid)
	assert.Equal(t, "", result.String)
}

func TestStringToPgTextDirect_ValidString(t *testing.T) {
	// Arrange
	str := "test string"

	// Act
	result := stringToPgTextDirect(str)

	// Assert
	assert.True(t, result.Valid)
	assert.Equal(t, "test string", result.String)
}

func TestStringToPgTextDirect_EmptyString(t *testing.T) {
	// Arrange
	str := ""

	// Act
	result := stringToPgTextDirect(str)

	// Assert
	assert.False(t, result.Valid)
	assert.Equal(t, "", result.String)
}

func TestGetStringOrDefault_WithNewValue(t *testing.T) {
	// Arrange
	newValue := "new value"
	existing := "existing value"

	// Act
	result := getStringOrDefault(&newValue, existing)

	// Assert
	assert.Equal(t, "new value", result)
}

func TestGetStringOrDefault_WithNilValue(t *testing.T) {
	// Arrange
	var newValue *string = nil
	existing := "existing value"

	// Act
	result := getStringOrDefault(newValue, existing)

	// Assert
	assert.Equal(t, "existing value", result)
}

func TestGetStringPtrOrDefault_WithNewValue(t *testing.T) {
	// Arrange
	newValue := "new value"
	existing := pgtype.Text{String: "existing value", Valid: true}

	// Act
	result := getStringPtrOrDefault(&newValue, existing)

	// Assert
	assert.True(t, result.Valid)
	assert.Equal(t, "new value", result.String)
}

func TestGetStringPtrOrDefault_WithNilValue(t *testing.T) {
	// Arrange
	var newValue *string = nil
	existing := pgtype.Text{String: "existing value", Valid: true}

	// Act
	result := getStringPtrOrDefault(newValue, existing)

	// Assert
	assert.True(t, result.Valid)
	assert.Equal(t, "existing value", result.String)
}
