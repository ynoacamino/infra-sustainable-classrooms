package controllers

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
)

// TestHelperFunctions tests the helper functions used in the service
func TestValidateSessionAndGetUserID(t *testing.T) {
	// This test would require setting up the service with mocks
	// and testing the validateSessionAndGetUserID helper function
	t.Skip("Helper function tests require more complex setup")
}

// TestConvertVideoToAPIVideo tests the video conversion helper
func TestConvertVideoToAPIVideo(t *testing.T) {
	t.Skip("Video conversion tests require service setup with mocks")
}

// TestRandomSelectVideos tests the random video selection function
func TestRandomSelectVideos(t *testing.T) {
	tests := []struct {
		name     string
		videos   []int
		amount   int
		expected int // expected length of result
	}{
		{
			name:     "select all videos when amount >= length",
			videos:   []int{1, 2, 3},
			amount:   5,
			expected: 3,
		},
		{
			name:     "select subset when amount < length",
			videos:   []int{1, 2, 3, 4, 5},
			amount:   3,
			expected: 3,
		},
		{
			name:     "select none when amount is 0",
			videos:   []int{1, 2, 3},
			amount:   0,
			expected: 0,
		},
		{
			name:     "select none when amount is negative",
			videos:   []int{1, 2, 3},
			amount:   -1,
			expected: 0,
		},
		{
			name:     "empty input returns empty",
			videos:   []int{},
			amount:   3,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := randomSelectVideos(tt.videos, tt.amount)
			assert.Equal(t, tt.expected, len(result))

			// Verify all results are from original slice
			for _, item := range result {
				assert.Contains(t, tt.videos, item)
			}
		})
	}
}

// TestGenerateObjectName tests the object name generation helper
func TestGenerateObjectName(t *testing.T) {
	service := &videolearningsrvc{}

	tests := []struct {
		name     string
		filename string
		userID   int64
	}{
		{
			name:     "normal filename",
			filename: "video.mp4",
			userID:   123,
		},
		{
			name:     "filename with spaces",
			filename: "my video file.mp4",
			userID:   456,
		},
		{
			name:     "filename with special characters",
			filename: "test@video#file!.mp4",
			userID:   789,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.generateObjectName(tt.filename, tt.userID)

			// Check that result is not empty
			assert.NotEmpty(t, result)

			// The result should contain the userID as string somewhere
			// Since each test has different userID, we need to check dynamically
			userIDStr := ""
			switch tt.userID {
			case 123:
				userIDStr = "123"
			case 456:
				userIDStr = "456"
			case 789:
				userIDStr = "789"
			}
			assert.Contains(t, result, userIDStr)

			// Check that it doesn't contain dangerous characters
			assert.NotContains(t, result, " ")
			assert.NotContains(t, result, "@")
			assert.NotContains(t, result, "#")
			assert.NotContains(t, result, "!")
		})
	}
}

// TestSanitizeFilename tests the filename sanitization helper
func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normal filename",
			input:    "video.mp4",
			expected: "video.mp4",
		},
		{
			name:     "filename with spaces",
			input:    "my video.mp4",
			expected: "my_video.mp4",
		},
		{
			name:     "filename with special characters",
			input:    "test@video#file!.mp4",
			expected: "test_video_file_.mp4",
		},
		{
			name:     "filename with allowed characters",
			input:    "test-video_file.mp4",
			expected: "test-video_file.mp4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizeFilename(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestDatabaseModelCreation tests creating test models
func TestDatabaseModelCreation(t *testing.T) {
	t.Run("create test video", func(t *testing.T) {
		video := createTestVideo()
		assert.Equal(t, int64(1), video.ID)
		assert.Equal(t, "Test Video", video.Title)
		assert.Equal(t, int64(1), video.UserID)
		assert.True(t, video.Description.Valid)
		assert.Equal(t, "Test description", video.Description.String)
	})

	t.Run("create test video comment", func(t *testing.T) {
		comment := createTestVideoComment()
		assert.Equal(t, int64(1), comment.ID)
		assert.Equal(t, int64(1), comment.VideoID)
		assert.Equal(t, int64(1), comment.UserID)
		assert.Equal(t, "Test Comment", comment.Title)
		assert.Equal(t, "This is a test comment", comment.Content)
	})

	t.Run("create test video category", func(t *testing.T) {
		category := createTestVideoCategory()
		assert.Equal(t, int64(1), category.ID)
		assert.Equal(t, "Education", category.Name)
	})

	t.Run("create test video tag", func(t *testing.T) {
		tag := createTestVideoTag()
		assert.Equal(t, int64(1), tag.ID)
		assert.Equal(t, "test-tag", tag.Name)
	})
}

// TestProfileCreation tests creating test profiles
func TestProfileCreation(t *testing.T) {
	t.Run("create complete profile", func(t *testing.T) {
		profile := createTestCompleteProfile("teacher")
		assert.Equal(t, int64(1), profile.UserID)
		assert.Equal(t, "teacher", profile.Role)
		assert.Equal(t, "John", profile.FirstName)
		assert.Equal(t, "Doe", profile.LastName)
		assert.Equal(t, "john.doe@example.com", profile.Email)
		assert.True(t, profile.IsActive)
		assert.NotNil(t, profile.Phone)
		assert.NotNil(t, profile.AvatarURL)
		assert.NotNil(t, profile.Bio)
	})

	t.Run("create public profile", func(t *testing.T) {
		profile := createTestPublicProfile()
		assert.Equal(t, int64(1), profile.UserID)
		assert.Equal(t, "John", profile.FirstName)
		assert.Equal(t, "Doe", profile.LastName)
		assert.NotNil(t, profile.AvatarURL)
		assert.NotNil(t, profile.Bio)
	})

	t.Run("create different roles", func(t *testing.T) {
		teacher := createTestCompleteProfile("teacher")
		student := createTestCompleteProfile("student")

		assert.Equal(t, "teacher", teacher.Role)
		assert.Equal(t, "student", student.Role)
	})
}

// Benchmark tests for performance-critical functions
func BenchmarkRandomSelectVideos(b *testing.B) {
	videos := make([]int, 1000)
	for i := range videos {
		videos[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randomSelectVideos(videos, 10)
	}
}

func BenchmarkSanitizeFilename(b *testing.B) {
	filename := "this@is#a$very%long^filename&with*many(special)characters+and-spaces.mp4"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sanitizeFilename(filename)
	}
}

// Test data factories for consistent test data creation
func createTestVideosByUser() videolearningdb.Video {
	return videolearningdb.Video{
		ID:           1,
		Title:        "User Video",
		UserID:       1,
		Views:        100,
		Likes:        50,
		ThumbObjName: pgtype.Text{String: "user_thumb.jpg", Valid: true},
		CategoryID:   1,
		CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

func createTestRecentVideo() videolearningdb.Video {
	return videolearningdb.Video{
		ID:           1,
		Title:        "Recent Video",
		UserID:       1,
		Views:        100,
		Likes:        50,
		ThumbObjName: pgtype.Text{String: "recent_thumb.jpg", Valid: true},
		CategoryID:   1,
		CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

func createTestUserCategoryLike() videolearningdb.GetUserCategoryLikesRow {
	return videolearningdb.GetUserCategoryLikesRow{
		Name:  "Education",
		Likes: pgtype.Int4{Int32: 10, Valid: true},
	}
}

func createTestUserVideoLike() videolearningdb.UserVideoLike {
	return videolearningdb.UserVideoLike{
		UserID:    1,
		VideoID:   1,
		Liked:     true,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// Test utilities for common assertions
func assertVideoListResponse(t *testing.T, expected, actual *videolearning.VideoList) {
	assert.NotNil(t, actual)
	assert.Equal(t, len(expected.Videos), len(actual.Videos))

	for i, expectedVideo := range expected.Videos {
		if i < len(actual.Videos) {
			assert.Equal(t, expectedVideo.ID, actual.Videos[i].ID)
			assert.Equal(t, expectedVideo.Title, actual.Videos[i].Title)
			assert.Equal(t, expectedVideo.Views, actual.Videos[i].Views)
			assert.Equal(t, expectedVideo.Likes, actual.Videos[i].Likes)
		}
	}
}

func assertSimpleResponse(t *testing.T, expected, actual *videolearning.SimpleResponse) {
	assert.NotNil(t, actual)
	assert.Equal(t, expected.Success, actual.Success)
	assert.Equal(t, expected.Message, actual.Message)
	if expected.ID != nil {
		assert.NotNil(t, actual.ID)
		assert.Equal(t, *expected.ID, *actual.ID)
	}
}
