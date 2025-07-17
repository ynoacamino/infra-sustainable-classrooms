package controllers

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/repositories/mocks"
)

// setupTestService creates a service instance with mocked repositories for testing
func setupTestService(
	videoRepo *mocks.MockVideoRepository,
	videoCommentRepo *mocks.MockVideoCommentRepository,
	videoCategoryRepo *mocks.MockVideoCategoryRepository,
	videoTagRepo *mocks.MockVideoTagRepository,
	userCategoryLikeRepo *mocks.MockUserCategoryLikeRepository,
	profilesServiceRepo *mocks.MockProfilesServiceRepository,
	cacheRepo *mocks.MockCacheRepository,
	storageRepo *mocks.MockStorageRepository,
) *videolearningsrvc {
	// If mocks are nil, create empty ones
	if videoRepo == nil {
		videoRepo = &mocks.MockVideoRepository{}
	}
	if videoCommentRepo == nil {
		videoCommentRepo = &mocks.MockVideoCommentRepository{}
	}
	if videoCategoryRepo == nil {
		videoCategoryRepo = &mocks.MockVideoCategoryRepository{}
	}
	if videoTagRepo == nil {
		videoTagRepo = &mocks.MockVideoTagRepository{}
	}
	if userCategoryLikeRepo == nil {
		userCategoryLikeRepo = &mocks.MockUserCategoryLikeRepository{}
	}
	if profilesServiceRepo == nil {
		profilesServiceRepo = &mocks.MockProfilesServiceRepository{}
	}
	if cacheRepo == nil {
		cacheRepo = &mocks.MockCacheRepository{}
	}
	if storageRepo == nil {
		storageRepo = &mocks.MockStorageRepository{}
	}

	return &videolearningsrvc{
		videoRepo:            videoRepo,
		videoCommentRepo:     videoCommentRepo,
		videoCategoryRepo:    videoCategoryRepo,
		videoTagRepo:         videoTagRepo,
		userCategoryLikeRepo: userCategoryLikeRepo,
		profileServiceRepo:   profilesServiceRepo,
		cacheRepo:            cacheRepo,
		// storageRepo:          storageRepo,
	}
}

// createTestCompleteProfile creates a test complete profile for use in tests
func createTestCompleteProfile(role string) *profiles.CompleteProfileResponse {
	updatedAt := time.Now().UnixMilli()
	return &profiles.CompleteProfileResponse{
		UserID:    1,
		Role:      role,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     stringPtr("+1234567890"),
		AvatarURL: stringPtr("https://example.com/avatar.jpg"),
		Bio:       stringPtr("Test bio"),
		IsActive:  true,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: &updatedAt,
	}
}

// createTestPublicProfile creates a test public profile for use in tests
func createTestPublicProfile() *profiles.PublicProfileResponse {
	return &profiles.PublicProfileResponse{
		UserID:    1,
		FirstName: "John",
		LastName:  "Doe",
		AvatarURL: stringPtr("https://example.com/avatar.jpg"),
		Bio:       stringPtr("Test bio"),
	}
}

// createTestVideo creates a test Video for use in tests
func createTestVideo() videolearningdb.Video {
	now := time.Now()
	return videolearningdb.Video{
		ID:           1,
		Title:        "Test Video",
		UserID:       1,
		Description:  pgtype.Text{String: "Test description", Valid: true},
		Views:        100,
		Likes:        50,
		VideoObjName: pgtype.Text{String: "test_video.mp4", Valid: true},
		ThumbObjName: pgtype.Text{String: "test_thumb.jpg", Valid: true},
		CategoryID:   1,
		CreatedAt:    pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:    pgtype.Timestamptz{Time: now, Valid: true},
	}
}

// createTestSearchVideo creates a test Video for use in search tests
func createTestSearchVideo() videolearningdb.Video {
	return videolearningdb.Video{
		ID:           1,
		Title:        "Test Video",
		UserID:       1,
		Views:        100,
		Likes:        50,
		ThumbObjName: pgtype.Text{String: "test_thumb.jpg", Valid: true},
		CategoryID:   1,
		CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestVideoComment creates a test VideoComment for use in tests
func createTestVideoComment() videolearningdb.VideoComment {
	return videolearningdb.VideoComment{
		ID:        1,
		VideoID:   1,
		UserID:    1,
		Title:     "Test Comment",
		Content:   "This is a test comment",
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestVideoCategory creates a test VideoCategory for use in tests
func createTestVideoCategory() videolearningdb.VideoCategory {
	return videolearningdb.VideoCategory{
		ID:        1,
		Name:      "Education",
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// createTestVideoTag creates a test VideoTag for use in tests
func createTestVideoTag() videolearningdb.VideoTag {
	return videolearningdb.VideoTag{
		ID:        1,
		Name:      "test-tag",
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
}

// Helper functions for tests
func stringPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}

// Test SearchVideos endpoint
func TestSearchVideos(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.SearchVideosPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository, *mocks.MockStorageRepository)
		expectedResult *videolearning.VideoList
		expectedError  error
	}{
		{
			name: "successful video search without category",
			payload: &videolearning.SearchVideosPayload{
				SessionToken: "valid_token",
				Query:        "test",
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				searchParams := videolearningdb.SearchVideosParams{
					Column1: pgtype.Text{String: "test", Valid: true},
					Column2: int64(0),
					Limit:   int32(10),
					Offset:  int32(0),
				}
				videos := []videolearningdb.Video{createTestSearchVideo()}
				videoRepo.On("SearchVideos", mock.Anything, searchParams).Return(videos, nil)

				// Setup cache expectations for video views and likes
				cacheRepo.On("GetInt", mock.Anything, "video:views:1").Return(int64(0), nil)
				cacheRepo.On("GetInt", mock.Anything, "video:likes:1").Return(int64(0), nil)

				// Setup cache expectations for presigned URL (return error to skip URL generation)
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:test_thumb.jpg").Return("", errors.New("not cached"))

				// Storage repository should fail to generate URL (so thumbnail URL stays empty)
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", "test_thumb.jpg", mock.Anything).Return("", errors.New("storage error"))

				// Mock the author name lookup
				profilesRepo.On("GetPublicProfileByID", mock.Anything, &profiles.GetPublicProfileByIDPayload{
					UserID: int64(1),
				}).Return(createTestPublicProfile(), nil)
			},
			expectedResult: &videolearning.VideoList{
				Videos: []*videolearning.Video{
					{
						ID:           1,
						Title:        "Test Video",
						Author:       "John Doe",
						Views:        100,
						Likes:        50,
						ThumbnailURL: "",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "successful video search with category",
			payload: &videolearning.SearchVideosPayload{
				SessionToken: "valid_token",
				Query:        "test",
				CategoryID:   int64Ptr(1),
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				searchParams := videolearningdb.SearchVideosParams{
					Column1: pgtype.Text{String: "test", Valid: true},
					Column2: int64(1),
					Limit:   int32(10),
					Offset:  int32(0),
				}
				videos := []videolearningdb.Video{createTestSearchVideo()}
				videoRepo.On("SearchVideos", mock.Anything, searchParams).Return(videos, nil)

				// Setup cache expectations for video views and likes
				cacheRepo.On("GetInt", mock.Anything, "video:views:1").Return(int64(0), nil)
				cacheRepo.On("GetInt", mock.Anything, "video:likes:1").Return(int64(0), nil)

				// Setup cache expectations for presigned URL (return error to skip URL generation)
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:test_thumb.jpg").Return("", errors.New("not cached"))

				// Storage repository should fail to generate URL (so thumbnail URL stays empty)
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", "test_thumb.jpg", mock.Anything).Return("", errors.New("storage error"))

				// Mock the author name lookup
				profilesRepo.On("GetPublicProfileByID", mock.Anything, &profiles.GetPublicProfileByIDPayload{
					UserID: int64(1),
				}).Return(createTestPublicProfile(), nil)
			},
			expectedResult: &videolearning.VideoList{
				Videos: []*videolearning.Video{
					{
						ID:           1,
						Title:        "Test Video",
						Author:       "John Doe",
						Views:        100,
						Likes:        50,
						ThumbnailURL: "",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "invalid session",
			payload: &videolearning.SearchVideosPayload{
				SessionToken: "invalid_token",
				Query:        "test",
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "invalid_token",
				}).Return((*profiles.CompleteProfileResponse)(nil), errors.New("invalid session"))
			},
			expectedResult: nil,
			expectedError:  videolearning.InvalidSession("invalid session"),
		},
		{
			name: "database error",
			payload: &videolearning.SearchVideosPayload{
				SessionToken: "valid_token",
				Query:        "test",
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				searchParams := videolearningdb.SearchVideosParams{
					Column1: pgtype.Text{String: "test", Valid: true},
					Column2: int64(0),
					Limit:   int32(10),
					Offset:  int32(0),
				}
				videoRepo.On("SearchVideos", mock.Anything, searchParams).Return([]videolearningdb.Video(nil), errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  videolearning.ServiceUnavailable("failed to search videos"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, profilesRepo, cacheRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, nil, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.SearchVideos(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult.Videos), len(result.Videos))
				if len(result.Videos) > 0 {
					assert.Equal(t, tt.expectedResult.Videos[0].ID, result.Videos[0].ID)
					assert.Equal(t, tt.expectedResult.Videos[0].Title, result.Videos[0].Title)
				}
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
			storageRepo.AssertExpectations(t)
		})
	}
}

// Test GetVideo endpoint
func TestGetVideo(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetVideoPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockVideoTagRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository, *mocks.MockStorageRepository)
		expectedResult *videolearning.VideoDetails
		expectedError  error
	}{
		{
			name: "successful get video",
			payload: &videolearning.GetVideoPayload{
				SessionToken: "valid_token",
				VideoID:      1,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				videoData := createTestVideo()
				videoRepo.On("GetVideoByID", mock.Anything, int64(1)).Return(videoData, nil)

				// Mock cache operations for views and likes
				cacheRepo.On("GetInt", mock.Anything, "video:views:1").Return(int64(10), nil)
				cacheRepo.On("GetInt", mock.Anything, "video:likes:1").Return(int64(5), nil)
				cacheRepo.On("IncrBy", mock.Anything, "video:views:1", int64(1)).Return(int64(11), nil)

				// Mock presigned URL generation
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-videos-confirmed:test_video.mp4").Return("", errors.New("not found"))
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:test_thumb.jpg").Return("", errors.New("not found"))

				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-videos-confirmed", "test_video.mp4", mock.Anything).Return("https://example.com/video.mp4", nil)
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", "test_thumb.jpg", mock.Anything).Return("https://example.com/thumb.jpg", nil)

				cacheRepo.On("Set", mock.Anything, "presigned_url:video-learning-videos-confirmed:test_video.mp4", "https://example.com/video.mp4", mock.Anything).Return(nil)
				cacheRepo.On("Set", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:test_thumb.jpg", "https://example.com/thumb.jpg", mock.Anything).Return(nil)

				// Mock profile service for author name
				profilesRepo.On("GetPublicProfileByID", mock.Anything, &profiles.GetPublicProfileByIDPayload{
					UserID: int64(1),
				}).Return(createTestPublicProfile(), nil)

				// Mock tag retrieval
				tags := []videolearningdb.VideoTag{createTestVideoTag()}
				tagRepo.On("GetTagsByVideoID", mock.Anything, int64(1)).Return(tags, nil)
			},
			expectedResult: &videolearning.VideoDetails{
				ID:           1,
				Title:        "Test Video",
				Description:  "Test description",
				Author:       "John Doe",
				Views:        110, // 100 + 10 cached
				Likes:        55,  // 50 + 5 cached
				VideoURL:     "https://example.com/video.mp4",
				ThumbnailURL: "https://example.com/thumb.jpg",
				UploadDate:   1234567890000,
				CategoryID:   1,
				TagIds:       []int64{1},
			},
			expectedError: nil,
		},
		{
			name: "video not found",
			payload: &videolearning.GetVideoPayload{
				SessionToken: "valid_token",
				VideoID:      999,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				videoRepo.On("GetVideoByID", mock.Anything, int64(999)).Return(videolearningdb.Video{}, errors.New("video not found"))
			},
			expectedResult: nil,
			expectedError:  videolearning.VideoNotFound("video not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			tagRepo := &mocks.MockVideoTagRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, tagRepo, profilesRepo, cacheRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, tagRepo, nil, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.GetVideo(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.Title, result.Title)
				assert.Equal(t, tt.expectedResult.Author, result.Author)
				assert.Equal(t, tt.expectedResult.Views, result.Views)
				assert.Equal(t, tt.expectedResult.Likes, result.Likes)
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			tagRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
			storageRepo.AssertExpectations(t)
		})
	}
}

// Test CreateComment endpoint
func TestCreateComment(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.CreateCommentPayload
		setupMocks     func(*mocks.MockVideoCommentRepository, *mocks.MockProfilesServiceRepository)
		expectedResult *videolearning.SimpleResponse
		expectedError  error
	}{
		{
			name: "successful comment creation",
			payload: &videolearning.CreateCommentPayload{
				SessionToken: "valid_token",
				VideoID:      1,
				Title:        "Great video!",
				Body:         "This video was very helpful.",
			},
			setupMocks: func(commentRepo *mocks.MockVideoCommentRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				createParams := videolearningdb.CreateCommentParams{
					VideoID: 1,
					UserID:  1,
					Title:   "Great video!",
					Content: "This video was very helpful.",
				}
				comment := createTestVideoComment()
				commentRepo.On("CreateComment", mock.Anything, createParams).Return(comment, nil)
			},
			expectedResult: &videolearning.SimpleResponse{
				Success: true,
				Message: "Comment created successfully",
				ID:      int64Ptr(1),
			},
			expectedError: nil,
		},
		{
			name: "invalid session",
			payload: &videolearning.CreateCommentPayload{
				SessionToken: "invalid_token",
				VideoID:      1,
				Title:        "Great video!",
				Body:         "This video was very helpful.",
			},
			setupMocks: func(commentRepo *mocks.MockVideoCommentRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "invalid_token",
				}).Return((*profiles.CompleteProfileResponse)(nil), errors.New("invalid session"))
			},
			expectedResult: nil,
			expectedError:  videolearning.InvalidSession("invalid session"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			commentRepo := &mocks.MockVideoCommentRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(commentRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, commentRepo, nil, nil, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.CreateComment(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.Success, result.Success)
				assert.Equal(t, tt.expectedResult.Message, result.Message)
			}

			// Verify mocks
			commentRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test GetComments endpoint
func TestGetComments(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetCommentsPayload
		setupMocks     func(*mocks.MockVideoCommentRepository, *mocks.MockProfilesServiceRepository)
		expectedResult *videolearning.CommentList
		expectedError  error
	}{
		{
			name: "successful get comments",
			payload: &videolearning.GetCommentsPayload{
				SessionToken: "valid_token",
				VideoID:      1,
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(commentRepo *mocks.MockVideoCommentRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				getParams := videolearningdb.GetCommentsForVideoParams{
					VideoID: 1,
					Limit:   10,
					Offset:  0,
				}
				comments := []videolearningdb.VideoComment{createTestVideoComment()}
				commentRepo.On("GetCommentsForVideo", mock.Anything, getParams).Return(comments, nil)

				// Mock profile service for author name
				profilesRepo.On("GetPublicProfileByID", mock.Anything, &profiles.GetPublicProfileByIDPayload{
					UserID: int64(1),
				}).Return(createTestPublicProfile(), nil)
			},
			expectedResult: &videolearning.CommentList{
				Comments: []*videolearning.Comment{
					{
						ID:      1,
						Author:  "John Doe",
						Date:    1234567890000,
						Title:   "Test Comment",
						Body:    "This is a test comment",
						VideoID: 1,
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			commentRepo := &mocks.MockVideoCommentRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(commentRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, commentRepo, nil, nil, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.GetComments(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult.Comments), len(result.Comments))
				if len(result.Comments) > 0 {
					assert.Equal(t, tt.expectedResult.Comments[0].ID, result.Comments[0].ID)
					assert.Equal(t, tt.expectedResult.Comments[0].Author, result.Comments[0].Author)
					assert.Equal(t, tt.expectedResult.Comments[0].Title, result.Comments[0].Title)
				}
			}

			// Verify mocks
			commentRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test ToggleVideoLike endpoint
func TestToggleVideoLike(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.ToggleVideoLikePayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockUserCategoryLikeRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository)
		expectedResult *videolearning.SimpleResponse
		expectedError  error
	}{
		{
			name: "successful like toggle - like video",
			payload: &videolearning.ToggleVideoLikePayload{
				SessionToken: "valid_token",
				VideoID:      1,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, userCatRepo *mocks.MockUserCategoryLikeRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				videoData := createTestVideo()
				videoRepo.On("GetVideoByID", mock.Anything, int64(1)).Return(videoData, nil)

				// Cache key for user like status
				cacheKey := "user:like:1:1"
				cacheRepo.On("Exists", mock.Anything, cacheKey).Return(false, nil)

				// Check database for existing like - return sql.ErrNoRows for no existing like
				likeParams := videolearningdb.GetUserVideoLikeParams{
					UserID:  1,
					VideoID: 1,
				}
				userCatRepo.On("GetUserVideoLike", mock.Anything, likeParams).Return(videolearningdb.UserVideoLike{}, sql.ErrNoRows)

				// Set cache for new like
				cacheRepo.On("Set", mock.Anything, cacheKey, "1", mock.Anything).Return(nil)

				// Increment cached video likes - first call succeeds
				cacheRepo.On("IncrBy", mock.Anything, "video:likes:1", int64(1)).Return(int64(51), nil)

				// Increment cached user category likes - first call succeeds
				cacheRepo.On("IncrBy", mock.Anything, "user:category:likes:1:1", int64(1)).Return(int64(1), nil)
			},
			expectedResult: &videolearning.SimpleResponse{
				Success: true,
				Message: "Video liked",
			},
			expectedError: nil,
		},
		{
			name: "invalid session",
			payload: &videolearning.ToggleVideoLikePayload{
				SessionToken: "invalid_token",
				VideoID:      1,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, userCatRepo *mocks.MockUserCategoryLikeRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "invalid_token",
				}).Return((*profiles.CompleteProfileResponse)(nil), errors.New("invalid session"))
			},
			expectedResult: nil,
			expectedError:  videolearning.InvalidSession("invalid session"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			userCatRepo := &mocks.MockUserCategoryLikeRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			tt.setupMocks(videoRepo, userCatRepo, profilesRepo, cacheRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, userCatRepo, profilesRepo, cacheRepo, nil)

			// Call method
			result, err := service.ToggleVideoLike(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.Success, result.Success)
				assert.Equal(t, tt.expectedResult.Message, result.Message)
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			userCatRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
		})
	}
}

// Test GetAllCategories endpoint
func TestGetAllCategories(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetAllCategoriesPayload
		setupMocks     func(*mocks.MockVideoCategoryRepository, *mocks.MockProfilesServiceRepository)
		expectedResult []*videolearning.VideoCategory
		expectedError  error
	}{
		{
			name: "successful get all categories",
			payload: &videolearning.GetAllCategoriesPayload{
				SessionToken: "valid_token",
			},
			setupMocks: func(categoryRepo *mocks.MockVideoCategoryRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				categories := []videolearningdb.VideoCategory{createTestVideoCategory()}
				categoryRepo.On("GetAllCategories", mock.Anything).Return(categories, nil)
			},
			expectedResult: []*videolearning.VideoCategory{
				{
					ID:   1,
					Name: "Education",
				},
			},
			expectedError: nil,
		},
		{
			name: "invalid session",
			payload: &videolearning.GetAllCategoriesPayload{
				SessionToken: "invalid_token",
			},
			setupMocks: func(categoryRepo *mocks.MockVideoCategoryRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "invalid_token",
				}).Return((*profiles.CompleteProfileResponse)(nil), errors.New("invalid session"))
			},
			expectedResult: nil,
			expectedError:  videolearning.InvalidSession("invalid session"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			categoryRepo := &mocks.MockVideoCategoryRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(categoryRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, nil, categoryRepo, nil, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.GetAllCategories(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult), len(result))
				if len(result) > 0 {
					assert.Equal(t, tt.expectedResult[0].ID, result[0].ID)
					assert.Equal(t, tt.expectedResult[0].Name, result[0].Name)
				}
			}

			// Verify mocks
			categoryRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test GetAllTags endpoint
func TestGetAllTags(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetAllTagsPayload
		setupMocks     func(*mocks.MockVideoTagRepository, *mocks.MockProfilesServiceRepository)
		expectedResult []*videolearning.VideoTag
		expectedError  error
	}{
		{
			name: "successful get all tags",
			payload: &videolearning.GetAllTagsPayload{
				SessionToken: "valid_token",
			},
			setupMocks: func(tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				tags := []videolearningdb.VideoTag{createTestVideoTag()}
				tagRepo.On("GetAllTags", mock.Anything).Return(tags, nil)
			},
			expectedResult: []*videolearning.VideoTag{
				{
					ID:   1,
					Name: "test-tag",
				},
			},
			expectedError: nil,
		},
		{
			name: "database error",
			payload: &videolearning.GetAllTagsPayload{
				SessionToken: "valid_token",
			},
			setupMocks: func(tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				tagRepo.On("GetAllTags", mock.Anything).Return([]videolearningdb.VideoTag(nil), errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  videolearning.ServiceUnavailable("failed to get tags"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			tagRepo := &mocks.MockVideoTagRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(tagRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, nil, nil, tagRepo, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.GetAllTags(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult), len(result))
				if len(result) > 0 {
					assert.Equal(t, tt.expectedResult[0].ID, result[0].ID)
					assert.Equal(t, tt.expectedResult[0].Name, result[0].Name)
				}
			}

			// Verify mocks
			tagRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test InitialUpload endpoint (teacher only)
// func TestInitialUpload(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		payload        *videolearning.InitialUploadPayload
// 		setupMocks     func(*mocks.MockProfilesServiceRepository, *mocks.MockStorageRepository)
// 		expectedResult *videolearning.UploadResponse
// 		expectedError  error
// 	}{
// 		{
// 			name: "successful initial upload by teacher",
// 			payload: &videolearning.InitialUploadPayload{
// 				SessionToken: "teacher_token",
// 				Filename:     "test_video.mp4",
// 				File:         []byte("fake video content"),
// 			},
// 			setupMocks: func(profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository) {
// 				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
// 					SessionToken: "teacher_token",
// 				}).Return(createTestCompleteProfile("teacher"), nil)

// 				storageRepo.On("UploadFile", mock.Anything, "video-learning-videos-staging", mock.AnythingOfType("string"), mock.Anything, int64(18), "video/mp4").Return(nil)
// 			},
// 			expectedResult: &videolearning.UploadResponse{
// 				ObjectName: "", // We'll check that it's not empty in the assertion
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name: "permission denied - not a teacher",
// 			payload: &videolearning.InitialUploadPayload{
// 				SessionToken: "student_token",
// 				Filename:     "test_video.mp4",
// 				File:         []byte("fake video content"),
// 			},
// 			setupMocks: func(profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository) {
// 				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
// 					SessionToken: "student_token",
// 				}).Return(createTestCompleteProfile("student"), nil)
// 			},
// 			expectedResult: nil,
// 			expectedError:  videolearning.PermissionDenied("only teachers can upload videos"),
// 		},
// 		{
// 			name: "upload failed",
// 			payload: &videolearning.InitialUploadPayload{
// 				SessionToken: "teacher_token",
// 				Filename:     "test_video.mp4",
// 				File:         []byte("fake video content"),
// 			},
// 			setupMocks: func(profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository) {
// 				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
// 					SessionToken: "teacher_token",
// 				}).Return(createTestCompleteProfile("teacher"), nil)

// 				storageRepo.On("UploadFile", mock.Anything, "video-learning-videos-staging", mock.AnythingOfType("string"), mock.Anything, int64(18), "video/mp4").Return(errors.New("upload failed"))
// 			},
// 			expectedResult: nil,
// 			expectedError:  videolearning.UploadFailed("failed to upload video"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Setup mocks
// 			profilesRepo := &mocks.MockProfilesServiceRepository{}
// 			storageRepo := &mocks.MockStorageRepository{}
// 			tt.setupMocks(profilesRepo, storageRepo)

// 			// Create service
// 			service := setupTestService(nil, nil, nil, nil, nil, profilesRepo, nil, storageRepo)

// 			// Call method
// 			result, err := service.InitialUpload(context.Background(), tt.payload)

// 			// Assertions
// 			if tt.expectedError != nil {
// 				assert.Error(t, err)
// 				assert.Equal(t, tt.expectedError.Error(), err.Error())
// 				assert.Nil(t, result)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, result)
// 				assert.NotEmpty(t, result.ObjectName)
// 			}

// 			// Verify mocks
// 			profilesRepo.AssertExpectations(t)
// 			storageRepo.AssertExpectations(t)
// 		})
// 	}
// }

// Test CompleteUpload endpoint (teacher only)
func TestCompleteUpload(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.CompleteUploadPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockVideoTagRepository, *mocks.MockProfilesServiceRepository, *mocks.MockStorageRepository)
		expectedResult *videolearning.SimpleResponse
		expectedError  error
	}{
		{
			name: "successful complete upload by teacher",
			payload: &videolearning.CompleteUploadPayload{
				SessionToken:        "teacher_token",
				Title:               "Test Video",
				Description:         stringPtr("Test description"),
				VideoObjectName:     "test_video.mp4",
				ThumbnailObjectName: "test_thumb.jpg",
				CategoryID:          1,
				Tags:                []string{"education", "test"},
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "teacher_token",
				}).Return(createTestCompleteProfile("teacher"), nil)

				// Mock file operations
				storageRepo.On("CopyFile", mock.Anything, "video-learning-videos-staging", "test_video.mp4", "video-learning-videos-confirmed", "test_video.mp4").Return(nil)
				storageRepo.On("CopyFile", mock.Anything, "video-learning-thumbnails-staging", "test_thumb.jpg", "video-learning-thumbnails-confirmed", "test_thumb.jpg").Return(nil)

				// Mock video creation
				createParams := videolearningdb.CreateVideoParams{
					Title:        "Test Video",
					UserID:       1,
					Description:  pgtype.Text{String: "Test description", Valid: true},
					Views:        0,
					Likes:        0,
					VideoObjName: pgtype.Text{String: "test_video.mp4", Valid: true},
					ThumbObjName: pgtype.Text{String: "test_thumb.jpg", Valid: true},
					CategoryID:   1,
				}
				video := createTestVideo()
				videoRepo.On("CreateVideo", mock.Anything, createParams).Return(video, nil)

				// Mock tag operations
				tag1 := videolearningdb.VideoTag{ID: 1, Name: "education"}
				tag2 := videolearningdb.VideoTag{ID: 2, Name: "test"}
				tagRepo.On("GetOrCreateTag", mock.Anything, "education").Return(tag1, nil)
				tagRepo.On("GetOrCreateTag", mock.Anything, "test").Return(tag2, nil)

				videoRepo.On("AssignTagToVideo", mock.Anything, videolearningdb.AssignTagToVideoParams{
					VideoID: 1,
					TagID:   1,
				}).Return(nil)
				videoRepo.On("AssignTagToVideo", mock.Anything, videolearningdb.AssignTagToVideoParams{
					VideoID: 1,
					TagID:   2,
				}).Return(nil)

				// Mock cleanup operations
				storageRepo.On("DeleteFile", mock.Anything, "video-learning-videos-staging", "test_video.mp4").Return(nil)
				storageRepo.On("DeleteFile", mock.Anything, "video-learning-thumbnails-staging", "test_thumb.jpg").Return(nil)
			},
			expectedResult: &videolearning.SimpleResponse{
				Success: true,
				Message: "Video upload completed successfully",
				ID:      int64Ptr(1),
			},
			expectedError: nil,
		},
		{
			name: "permission denied - not a teacher",
			payload: &videolearning.CompleteUploadPayload{
				SessionToken:        "student_token",
				Title:               "Test Video",
				VideoObjectName:     "test_video.mp4",
				ThumbnailObjectName: "test_thumb.jpg",
				CategoryID:          1,
				Tags:                []string{"education"},
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "student_token",
				}).Return(createTestCompleteProfile("student"), nil)
			},
			expectedResult: nil,
			expectedError:  videolearning.PermissionDenied("only teachers can upload videos"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			tagRepo := &mocks.MockVideoTagRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, tagRepo, profilesRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, tagRepo, nil, profilesRepo, nil, storageRepo)

			// Call method
			result, err := service.CompleteUpload(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.Success, result.Success)
				assert.Equal(t, tt.expectedResult.Message, result.Message)
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			tagRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			storageRepo.AssertExpectations(t)
		})
	}
}

// Test DeleteVideo endpoint (teacher only)
func TestDeleteVideo(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.DeleteVideoPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockProfilesServiceRepository, *mocks.MockStorageRepository, *mocks.MockCacheRepository)
		expectedResult *videolearning.SimpleResponse
		expectedError  error
	}{
		{
			name: "successful video deletion by owner",
			payload: &videolearning.DeleteVideoPayload{
				SessionToken: "teacher_token",
				VideoID:      1,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository, cacheRepo *mocks.MockCacheRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "teacher_token",
				}).Return(createTestCompleteProfile("teacher"), nil)

				videoData := createTestVideo()
				videoRepo.On("GetVideoByID", mock.Anything, int64(1)).Return(videoData, nil)

				// Mock file deletion
				storageRepo.On("DeleteFile", mock.Anything, "video-learning-videos-confirmed", "test_video.mp4").Return(nil)
				storageRepo.On("DeleteFile", mock.Anything, "video-learning-thumbnails-confirmed", "test_thumb.jpg").Return(nil)

				// Mock cache clearing
				cacheRepo.On("Delete", mock.Anything, "presigned_url:video-learning-videos-confirmed:test_video.mp4").Return(nil)
				cacheRepo.On("Delete", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:test_thumb.jpg").Return(nil)

				// Mock video deletion from database
				videoRepo.On("DeleteVideo", mock.Anything, int64(1)).Return(nil)
			},
			expectedResult: &videolearning.SimpleResponse{
				Success: true,
				Message: "Video deleted successfully",
			},
			expectedError: nil,
		},
		{
			name: "permission denied - not owner",
			payload: &videolearning.DeleteVideoPayload{
				SessionToken: "other_teacher_token",
				VideoID:      1,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository, cacheRepo *mocks.MockCacheRepository) {
				// Different user ID (2) trying to delete video owned by user 1
				otherTeacher := createTestCompleteProfile("teacher")
				otherTeacher.UserID = 2
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "other_teacher_token",
				}).Return(otherTeacher, nil)

				videoData := createTestVideo()
				videoRepo.On("GetVideoByID", mock.Anything, int64(1)).Return(videoData, nil)
			},
			expectedResult: nil,
			expectedError:  videolearning.PermissionDenied("you can only delete your own videos"),
		},
		{
			name: "video not found",
			payload: &videolearning.DeleteVideoPayload{
				SessionToken: "teacher_token",
				VideoID:      999,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, storageRepo *mocks.MockStorageRepository, cacheRepo *mocks.MockCacheRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "teacher_token",
				}).Return(createTestCompleteProfile("teacher"), nil)

				videoRepo.On("GetVideoByID", mock.Anything, int64(999)).Return(videolearningdb.Video{}, errors.New("video not found"))
			},
			expectedResult: nil,
			expectedError:  videolearning.VideoNotFound("video not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			tt.setupMocks(videoRepo, profilesRepo, storageRepo, cacheRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, nil, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.DeleteVideo(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.Success, result.Success)
				assert.Equal(t, tt.expectedResult.Message, result.Message)
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			storageRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
		})
	}
}

// Test GetVideosByCategory endpoint
func TestGetVideosByCategory(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetVideosByCategoryPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository, *mocks.MockStorageRepository)
		expectedResult *videolearning.VideoList
		expectedError  error
	}{
		{
			name: "successful get videos by category",
			payload: &videolearning.GetVideosByCategoryPayload{
				SessionToken: "valid_token",
				CategoryID:   1,
				Amount:       5,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				categoryVideos := []videolearningdb.Video{
					{
						ID:           1,
						Title:        "Category Video 1",
						UserID:       1,
						Views:        100,
						Likes:        50,
						ThumbObjName: pgtype.Text{String: "thumb1.jpg", Valid: true},
						CategoryID:   1,
						CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
						UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
					},
				}
				videoRepo.On("GetVideosByCategory", mock.Anything, int64(1)).Return(categoryVideos, nil)

				// Cache expectations for video views and likes
				cacheRepo.On("GetInt", mock.Anything, "video:views:1").Return(int64(0), nil)
				cacheRepo.On("GetInt", mock.Anything, "video:likes:1").Return(int64(0), nil)

				// Cache expectations for presigned URL (return error to skip URL generation)
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:thumb1.jpg").Return("", errors.New("not cached"))

				// Storage repository should fail to generate URL (so thumbnail URL stays empty)
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", "thumb1.jpg", mock.Anything).Return("", errors.New("storage error"))

				// Mock the author name lookup
				profilesRepo.On("GetPublicProfileByID", mock.Anything, &profiles.GetPublicProfileByIDPayload{
					UserID: int64(1),
				}).Return(createTestPublicProfile(), nil)
			},
			expectedResult: &videolearning.VideoList{
				Videos: []*videolearning.Video{
					{
						ID:           1,
						Title:        "Category Video 1",
						Author:       "John Doe",
						Views:        100,
						Likes:        50,
						ThumbnailURL: "",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "database error",
			payload: &videolearning.GetVideosByCategoryPayload{
				SessionToken: "valid_token",
				CategoryID:   1,
				Amount:       5,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				videoRepo.On("GetVideosByCategory", mock.Anything, int64(1)).Return([]videolearningdb.Video(nil), errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  videolearning.ServiceUnavailable("failed to get videos by category"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, profilesRepo, cacheRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, nil, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.GetVideosByCategory(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult.Videos), len(result.Videos))
				if len(result.Videos) > 0 {
					assert.Equal(t, tt.expectedResult.Videos[0].ID, result.Videos[0].ID)
					assert.Equal(t, tt.expectedResult.Videos[0].Title, result.Videos[0].Title)
				}
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
			storageRepo.AssertExpectations(t)
		})
	}
}

// Test GetSimilarVideos endpoint
func TestGetSimilarVideos(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetSimilarVideosPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository, *mocks.MockStorageRepository)
		expectedResult *videolearning.VideoList
		expectedError  error
	}{
		{
			name: "successful get similar videos",
			payload: &videolearning.GetSimilarVideosPayload{
				SessionToken: "valid_token",
				VideoID:      1,
				Amount:       3,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				similarVideos := []videolearningdb.Video{
					{
						ID:           2,
						Title:        "Similar Video",
						UserID:       1,
						Views:        80,
						Likes:        40,
						ThumbObjName: pgtype.Text{String: "similar_thumb.jpg", Valid: true},
						CategoryID:   1,
						CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
						UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
					},
				}
				videoRepo.On("GetSimilarVideos", mock.Anything, int64(1)).Return(similarVideos, nil)

				// Cache expectations for video views and likes (note video ID 2)
				cacheRepo.On("GetInt", mock.Anything, "video:views:2").Return(int64(0), nil)
				cacheRepo.On("GetInt", mock.Anything, "video:likes:2").Return(int64(0), nil)

				// Cache expectations for presigned URL (return error to skip URL generation)
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:similar_thumb.jpg").Return("", errors.New("not cached"))

				// Storage repository should fail to generate URL (so thumbnail URL stays empty)
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", "similar_thumb.jpg", mock.Anything).Return("", errors.New("storage error"))

				// Mock the author name lookup
				profilesRepo.On("GetPublicProfileByID", mock.Anything, &profiles.GetPublicProfileByIDPayload{
					UserID: int64(1),
				}).Return(createTestPublicProfile(), nil)
			},
			expectedResult: &videolearning.VideoList{
				Videos: []*videolearning.Video{
					{
						ID:           2,
						Title:        "Similar Video",
						Author:       "John Doe",
						Views:        80,
						Likes:        40,
						ThumbnailURL: "",
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, profilesRepo, cacheRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, nil, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.GetSimilarVideos(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult.Videos), len(result.Videos))
				if len(result.Videos) > 0 {
					assert.Equal(t, tt.expectedResult.Videos[0].ID, result.Videos[0].ID)
					assert.Equal(t, tt.expectedResult.Videos[0].Title, result.Videos[0].Title)
				}
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test DeleteComment endpoint
func TestDeleteComment(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.DeleteCommentPayload
		setupMocks     func(*mocks.MockVideoCommentRepository, *mocks.MockProfilesServiceRepository)
		expectedResult *videolearning.SimpleResponse
		expectedError  error
	}{
		{
			name: "successful comment deletion",
			payload: &videolearning.DeleteCommentPayload{
				SessionToken: "valid_token",
				CommentID:    1,
			},
			setupMocks: func(commentRepo *mocks.MockVideoCommentRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				deleteParams := videolearningdb.DeleteCommentParams{
					ID:     1,
					UserID: 1,
				}
				commentRepo.On("DeleteComment", mock.Anything, deleteParams).Return(nil)
			},
			expectedResult: &videolearning.SimpleResponse{
				Success: true,
				Message: "Comment deleted successfully",
			},
			expectedError: nil,
		},
		{
			name: "database error",
			payload: &videolearning.DeleteCommentPayload{
				SessionToken: "valid_token",
				CommentID:    1,
			},
			setupMocks: func(commentRepo *mocks.MockVideoCommentRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				deleteParams := videolearningdb.DeleteCommentParams{
					ID:     1,
					UserID: 1,
				}
				commentRepo.On("DeleteComment", mock.Anything, deleteParams).Return(errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  videolearning.ServiceUnavailable("failed to delete comment"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			commentRepo := &mocks.MockVideoCommentRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(commentRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, commentRepo, nil, nil, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.DeleteComment(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.Success, result.Success)
				assert.Equal(t, tt.expectedResult.Message, result.Message)
			}

			// Verify mocks
			commentRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test GetOrCreateCategory endpoint
func TestGetOrCreateCategory(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetOrCreateCategoryPayload
		setupMocks     func(*mocks.MockVideoCategoryRepository, *mocks.MockProfilesServiceRepository)
		expectedResult *videolearning.VideoCategory
		expectedError  error
	}{
		{
			name: "successful get or create category",
			payload: &videolearning.GetOrCreateCategoryPayload{
				SessionToken: "valid_token",
				Name:         "New Category",
			},
			setupMocks: func(categoryRepo *mocks.MockVideoCategoryRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				category := videolearningdb.VideoCategory{
					ID:   2,
					Name: "New Category",
				}
				categoryRepo.On("GetOrCreateCategory", mock.Anything, "New Category").Return(category, nil)
			},
			expectedResult: &videolearning.VideoCategory{
				ID:   2,
				Name: "New Category",
			},
			expectedError: nil,
		},
		{
			name: "database error",
			payload: &videolearning.GetOrCreateCategoryPayload{
				SessionToken: "valid_token",
				Name:         "New Category",
			},
			setupMocks: func(categoryRepo *mocks.MockVideoCategoryRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				categoryRepo.On("GetOrCreateCategory", mock.Anything, "New Category").Return(videolearningdb.VideoCategory{}, errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  videolearning.ServiceUnavailable("failed to get or create category"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			categoryRepo := &mocks.MockVideoCategoryRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(categoryRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, nil, categoryRepo, nil, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.GetOrCreateCategory(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.Name, result.Name)
			}

			// Verify mocks
			categoryRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test GetOrCreateTag endpoint
func TestGetOrCreateTag(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetOrCreateTagPayload
		setupMocks     func(*mocks.MockVideoTagRepository, *mocks.MockProfilesServiceRepository)
		expectedResult *videolearning.VideoTag
		expectedError  error
	}{
		{
			name: "successful get or create tag",
			payload: &videolearning.GetOrCreateTagPayload{
				SessionToken: "valid_token",
				Name:         "new-tag",
			},
			setupMocks: func(tagRepo *mocks.MockVideoTagRepository, profilesRepo *mocks.MockProfilesServiceRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				tag := videolearningdb.VideoTag{
					ID:   2,
					Name: "new-tag",
				}
				tagRepo.On("GetOrCreateTag", mock.Anything, "new-tag").Return(tag, nil)
			},
			expectedResult: &videolearning.VideoTag{
				ID:   2,
				Name: "new-tag",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			tagRepo := &mocks.MockVideoTagRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			tt.setupMocks(tagRepo, profilesRepo)

			// Create service
			service := setupTestService(nil, nil, nil, tagRepo, nil, profilesRepo, nil, nil)

			// Call method
			result, err := service.GetOrCreateTag(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResult.ID, result.ID)
				assert.Equal(t, tt.expectedResult.Name, result.Name)
			}

			// Verify mocks
			tagRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}

// Test GetOwnVideos endpoint (teacher only)
func TestGetOwnVideos(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetOwnVideosPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository, *mocks.MockStorageRepository)
		expectedResult []*videolearning.OwnVideo
		expectedError  error
	}{
		{
			name: "successful get own videos by teacher",
			payload: &videolearning.GetOwnVideosPayload{
				SessionToken: "teacher_token",
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "teacher_token",
				}).Return(createTestCompleteProfile("teacher"), nil)

				getParams := videolearningdb.GetVideosByUserParams{
					UserID: 1,
					Limit:  10,
					Offset: 0,
				}
				userVideos := []videolearningdb.Video{
					{
						ID:           1,
						Title:        "My Video",
						UserID:       1,
						Views:        100,
						Likes:        50,
						ThumbObjName: pgtype.Text{String: "my_thumb.jpg", Valid: true},
						CategoryID:   1,
						CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
						UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
					},
				}
				videoRepo.On("GetVideosByUser", mock.Anything, getParams).Return(userVideos, nil)

				// Mock cache operations for views and likes
				cacheRepo.On("GetInt", mock.Anything, "video:views:1").Return(int64(10), nil)
				cacheRepo.On("GetInt", mock.Anything, "video:likes:1").Return(int64(5), nil)

				// Mock presigned URL generation for thumbnail
				cacheRepo.On("Get", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:my_thumb.jpg").Return("", errors.New("not found"))
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", "my_thumb.jpg", mock.Anything).Return("https://example.com/my_thumb.jpg", nil)
				cacheRepo.On("Set", mock.Anything, "presigned_url:video-learning-thumbnails-confirmed:my_thumb.jpg", "https://example.com/my_thumb.jpg", mock.Anything).Return(nil)
			},
			expectedResult: []*videolearning.OwnVideo{
				{
					ID:           1,
					Title:        "My Video",
					Views:        110, // 100 + 10 cached
					Likes:        55,  // 50 + 5 cached
					ThumbnailURL: "https://example.com/my_thumb.jpg",
				},
			},
			expectedError: nil,
		},
		{
			name: "permission denied - not a teacher",
			payload: &videolearning.GetOwnVideosPayload{
				SessionToken: "student_token",
				Page:         1,
				PageSize:     10,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "student_token",
				}).Return(createTestCompleteProfile("student"), nil)
			},
			expectedResult: nil,
			expectedError:  videolearning.PermissionDenied("only teachers can view their uploaded videos"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, profilesRepo, cacheRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, nil, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.GetOwnVideos(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expectedResult), len(result))
				if len(result) > 0 {
					assert.Equal(t, tt.expectedResult[0].ID, result[0].ID)
					assert.Equal(t, tt.expectedResult[0].Title, result[0].Title)
					assert.Equal(t, tt.expectedResult[0].Views, result[0].Views)
					assert.Equal(t, tt.expectedResult[0].Likes, result[0].Likes)
				}
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
			cacheRepo.AssertExpectations(t)
			storageRepo.AssertExpectations(t)
		})
	}
}

// Test GetRecommendations endpoint
func TestGetRecommendations(t *testing.T) {
	tests := []struct {
		name           string
		payload        *videolearning.GetRecommendationsPayload
		setupMocks     func(*mocks.MockVideoRepository, *mocks.MockUserCategoryLikeRepository, *mocks.MockProfilesServiceRepository, *mocks.MockCacheRepository, *mocks.MockStorageRepository)
		expectedResult *videolearning.VideoList
		expectedError  error
	}{
		{
			name: "successful get recommendations without user preferences",
			payload: &videolearning.GetRecommendationsPayload{
				SessionToken: "valid_token",
				Amount:       1, // Set to 1 so we have more videos than requested amount
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, userCatRepo *mocks.MockUserCategoryLikeRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				// Mock recent videos
				interval := pgtype.Interval{Days: 7, Valid: true}
				recentVideos := []videolearningdb.Video{
					{
						ID:           1,
						Title:        "Recent Video 1",
						UserID:       1,
						Views:        100,
						Likes:        50,
						ThumbObjName: pgtype.Text{String: "recent_thumb1.jpg", Valid: true},
						CategoryID:   1,
						CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
						UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
					},
					{
						ID:           2,
						Title:        "Recent Video 2",
						UserID:       2,
						Views:        80,
						Likes:        40,
						ThumbObjName: pgtype.Text{String: "recent_thumb2.jpg", Valid: true},
						CategoryID:   2,
						CreatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
						UpdatedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
					},
				}
				videoRepo.On("GetRecentVideos", mock.Anything, interval).Return(recentVideos, nil)

				// Mock user category likes (empty - no preferences)
				userCatRepo.On("GetUserCategoryLikes", mock.Anything, int64(1)).Return([]videolearningdb.GetUserCategoryLikesRow{}, nil)

				// Mock public profiles for video authors (only need one since amount=1)
				profilesRepo.On("GetPublicProfileByID", mock.Anything, mock.MatchedBy(func(payload *profiles.GetPublicProfileByIDPayload) bool {
					return payload.UserID == 1 || payload.UserID == 2 // Accept either user ID since randomSelectVideos picks randomly
				})).Return(createTestPublicProfile(), nil)

				// Cache expectations for video views and likes (only need one since amount=1)
				cacheRepo.On("GetInt", mock.Anything, mock.MatchedBy(func(key string) bool {
					return key == "video:views:1" || key == "video:views:2"
				})).Return(int64(5), nil)
				cacheRepo.On("GetInt", mock.Anything, mock.MatchedBy(func(key string) bool {
					return key == "video:likes:1" || key == "video:likes:2"
				})).Return(int64(10), nil)

				// Cache expectations for presigned URLs (only need one since amount=1)
				cacheRepo.On("Get", mock.Anything, mock.MatchedBy(func(key string) bool {
					return key == "presigned_url:video-learning-thumbnails-confirmed:recent_thumb1.jpg" ||
						key == "presigned_url:video-learning-thumbnails-confirmed:recent_thumb2.jpg"
				})).Return("", errors.New("not cached"))

				// Storage repository presigned URL generation (only need one since amount=1)
				storageRepo.On("GeneratePresignedURL", mock.Anything, "video-learning-thumbnails-confirmed", mock.MatchedBy(func(objName string) bool {
					return objName == "recent_thumb1.jpg" || objName == "recent_thumb2.jpg"
				}), mock.Anything).Return("https://example.com/thumb.jpg", nil)

				// Cache expectations for storing presigned URLs (only need one since amount=1)
				cacheRepo.On("Set", mock.Anything, mock.MatchedBy(func(key string) bool {
					return key == "presigned_url:video-learning-thumbnails-confirmed:recent_thumb1.jpg" ||
						key == "presigned_url:video-learning-thumbnails-confirmed:recent_thumb2.jpg"
				}), "https://example.com/thumb.jpg", mock.Anything).Return(nil)
			},
			expectedResult: &videolearning.VideoList{
				Videos: []*videolearning.Video{}, // Results will be randomized, so we check length instead
			},
			expectedError: nil,
		},
		{
			name: "invalid session",
			payload: &videolearning.GetRecommendationsPayload{
				SessionToken: "invalid_token",
				Amount:       5,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, userCatRepo *mocks.MockUserCategoryLikeRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "invalid_token",
				}).Return((*profiles.CompleteProfileResponse)(nil), errors.New("invalid session"))
			},
			expectedResult: nil,
			expectedError:  videolearning.InvalidSession("invalid session"),
		},
		{
			name: "database error getting recent videos",
			payload: &videolearning.GetRecommendationsPayload{
				SessionToken: "valid_token",
				Amount:       5,
			},
			setupMocks: func(videoRepo *mocks.MockVideoRepository, userCatRepo *mocks.MockUserCategoryLikeRepository, profilesRepo *mocks.MockProfilesServiceRepository, cacheRepo *mocks.MockCacheRepository, storageRepo *mocks.MockStorageRepository) {
				profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
					SessionToken: "valid_token",
				}).Return(createTestCompleteProfile("student"), nil)

				interval := pgtype.Interval{Days: 7, Valid: true}
				videoRepo.On("GetRecentVideos", mock.Anything, interval).Return([]videolearningdb.Video(nil), errors.New("database error"))
			},
			expectedResult: nil,
			expectedError:  videolearning.ServiceUnavailable("failed to get recommendations"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			videoRepo := &mocks.MockVideoRepository{}
			userCatRepo := &mocks.MockUserCategoryLikeRepository{}
			profilesRepo := &mocks.MockProfilesServiceRepository{}
			cacheRepo := &mocks.MockCacheRepository{}
			storageRepo := &mocks.MockStorageRepository{}
			tt.setupMocks(videoRepo, userCatRepo, profilesRepo, cacheRepo, storageRepo)

			// Create service
			service := setupTestService(videoRepo, nil, nil, nil, userCatRepo, profilesRepo, cacheRepo, storageRepo)

			// Call method
			result, err := service.GetRecommendations(context.Background(), tt.payload)

			// Assertions
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotNil(t, result.Videos)
				// For recommendations, we expect some videos but the exact count may vary due to randomization
				assert.LessOrEqual(t, len(result.Videos), tt.payload.Amount)
			}

			// Verify mocks
			videoRepo.AssertExpectations(t)
			userCatRepo.AssertExpectations(t)
			profilesRepo.AssertExpectations(t)
		})
	}
}
