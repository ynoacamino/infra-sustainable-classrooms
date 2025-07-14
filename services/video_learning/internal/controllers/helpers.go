package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
)

func (s *videolearningsrvc) validateSessionAndGetUserID(ctx context.Context, sessionToken string) (int64, error) {
	profile, err := s.profileServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: sessionToken,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to validate session: %w", err)
	}

	return profile.UserID, nil
}

// validateSessionAndGetUserProfile validates session and returns complete user profile
func (s *videolearningsrvc) validateSessionAndGetUserProfile(ctx context.Context, sessionToken string) (*profiles.CompleteProfileResponse, error) {
	profile, err := s.profileServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: sessionToken,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to validate session: %w", err)
	}

	return profile, nil
}

// validateTeacherRole validates that the user is a teacher
func (s *videolearningsrvc) validateTeacherRole(ctx context.Context, sessionToken string) (*profiles.CompleteProfileResponse, error) {
	profile, err := s.validateSessionAndGetUserProfile(ctx, sessionToken)
	if err != nil {
		return nil, err
	}

	if profile.Role != "teacher" {
		return nil, fmt.Errorf("access denied: only teachers can perform this action")
	}

	return profile, nil
}

func (s *videolearningsrvc) convertTagsToAPITags(ctx context.Context, tags []videolearningdb.VideoTag) []int64 {
	apiTags := make([]int64, len(tags))
	for i, tag := range tags {
		apiTags[i] = tag.ID
	}
	return apiTags
}

// convertVideoToAPIVideo converts database video row to API video struct
func (s *videolearningsrvc) convertVideoToAPIVideo(ctx context.Context, video interface{}) (*videolearning.Video, error) {
	var id int64
	var title string
	var userID int64
	var views int32
	var likes int32
	var thumbObjName pgtype.Text

	// Handle different video row types
	v := video.(videolearningdb.Video)
	id = v.ID
	title = v.Title
	userID = v.UserID
	views = v.Views
	likes = v.Likes
	thumbObjName = v.ThumbObjName

	// Get cached views and likes
	cachedViews, _ := s.getCachedVideoViews(ctx, id)
	cachedLikes, _ := s.getCachedVideoLikes(ctx, id)

	totalViews := int(views) + cachedViews
	totalLikes := int(likes) + cachedLikes

	// Generate thumbnail URL with caching
	thumbnailURL := ""
	if thumbObjName.Valid && thumbObjName.String != "" {
		presignedURL, err := s.getOrGeneratePresignedURL(ctx, "video-learning-thumbnails-confirmed", thumbObjName.String, 24*time.Hour)
		if err == nil {
			thumbnailURL = presignedURL
		}
	}

	// Get author name from profile service
	author := fmt.Sprintf("User_%d", userID) // Default fallback
	if publicProfile, err := s.profileServiceRepo.GetPublicProfileByID(ctx, &profiles.GetPublicProfileByIDPayload{
		UserID: userID,
	}); err == nil {
		author = fmt.Sprintf("%s %s", publicProfile.FirstName, publicProfile.LastName)
	}

	return &videolearning.Video{
		ID:           id,
		Title:        title,
		Author:       author,
		Views:        totalViews,
		Likes:        totalLikes,
		ThumbnailURL: thumbnailURL,
	}, nil
}

// getCachedVideoViews gets cached view count for a video
func (s *videolearningsrvc) getCachedVideoViews(ctx context.Context, videoID int64) (int, error) {
	key := fmt.Sprintf("video:views:%d", videoID)

	// Try to get as integer first
	result, err := s.cacheRepo.GetInt(ctx, key)
	if err != nil {
		return 0, nil // Return 0 if not found or error
	}

	return int(result), nil
}

// getCachedVideoLikes gets cached like count for a video
func (s *videolearningsrvc) getCachedVideoLikes(ctx context.Context, videoID int64) (int, error) {
	key := fmt.Sprintf("video:likes:%d", videoID)

	// Try to get as integer first
	result, err := s.cacheRepo.GetInt(ctx, key)
	if err != nil {
		return 0, nil // Return 0 if not found or error
	}

	return int(result), nil
}

// incrementCachedVideoViews increments cached view count atomically
func (s *videolearningsrvc) incrementCachedVideoViews(ctx context.Context, videoID int64) error {
	key := fmt.Sprintf("video:views:%d", videoID)

	// Use atomic increment operation
	_, err := s.cacheRepo.IncrBy(ctx, key, 1)
	if err != nil {
		// If increment fails (key doesn't exist), set it to 1
		return s.cacheRepo.Set(ctx, key, "1", 24*time.Hour)
	}

	return nil
}

// incrementCachedVideoLikes increments cached like count atomically
func (s *videolearningsrvc) incrementCachedVideoLikes(ctx context.Context, videoID int64, increment int) error {
	key := fmt.Sprintf("video:likes:%d", videoID)

	// Use atomic increment operation
	_, err := s.cacheRepo.IncrBy(ctx, key, int64(increment))
	if err != nil {
		// If increment fails (key doesn't exist), set it to the increment value
		return s.cacheRepo.Set(ctx, key, strconv.Itoa(increment), 24*time.Hour)
	}

	return nil
}

// incrementCachedUserCategoryLike increments cached user category like count atomically
func (s *videolearningsrvc) incrementCachedUserCategoryLike(ctx context.Context, userID, categoryID int64, increment int) error {
	key := fmt.Sprintf("user:category:likes:%d:%d", userID, categoryID)

	// Use atomic increment operation
	_, err := s.cacheRepo.IncrBy(ctx, key, int64(increment))
	if err != nil {
		// If increment fails (key doesn't exist), set it to the increment value
		return s.cacheRepo.Set(ctx, key, strconv.Itoa(increment), 24*time.Hour)
	}

	return nil
}

// randomSelectVideos randomly selects up to 'amount' videos from the slice
func randomSelectVideos[T any](videos []T, amount int) []T {
	if amount >= len(videos) {
		return videos
	}

	if amount <= 0 {
		return []T{}
	}

	// Create a copy of indices
	indices := make([]int, len(videos))
	for i := range indices {
		indices[i] = i
	}

	// Shuffle indices
	rand.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	// Select first 'amount' videos
	result := make([]T, amount)
	for i := range amount {
		result[i] = videos[indices[i]]
	}

	return result
}

// generateObjectName generates a unique object name for file uploads
func (s *videolearningsrvc) generateObjectName(filename string, userID int64) string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%d_%d_%s", userID, timestamp, sanitizeFilename(filename))
}

// sanitizeFilename removes unsafe characters from filename
func sanitizeFilename(filename string) string {
	result := make([]rune, 0, len(filename))
	for _, r := range filename {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '-' || r == '_' {
			result = append(result, r)
		} else {
			result = append(result, '_')
		}
	}
	return string(result)
}

// getCachedPresignedURL gets a cached presigned URL from Redis
func (s *videolearningsrvc) getCachedPresignedURL(ctx context.Context, bucket, objectName string) (string, error) {
	key := fmt.Sprintf("presigned_url:%s:%s", bucket, objectName)
	return s.cacheRepo.Get(ctx, key)
}

// setCachedPresignedURL stores a presigned URL in Redis with expiration
func (s *videolearningsrvc) setCachedPresignedURL(ctx context.Context, bucket, objectName, url string, expiration time.Duration) error {
	key := fmt.Sprintf("presigned_url:%s:%s", bucket, objectName)
	// Set cache expiration to 90% of URL expiration to ensure we refresh before URL expires
	cacheExpiration := time.Duration(float64(expiration) * 0.9)
	return s.cacheRepo.Set(ctx, key, url, cacheExpiration)
}

// getOrGeneratePresignedURL gets a cached presigned URL or generates a new one
func (s *videolearningsrvc) getOrGeneratePresignedURL(ctx context.Context, bucket, objectName string, expiration time.Duration) (string, error) {
	// Try to get from cache first
	if cachedURL, err := s.getCachedPresignedURL(ctx, bucket, objectName); err == nil && cachedURL != "" {
		return cachedURL, nil
	}

	// Generate new presigned URL
	url, err := s.storageRepo.GeneratePresignedURL(ctx, bucket, objectName, expiration)
	if err != nil {
		return "", err
	}

	// Cache the URL
	s.setCachedPresignedURL(ctx, bucket, objectName, url, expiration)

	return url, nil
}

// clearCachedPresignedURL removes a cached presigned URL from Redis
func (s *videolearningsrvc) clearCachedPresignedURL(ctx context.Context, bucket, objectName string) error {
	key := fmt.Sprintf("presigned_url:%s:%s", bucket, objectName)
	return s.cacheRepo.Delete(ctx, key)
}
