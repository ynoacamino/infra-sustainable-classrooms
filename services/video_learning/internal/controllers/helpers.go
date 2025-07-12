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

// convertVideoToAPIVideo converts database video row to API video struct
func (s *videolearningsrvc) convertVideoToAPIVideo(ctx context.Context, video interface{}) (*videolearning.Video, error) {
	var id int64
	var title string
	var userID int64
	var views int32
	var likes int32
	var thumbObjName pgtype.Text

	// Handle different video row types
	switch v := video.(type) {
	case videolearningdb.SearchVideosRow:
		id = v.ID
		title = v.Title
		userID = v.UserID
		views = v.Views
		likes = v.Likes
		thumbObjName = v.ThumbObjName
	case videolearningdb.GetVideosByCategoryRow:
		id = v.ID
		title = v.Title
		userID = v.UserID
		views = v.Views
		likes = v.Likes
		thumbObjName = v.ThumbObjName
	case videolearningdb.GetSimilarVideosRow:
		id = v.ID
		title = v.Title
		userID = v.UserID
		views = v.Views
		likes = v.Likes
		thumbObjName = v.ThumbObjName
	case videolearningdb.GetRecentVideosRow:
		id = v.ID
		title = v.Title
		userID = v.UserID
		views = v.Views
		likes = v.Likes
		thumbObjName = v.ThumbObjName
	case videolearningdb.GetVideosByUserRow:
		id = v.ID
		title = v.Title
		userID = v.UserID
		views = v.Views
		likes = v.Likes
		thumbObjName = v.ThumbObjName
	default:
		return nil, fmt.Errorf("unsupported video type")
	}

	// Get cached views and likes
	cachedViews, _ := s.getCachedVideoViews(ctx, id)
	cachedLikes, _ := s.getCachedVideoLikes(ctx, id)

	totalViews := int(views) + cachedViews
	totalLikes := int(likes) + cachedLikes

	// Generate thumbnail URL
	thumbnailURL := ""
	if thumbObjName.Valid && thumbObjName.String != "" {
		presignedURL, err := s.storageRepo.GeneratePresignedURL(ctx, "video-learning-thumbnails-confirmed", thumbObjName.String, 24*time.Hour)
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
	result, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		return 0, nil // Return 0 if not found or error
	}

	views, err := strconv.Atoi(result)
	if err != nil {
		return 0, nil
	}

	return views, nil
}

// getCachedVideoLikes gets cached like count for a video
func (s *videolearningsrvc) getCachedVideoLikes(ctx context.Context, videoID int64) (int, error) {
	key := fmt.Sprintf("video:likes:%d", videoID)
	result, err := s.cacheRepo.Get(ctx, key)
	if err != nil {
		return 0, nil // Return 0 if not found or error
	}

	likes, err := strconv.Atoi(result)
	if err != nil {
		return 0, nil
	}

	return likes, nil
}

// incrementCachedVideoViews increments cached view count
func (s *videolearningsrvc) incrementCachedVideoViews(ctx context.Context, videoID int64) error {
	key := fmt.Sprintf("video:views:%d", videoID)

	// Try to increment existing value
	currentViews, err := s.getCachedVideoViews(ctx, videoID)
	if err != nil {
		currentViews = 0
	}

	newViews := currentViews + 1
	return s.cacheRepo.Set(ctx, key, strconv.Itoa(newViews), 24*time.Hour)
}

// incrementCachedVideoLikes increments cached like count
func (s *videolearningsrvc) incrementCachedVideoLikes(ctx context.Context, videoID int64, increment int) error {
	key := fmt.Sprintf("video:likes:%d", videoID)

	currentLikes, err := s.getCachedVideoLikes(ctx, videoID)
	if err != nil {
		currentLikes = 0
	}

	newLikes := currentLikes + increment
	return s.cacheRepo.Set(ctx, key, strconv.Itoa(newLikes), 24*time.Hour)
}

// incrementCachedUserCategoryLike increments cached user category like count
func (s *videolearningsrvc) incrementCachedUserCategoryLike(ctx context.Context, userID, categoryID int64, increment int) error {
	key := fmt.Sprintf("user:category:likes:%d:%d", userID, categoryID)

	// Get current value
	result, err := s.cacheRepo.Get(ctx, key)
	current := 0
	if err == nil {
		current, _ = strconv.Atoi(result)
	}

	newValue := current + increment
	return s.cacheRepo.Set(ctx, key, strconv.Itoa(newValue), 24*time.Hour)
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
	for i := 0; i < amount; i++ {
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
