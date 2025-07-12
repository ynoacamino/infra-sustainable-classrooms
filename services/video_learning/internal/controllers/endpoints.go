package controllers

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
	videolearning "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/video_learning"
)

// SearchVideos searches videos by query string and category
func (s *videolearningsrvc) SearchVideos(ctx context.Context, p *videolearning.SearchVideosPayload) (res *videolearning.VideoList, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Prepare search parameters
	searchParams := videolearningdb.SearchVideosParams{
		Column1: pgtype.Text{String: p.Query, Valid: true},
		Column2: 0, // Default category (null case)
		Limit:   int32(p.PageSize),
		Offset:  int32((p.Page - 1) * p.PageSize),
	}

	if p.CategoryID != nil {
		searchParams.Column2 = *p.CategoryID
	}

	// Search videos
	videos, err := s.videoRepo.SearchVideos(ctx, searchParams)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to search videos")
	}

	// Convert to API format
	var apiVideos []*videolearning.Video
	for _, video := range videos {
		apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
		if err != nil {
			continue // Skip videos that fail conversion
		}
		apiVideos = append(apiVideos, apiVideo)
	}

	return &videolearning.VideoList{
		Videos: apiVideos,
	}, nil
}

// GetRecommendations returns recommended videos for user
func (s *videolearningsrvc) GetRecommendations(ctx context.Context, p *videolearning.GetRecommendationsPayload) (res *videolearning.VideoList, err error) {
	// Validate session
	userID, err := s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get all recent videos from the last 7 days
	interval := pgtype.Interval{
		Days:  7,
		Valid: true,
	}
	recentVideos, err := s.videoRepo.GetRecentVideos(ctx, interval)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get recommendations")
	}

	// If we have fewer or equal videos than requested, return all of them
	if len(recentVideos) <= p.Ammount {
		var apiVideos []*videolearning.Video
		for _, video := range recentVideos {
			apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
			if err != nil {
				continue
			}
			apiVideos = append(apiVideos, apiVideo)
		}

		return &videolearning.VideoList{
			Videos: apiVideos,
		}, nil
	}

	// Get user's category preferences
	categoryLikes, err := s.userCategoryLikeRepo.GetUserCategoryLikes(ctx, userID)
	if err != nil || len(categoryLikes) == 0 {
		// If no preferences found, randomly select from recent videos
		selectedVideos := randomSelectVideos(recentVideos, p.Ammount)

		var apiVideos []*videolearning.Video
		for _, video := range selectedVideos {
			apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
			if err != nil {
				continue
			}
			apiVideos = append(apiVideos, apiVideo)
		}

		return &videolearning.VideoList{
			Videos: apiVideos,
		}, nil
	}

	// Calculate distribution based on likes
	totalLikes := int32(0)
	minNonZeroLikes := int32(1000000) // Large number

	for _, catLike := range categoryLikes {
		if catLike.Likes.Valid && catLike.Likes.Int32 > 0 {
			totalLikes += catLike.Likes.Int32
			if catLike.Likes.Int32 < minNonZeroLikes {
				minNonZeroLikes = catLike.Likes.Int32
			}
		}
	}

	// If no likes found, set minimum to 1
	if minNonZeroLikes == 1000000 {
		minNonZeroLikes = 1
	}

	// Add minimum likes to null/zero categories and recalculate total
	for i := range categoryLikes {
		if !categoryLikes[i].Likes.Valid || categoryLikes[i].Likes.Int32 == 0 {
			categoryLikes[i].Likes = pgtype.Int4{Int32: minNonZeroLikes, Valid: true}
			totalLikes += minNonZeroLikes
		}
	}

	// Group recent videos by category
	videosByCategory := make(map[int64][]videolearningdb.GetRecentVideosRow)
	for _, video := range recentVideos {
		videosByCategory[video.CategoryID] = append(videosByCategory[video.CategoryID], video)
	}

	// Create a map from category names to category IDs from user likes
	categoryNameToID := make(map[string]int64)
	for _, catLike := range categoryLikes {
		// Find the category ID by looking through videos with matching category name
		for _, video := range recentVideos {
			if video.CategoryName == catLike.Name {
				categoryNameToID[catLike.Name] = video.CategoryID
				break
			}
		}
	}

	// Distribute amount across categories proportionally
	var allRecommendedVideos []*videolearning.Video
	remaining := p.Ammount

	for _, catLike := range categoryLikes {
		if remaining <= 0 {
			break
		}

		categoryID, exists := categoryNameToID[catLike.Name]
		if !exists {
			continue
		}

		// Check if we have videos in this category
		categoryVideos, hasVideos := videosByCategory[categoryID]
		if !hasVideos || len(categoryVideos) == 0 {
			continue
		}

		// Calculate how many videos to get from this category
		proportion := float64(catLike.Likes.Int32) / float64(totalLikes)
		videoCount := int(float64(p.Ammount) * proportion)
		if videoCount > remaining {
			videoCount = remaining
		}
		if videoCount == 0 {
			videoCount = 1 // At least 1 video per category
		}

		// Randomly select from category videos
		selectedFromCategory := randomSelectVideos(categoryVideos, videoCount)

		for _, video := range selectedFromCategory {
			if remaining <= 0 {
				break
			}
			apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
			if err != nil {
				continue
			}
			allRecommendedVideos = append(allRecommendedVideos, apiVideo)
			remaining--
		}
	}

	// If we still have remaining slots and there are unassigned videos, fill randomly
	if remaining > 0 {
		var allUnassignedVideos []videolearningdb.GetRecentVideosRow
		usedVideoIDs := make(map[int64]bool)

		// Mark used videos
		for _, video := range allRecommendedVideos {
			usedVideoIDs[video.ID] = true
		}

		// Collect unused videos
		for _, video := range recentVideos {
			if !usedVideoIDs[video.ID] {
				allUnassignedVideos = append(allUnassignedVideos, video)
			}
		}

		// Randomly select from unused videos
		selectedRemainingVideos := randomSelectVideos(allUnassignedVideos, remaining)
		for _, video := range selectedRemainingVideos {
			apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
			if err != nil {
				continue
			}
			allRecommendedVideos = append(allRecommendedVideos, apiVideo)
		}
	}

	return &videolearning.VideoList{
		Videos: allRecommendedVideos,
	}, nil
}

// GetVideo returns complete video information including presigned URL
func (s *videolearningsrvc) GetVideo(ctx context.Context, p *videolearning.GetVideoPayload) (res *videolearning.VideoDetails, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get video details
	video, err := s.videoRepo.GetVideoByID(ctx, p.VideoID)
	if err != nil {
		return nil, videolearning.VideoNotFound("video not found")
	}

	// Increment view count in cache
	s.incrementCachedVideoViews(ctx, p.VideoID)

	// Get cached metrics
	cachedViews, _ := s.getCachedVideoViews(ctx, p.VideoID)
	cachedLikes, _ := s.getCachedVideoLikes(ctx, p.VideoID)

	totalViews := int(video.Views) + cachedViews
	totalLikes := int(video.Likes) + cachedLikes

	// Generate presigned URLs with caching
	var videoURL, thumbnailURL string

	if video.VideoObjName.Valid && video.VideoObjName.String != "" {
		videoURL, _ = s.getOrGeneratePresignedURL(ctx, "video-learning-videos-confirmed", video.VideoObjName.String, 4*time.Hour)
	}

	if video.ThumbObjName.Valid && video.ThumbObjName.String != "" {
		thumbnailURL, _ = s.getOrGeneratePresignedURL(ctx, "video-learning-thumbnails-confirmed", video.ThumbObjName.String, 24*time.Hour)
	}

	// Get author name from profile service
	author := fmt.Sprintf("User_%d", video.UserID) // Default fallback
	if publicProfile, err := s.profileServiceRepo.GetPublicProfileByID(ctx, &profiles.GetPublicProfileByIDPayload{
		UserID: video.UserID,
	}); err == nil {
		author = fmt.Sprintf("%s %s", publicProfile.FirstName, publicProfile.LastName)
	}

	tags, err := s.videoTagRepo.GetTagsByVideoID(ctx, p.VideoID)

	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get video tags")
	}

	apiTags := s.convertTagsToAPITags(ctx, tags)

	return &videolearning.VideoDetails{
		ID:           video.ID,
		Title:        video.Title,
		Description:  video.Description.String,
		Author:       author,
		Views:        totalViews,
		Likes:        totalLikes,
		VideoURL:     videoURL,
		ThumbnailURL: thumbnailURL,
		UploadDate:   video.CreatedAt.Time.UnixMilli(),
		Category:     video.CategoryName,
		Tags:         apiTags,
	}, nil
}

// GetSimilarVideos returns videos similar to the given video
func (s *videolearningsrvc) GetSimilarVideos(ctx context.Context, p *videolearning.GetSimilarVideosPayload) (res *videolearning.VideoList, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get similar videos
	similarVideos, err := s.videoRepo.GetSimilarVideos(ctx, p.VideoID)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get similar videos")
	}

	// Randomly select videos
	selectedVideos := randomSelectVideos(similarVideos, p.Amount)

	var apiVideos []*videolearning.Video
	for _, video := range selectedVideos {
		apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
		if err != nil {
			continue
		}
		apiVideos = append(apiVideos, apiVideo)
	}

	return &videolearning.VideoList{
		Videos: apiVideos,
	}, nil
}

// GetVideosByCategory returns videos filtered by category
func (s *videolearningsrvc) GetVideosByCategory(ctx context.Context, p *videolearning.GetVideosByCategoryPayload) (res *videolearning.VideoList, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get videos by category
	categoryVideos, err := s.videoRepo.GetVideosByCategory(ctx, p.CategoryID)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get videos by category")
	}

	// Randomly select videos
	selectedVideos := randomSelectVideos(categoryVideos, p.Amount)

	var apiVideos []*videolearning.Video
	for _, video := range selectedVideos {
		apiVideo, err := s.convertVideoToAPIVideo(ctx, video)
		if err != nil {
			continue
		}
		apiVideos = append(apiVideos, apiVideo)
	}

	return &videolearning.VideoList{
		Videos: apiVideos,
	}, nil
}

// GetComments returns paginated comments for a video
func (s *videolearningsrvc) GetComments(ctx context.Context, p *videolearning.GetCommentsPayload) (res *videolearning.CommentList, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get comments
	comments, err := s.videoCommentRepo.GetCommentsForVideo(ctx, videolearningdb.GetCommentsForVideoParams{
		VideoID: p.VideoID,
		Limit:   int32(p.PageSize),
		Offset:  int32((p.Page - 1) * p.PageSize),
	})
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get comments")
	}
	// Convert to API format
	var apiComments []*videolearning.Comment
	for _, comment := range comments {
		// Get author name from profile service
		author := fmt.Sprintf("User_%d", comment.UserID) // Default fallback
		if publicProfile, err := s.profileServiceRepo.GetPublicProfileByID(ctx, &profiles.GetPublicProfileByIDPayload{
			UserID: comment.UserID,
		}); err == nil {
			author = fmt.Sprintf("%s %s", publicProfile.FirstName, publicProfile.LastName)
		}

		apiComments = append(apiComments, &videolearning.Comment{
			ID:     comment.ID,
			Author: author,
			Date:   comment.CreatedAt.Time.UnixMilli(),
			Title:  comment.Title,
			Body:   comment.Content,
		})
	}

	return &videolearning.CommentList{
		Comments: apiComments,
	}, nil
}

// CreateComment creates a new comment on a video
func (s *videolearningsrvc) CreateComment(ctx context.Context, p *videolearning.CreateCommentPayload) (res *videolearning.SimpleResponse, err error) {
	// Validate session
	userID, err := s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Create comment
	comment, err := s.videoCommentRepo.CreateComment(ctx, videolearningdb.CreateCommentParams{
		VideoID: p.VideoID,
		UserID:  userID,
		Title:   p.Title,
		Content: p.Body,
	})
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to create comment")
	}

	return &videolearning.SimpleResponse{
		Success: true,
		Message: "Comment created successfully",
		ID:      &comment.ID,
	}, nil
}

// DeleteComment deletes a comment by ID
func (s *videolearningsrvc) DeleteComment(ctx context.Context, p *videolearning.DeleteCommentPayload) (res *videolearning.SimpleResponse, err error) {
	// Validate session
	userID, err := s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Delete comment (only if user owns it)
	err = s.videoCommentRepo.DeleteComment(ctx, videolearningdb.DeleteCommentParams{
		ID:     p.CommentID,
		UserID: userID,
	})
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to delete comment")
	}

	return &videolearning.SimpleResponse{
		Success: true,
		Message: "Comment deleted successfully",
	}, nil
}

// GetOwnVideos returns authenticated user's uploaded videos
func (s *videolearningsrvc) GetOwnVideos(ctx context.Context, p *videolearning.GetOwnVideosPayload) (res []*videolearning.OwnVideo, err error) {
	// Validate session and check teacher role
	profile, err := s.validateTeacherRole(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.PermissionDenied("only teachers can view their uploaded videos")
	}

	// Get user's videos
	videos, err := s.videoRepo.GetVideosByUser(ctx, videolearningdb.GetVideosByUserParams{
		UserID: profile.UserID,
		Limit:  int32(p.PageSize),
		Offset: int32((p.Page - 1) * p.PageSize),
	})
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get user videos")
	}

	// Convert to API format
	var ownVideos []*videolearning.OwnVideo
	for _, video := range videos {
		// Get cached metrics
		cachedViews, _ := s.getCachedVideoViews(ctx, video.ID)
		cachedLikes, _ := s.getCachedVideoLikes(ctx, video.ID)

		totalViews := int(video.Views) + cachedViews
		totalLikes := int(video.Likes) + cachedLikes

		// Generate thumbnail URL with caching
		thumbnailURL := ""
		if video.ThumbObjName.Valid && video.ThumbObjName.String != "" {
			presignedURL, err := s.getOrGeneratePresignedURL(ctx, "video-learning-thumbnails-confirmed", video.ThumbObjName.String, 24*time.Hour)
			if err == nil {
				thumbnailURL = presignedURL
			}
		}

		ownVideos = append(ownVideos, &videolearning.OwnVideo{
			ID:           video.ID,
			Title:        video.Title,
			Views:        totalViews,
			Likes:        totalLikes,
			ThumbnailURL: thumbnailURL,
			UploadDate:   video.CreatedAt.Time.UnixMilli(),
		})
	}

	return ownVideos, nil
}

// InitialUpload uploads video file and returns object name
func (s *videolearningsrvc) InitialUpload(ctx context.Context, p *videolearning.InitialUploadPayload) (res *videolearning.UploadResponse, err error) {
	// Validate session and check teacher role
	profile, err := s.validateTeacherRole(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.PermissionDenied("only teachers can upload videos")
	}

	// Generate object name
	objectName := s.generateObjectName(p.Filename, profile.UserID)

	// Upload to staging bucket
	reader := bytes.NewReader(p.File)
	err = s.storageRepo.UploadFile(ctx, "video-learning-videos-staging", objectName, reader, int64(len(p.File)), "video/mp4")
	if err != nil {
		return nil, videolearning.UploadFailed("failed to upload video")
	}

	return &videolearning.UploadResponse{
		ObjectName: objectName,
	}, nil
}

// CompleteUpload completes video upload with metadata
func (s *videolearningsrvc) CompleteUpload(ctx context.Context, p *videolearning.CompleteUploadPayload) (res *videolearning.SimpleResponse, err error) {
	// Validate session and check teacher role
	profile, err := s.validateTeacherRole(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.PermissionDenied("only teachers can upload videos")
	}

	// Move video
	err = s.storageRepo.CopyFile(ctx, "video-learning-videos-staging", p.VideoObjectName, "video-learning-videos-confirmed", p.VideoObjectName)
	if err != nil {
		return nil, videolearning.UploadFailed("failed to move video to confirmed bucket")
	}

	// Move thumbnail
	err = s.storageRepo.CopyFile(ctx, "video-learning-thumbnails-staging", p.ThumbnailObjectName, "video-learning-thumbnails-confirmed", p.ThumbnailObjectName)
	if err != nil {
		return nil, videolearning.UploadFailed("failed to move thumbnail to confirmed bucket")
	}

	// Create video record in database
	description := pgtype.Text{}
	if p.Description != nil {
		description = pgtype.Text{String: *p.Description, Valid: true}
	}

	video, err := s.videoRepo.CreateVideo(ctx, videolearningdb.CreateVideoParams{
		Title:        p.Title,
		UserID:       profile.UserID,
		Description:  description,
		Views:        0,
		Likes:        0,
		VideoObjName: pgtype.Text{String: p.VideoObjectName, Valid: true},
		ThumbObjName: pgtype.Text{String: p.ThumbnailObjectName, Valid: true},
		CategoryID:   p.CategoryID,
	})
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to create video record")
	}

	// Assign tags to video
	for _, tagName := range p.Tags {
		// Get or create tag
		tag, err := s.videoTagRepo.GetOrCreateTag(ctx, tagName)
		if err != nil {
			continue // Skip failed tags
		}

		// Assign tag to video
		s.videoRepo.AssignTagToVideo(ctx, videolearningdb.AssignTagToVideoParams{
			VideoID: video.ID,
			TagID:   tag.ID,
		})
	}

	// Delete staging video
	err = s.storageRepo.DeleteFile(ctx, "video-learning-videos-staging", p.VideoObjectName)
	if err != nil {
		log.Print("failed to delete staging video file: ", err)
	}

	// Delete staging thumbnail
	err = s.storageRepo.DeleteFile(ctx, "video-learning-thumbnails-staging", p.ThumbnailObjectName)
	if err != nil {
		log.Print("failed to delete staging thumbnail file: ", err)
	}

	return &videolearning.SimpleResponse{
		Success: true,
		Message: "Video upload completed successfully",
		ID:      &video.ID,
	}, nil
}

// UploadThumbnail uploads custom thumbnail for video
func (s *videolearningsrvc) UploadThumbnail(ctx context.Context, p *videolearning.UploadThumbnailPayload) (res *videolearning.UploadResponse, err error) {
	// Validate session and check teacher role
	profile, err := s.validateTeacherRole(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.PermissionDenied("only teachers can upload thumbnails")
	}

	// Generate object name for thumbnail
	objectName := s.generateObjectName(p.Filename, profile.UserID)

	// Upload to staging bucket
	reader := bytes.NewReader(p.File)
	err = s.storageRepo.UploadFile(ctx, "video-learning-thumbnails-staging", objectName, reader, int64(len(p.File)), "image/jpeg")
	if err != nil {
		return nil, videolearning.UploadFailed("failed to upload thumbnail")
	}

	// Generate presigned URL for thumbnail
	presignedURL, err := s.getOrGeneratePresignedURL(ctx, "video-learning-thumbnails-staging", objectName, 4*time.Hour)
	if err != nil {
		return nil, videolearning.UploadFailed("failed to generate presigned URL for thumbnail")
	}

	return &videolearning.UploadResponse{
		ObjectName:   objectName,
		PresignedURL: &presignedURL,
	}, nil
}

// GetOrCreateCategory retrieves or creates a video category
func (s *videolearningsrvc) GetOrCreateCategory(ctx context.Context, p *videolearning.GetOrCreateCategoryPayload) (res *videolearning.VideoCategory, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get or create category
	category, err := s.videoCategoryRepo.GetOrCreateCategory(ctx, p.Name)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get or create category")
	}

	return &videolearning.VideoCategory{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

// GetAllCategories returns all video categories
func (s *videolearningsrvc) GetAllCategories(ctx context.Context, p *videolearning.GetAllCategoriesPayload) (res []*videolearning.VideoCategory, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get all categories
	categories, err := s.videoCategoryRepo.GetAllCategories(ctx)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get categories")
	}

	// Convert to API format
	var apiCategories []*videolearning.VideoCategory
	for _, category := range categories {
		apiCategories = append(apiCategories, &videolearning.VideoCategory{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return apiCategories, nil
}

// GetOrCreateTag retrieves or creates a video tag
func (s *videolearningsrvc) GetOrCreateTag(ctx context.Context, p *videolearning.GetOrCreateTagPayload) (res *videolearning.VideoTag, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get or create tag
	tag, err := s.videoTagRepo.GetOrCreateTag(ctx, p.Name)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get or create tag")
	}

	return &videolearning.VideoTag{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

// GetAllTags returns all video tags
func (s *videolearningsrvc) GetAllTags(ctx context.Context, p *videolearning.GetAllTagsPayload) (res []*videolearning.VideoTag, err error) {
	// Validate session
	_, err = s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get all tags
	tags, err := s.videoTagRepo.GetAllTags(ctx)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to get tags")
	}

	// Convert to API format
	var apiTags []*videolearning.VideoTag
	for _, tag := range tags {
		apiTags = append(apiTags, &videolearning.VideoTag{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return apiTags, nil
}

// ToggleVideoLike toggles like status for a video
func (s *videolearningsrvc) ToggleVideoLike(ctx context.Context, p *videolearning.ToggleVideoLikePayload) (res *videolearning.SimpleResponse, err error) {
	// Validate session
	userID, err := s.validateSessionAndGetUserID(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.InvalidSession("invalid session")
	}

	// Get video to check category
	video, err := s.videoRepo.GetVideoByID(ctx, p.VideoID)
	if err != nil {
		return nil, videolearning.VideoNotFound("video not found")
	}

	cacheKey := fmt.Sprintf("user:like:%d:%d", userID, p.VideoID)

	var increment int
	var message string

	cached, err := s.cacheRepo.Exists(ctx, cacheKey)

	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to check like status")
	}

	if cached {
		value, err := s.cacheRepo.Get(ctx, cacheKey)
		if err != nil {
			return nil, videolearning.ServiceUnavailable("failed to get like status")
		}

		hasLiked := value == "1"

		if hasLiked {
			increment = -1
			message = "Video unliked"
			s.cacheRepo.Set(ctx, cacheKey, "0", 2*time.Hour)
		} else {
			increment = 1
			message = "Video liked"
			s.cacheRepo.Set(ctx, cacheKey, "1", 2*time.Hour)
		}
	} else {
		// check db
		userLike, err := s.userCategoryLikeRepo.GetUserVideoLike(ctx, videolearningdb.GetUserVideoLikeParams{
			UserID:  userID,
			VideoID: p.VideoID,
		})

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// No like record, create new like
				increment = 1
				message = "Video liked"
				s.cacheRepo.Set(ctx, cacheKey, "1", 2*time.Hour)
			} else {
				return nil, videolearning.ServiceUnavailable("failed to get like status")
			}
		} else {
			if userLike.Liked {
				increment = 1
				message = "Video liked"
				s.cacheRepo.Set(ctx, cacheKey, "1", 2*time.Hour)
			} else {
				increment = -1
				message = "Video unliked"
				s.cacheRepo.Set(ctx, cacheKey, "0", 2*time.Hour)
			}
		}
	}

	// Update cached like count
	s.incrementCachedVideoLikes(ctx, p.VideoID, increment)

	// Update user category preferences
	s.incrementCachedUserCategoryLike(ctx, userID, video.CategoryID, increment)

	return &videolearning.SimpleResponse{
		Success: true,
		Message: message,
	}, nil
}

// DeleteVideo deletes user's own video
func (s *videolearningsrvc) DeleteVideo(ctx context.Context, p *videolearning.DeleteVideoPayload) (res *videolearning.SimpleResponse, err error) {
	// Validate session and check teacher role
	profile, err := s.validateTeacherRole(ctx, p.SessionToken)
	if err != nil {
		return nil, videolearning.PermissionDenied("only teachers can delete videos")
	}

	// Get video to verify ownership
	video, err := s.videoRepo.GetVideoByID(ctx, p.VideoID)
	if err != nil {
		return nil, videolearning.VideoNotFound("video not found")
	}

	// Check ownership
	if video.UserID != profile.UserID {
		return nil, videolearning.PermissionDenied("you can only delete your own videos")
	}

	// Delete files from storage and clear cached URLs
	if video.VideoObjName.Valid {
		s.storageRepo.DeleteFile(ctx, "video-learning-videos-confirmed", video.VideoObjName.String)
		s.clearCachedPresignedURL(ctx, "video-learning-videos-confirmed", video.VideoObjName.String)
	}
	if video.ThumbObjName.Valid {
		s.storageRepo.DeleteFile(ctx, "video-learning-thumbnails-confirmed", video.ThumbObjName.String)
		s.clearCachedPresignedURL(ctx, "video-learning-thumbnails-confirmed", video.ThumbObjName.String)
	}

	// Delete video from database
	err = s.videoRepo.DeleteVideo(ctx, p.VideoID)
	if err != nil {
		return nil, videolearning.ServiceUnavailable("failed to delete video")
	}

	return &videolearning.SimpleResponse{
		Success: true,
		Message: "Video deleted successfully",
	}, nil
}
