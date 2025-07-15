package controllers

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/gen/stats"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/mappers"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
)

// --- Leaderboard Methods Implementation ---

func (s *statssrvc) GetCourseLeaderboard(ctx context.Context, payload *stats.GetCourseLeaderboardPayload) (res *stats.CourseLeaderboard, err error) {
	// Validate user session through profiles service
	_, err = s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Get course leaderboard from text service
	textLeaderboard, err := s.textServiceRepo.GetCourseLeaderboard(ctx, &text.GetCourseLeaderboardPayload{
		SessionToken: payload.SessionToken,
		CourseID:     payload.CourseID,
		Limit:        payload.Limit,
	})
	if err != nil {
		return nil, stats.InternalError("Failed to get course leaderboard: " + err.Error())
	}

	// Convert from text service format to stats service format
	var statsEntries []*stats.LeaderboardEntry
	for _, textEntry := range textLeaderboard.Entries {
		statsEntry := &stats.LeaderboardEntry{
			UserID:               textEntry.UserID,
			Username:            textEntry.Username,
			CompletionPercentage: textEntry.CompletionPercentage,
			CompletedArticles:   textEntry.CompletedArticles,
			TotalArticles:       textEntry.TotalArticles,
			Rank:                textEntry.Rank,
			LastActivity:        textEntry.LastActivity,
		}
		statsEntries = append(statsEntries, statsEntry)
	}

	leaderboard := &stats.CourseLeaderboard{
		CourseID:         textLeaderboard.CourseID,
		CourseTitle:      textLeaderboard.CourseTitle,
		Entries:          statsEntries,
		TotalParticipants: textLeaderboard.TotalParticipants,
		GeneratedAt:      textLeaderboard.GeneratedAt,
	}

	return leaderboard, nil
}

func (s *statssrvc) GetUserOverallStats(ctx context.Context, payload *stats.GetUserOverallStatsPayload) (res *stats.UserOverallStats, err error) {
	// Validate user session
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Determine target user
	targetUserID := profileInfo.UserID
	targetUsername := profileInfo.FirstName + " " + profileInfo.LastName
	
	if payload.UserID != nil {
		// Only allow viewing other users' stats if user is teacher or admin
		if profileInfo.Role != "teacher" && profileInfo.Role != "admin" {
			return nil, stats.PermissionDenied("Only teachers and admins can view other users' statistics")
		}
		targetUserID = *payload.UserID
		
		// Get target user's profile
		targetProfile, err := s.profilesServiceRepo.GetPublicProfileByID(ctx, &profiles.GetPublicProfileByIDPayload{
			UserID: targetUserID,
		})
		if err != nil {
			return nil, stats.NotFound("Target user not found: " + err.Error())
		}
		targetUsername = targetProfile.FirstName + " " + targetProfile.LastName
	}

	// Get all courses for the user
	courses, err := s.textServiceRepo.ListCourses(ctx, &text.ListCoursesPayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.InternalError("Failed to get courses: " + err.Error())
	}

	// Calculate overall statistics
	var courseEntries []*stats.CourseStatsEntry
	var totalCompletedArticles int64
	var totalArticles int64
	var completedCourses int64

	for _, course := range courses {
		// Get completion stats for each course
		courseStats, err := s.textServiceRepo.GetCourseCompletionStats(ctx, &text.GetCourseCompletionStatsPayload{
			SessionToken: payload.SessionToken,
			CourseID:     course.ID,
		})
		if err != nil {
			// Skip this course if we can't get stats
			continue
		}

		completionPercentage := courseStats.CompletionPercentage
		if completionPercentage >= 100.0 {
			completedCourses++
		}

		courseEntry := &stats.CourseStatsEntry{
			CourseID:            course.ID,
			CourseTitle:         course.Title,
			CompletionPercentage: completionPercentage,
			CompletedArticles:   courseStats.CompletedArticles,
			TotalArticles:       courseStats.TotalArticles,
			LastAccessed:        mappers.Int64Ptr(mappers.TimestampToMillis(1672531200000)), // Mock timestamp
		}

		courseEntries = append(courseEntries, courseEntry)
		totalCompletedArticles += courseStats.CompletedArticles
		totalArticles += courseStats.TotalArticles
	}

	// Calculate overall completion percentage
	var overallCompletionPercentage float64
	if totalArticles > 0 {
		overallCompletionPercentage = float64(totalCompletedArticles) / float64(totalArticles) * 100
	}

	userStats := &stats.UserOverallStats{
		UserID:                       targetUserID,
		Username:                    targetUsername,
		TotalCoursesEnrolled:        int64(len(courses)),
		TotalCoursesCompleted:       completedCourses,
		TotalArticlesCompleted:      totalCompletedArticles,
		OverallCompletionPercentage: overallCompletionPercentage,
		Courses:                     courseEntries,
		LastActivity:                mappers.Int64Ptr(mappers.TimestampToMillis(1672531200000)), // Mock timestamp
	}

	return userStats, nil
}

// --- Progress Tracking Methods Implementation ---

func (s *statssrvc) GetUserCourseProgress(ctx context.Context, payload *stats.GetUserCourseProgressPayload) (res *stats.UserCourseProgressStats, err error) {
	// Validate user session
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Determine target user
	var targetUserID *int64
	if payload.UserID != nil {
		// Only allow viewing other users' progress if user is teacher or admin
		if profileInfo.Role != "teacher" && profileInfo.Role != "admin" {
			return nil, stats.PermissionDenied("Only teachers and admins can view other users' progress")
		}
		targetUserID = payload.UserID
	}

	// Get user course progress from text service
	progressData, err := s.textServiceRepo.GetUserCourseProgress(ctx, &text.GetUserCourseProgressPayload{
		SessionToken: payload.SessionToken,
		CourseID:     payload.CourseID,
		UserID:       targetUserID,
	})
	if err != nil {
		return nil, stats.InternalError("Failed to get user course progress: " + err.Error())
	}

	// Convert sections to stats format
	var sectionEntries []*stats.SectionProgressEntry
	for _, section := range progressData.Sections {
		sectionEntry := &stats.SectionProgressEntry{
			SectionID:            section.ID,
			SectionTitle:         section.Title,
			SectionOrder:         section.Order,
			CompletionPercentage: section.CompletionPercentage,
			CompletedArticles:    section.CompletedArticles,
			TotalArticles:        section.TotalArticles,
			LastAccessed:         mappers.Int64Ptr(section.UpdatedAt), // Using UpdatedAt as LastAccessed
		}
		sectionEntries = append(sectionEntries, sectionEntry)
	}

	// Calculate estimated completion time (mock calculation)
	remainingArticles := progressData.TotalArticles - progressData.CompletedArticles
	estimatedTime := remainingArticles * 10 // 10 minutes per article estimate

	userCourseProgress := &stats.UserCourseProgressStats{
		UserID:                      progressData.UserID,
		CourseID:                   progressData.Course.ID,
		CourseTitle:                progressData.Course.Title,
		OverallCompletionPercentage: progressData.CompletionPercentage,
		CompletedArticles:          progressData.CompletedArticles,
		TotalArticles:              progressData.TotalArticles,
		Sections:                   sectionEntries,
		EnrollmentDate:             mappers.Int64Ptr(mappers.TimestampToMillis(1672531200000)), // Mock timestamp
		LastAccessed:               mappers.Int64Ptr(progressData.LastAccessed),
		EstimatedCompletionTime:    mappers.Int64Ptr(estimatedTime),
	}

	return userCourseProgress, nil
}

func (s *statssrvc) GetUserCompletedArticles(ctx context.Context, payload *stats.GetUserCompletedArticlesPayload) (res *stats.UserCompletedArticles, err error) {
	// Validate user session
	profileInfo, err := s.profilesServiceRepo.GetCompleteProfile(ctx, &profiles.GetCompleteProfilePayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.Unauthorized("Invalid session: " + err.Error())
	}

	// Determine target user
	targetUserID := profileInfo.UserID
	if payload.UserID != nil {
		// Only allow viewing other users' completed articles if user is teacher or admin
		if profileInfo.Role != "teacher" && profileInfo.Role != "admin" {
			return nil, stats.PermissionDenied("Only teachers and admins can view other users' completed articles")
		}
		targetUserID = *payload.UserID
	}

	// Get all courses to find completed articles
	// Note: This is a workaround since we don't have a direct "GetUserCompletedArticles" method in text service
	courses, err := s.textServiceRepo.ListCourses(ctx, &text.ListCoursesPayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, stats.InternalError("Failed to get courses: " + err.Error())
	}

	// Collect completed articles from all courses
	var articleEntries []*stats.CompletedArticleEntry
	var lastCompletedAt int64
	limit := payload.Limit

	for _, course := range courses {
		if int64(len(articleEntries)) >= limit {
			break
		}

		// Get user progress for this course
		progressData, err := s.textServiceRepo.GetUserCourseProgress(ctx, &text.GetUserCourseProgressPayload{
			SessionToken: payload.SessionToken,
			CourseID:     course.ID,
			UserID:       &targetUserID,
		})
		if err != nil {
			continue // Skip this course if we can't get progress
		}

		// Extract completed articles from sections
		for _, section := range progressData.Sections {
			for _, article := range section.Articles {
				if article.Completed && int64(len(articleEntries)) < limit {
					articleEntry := &stats.CompletedArticleEntry{
						ArticleID:           article.ID,
						ArticleTitle:        article.Title,
						SectionID:           article.SectionID,
						SectionTitle:        section.Title,
						CourseID:            course.ID,
						CourseTitle:         course.Title,
						CompletedAt:         article.CompletedAt,
						ReadingTimeEstimate: mappers.Int64Ptr(10), // Mock estimate
					}
					articleEntries = append(articleEntries, articleEntry)
					
					if article.CompletedAt > lastCompletedAt {
						lastCompletedAt = article.CompletedAt
					}
				}
			}
		}
	}

	userCompletedArticles := &stats.UserCompletedArticles{
		UserID:          targetUserID,
		TotalCompleted:  int64(len(articleEntries)),
		Articles:        articleEntries,
		LastCompletedAt: mappers.Int64Ptr(lastCompletedAt),
	}

	return userCompletedArticles, nil
}
