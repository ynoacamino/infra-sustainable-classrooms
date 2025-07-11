package controllers

import (
	"github.com/jackc/pgx/v5/pgtype"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

// validateOwnership checks if the user can only access their own profile
func (s *profilessrvc) validateOwnership(userID int64, targetUserID int64) error {
	if userID != targetUserID {
		return profiles.PermissionDenied("cannot access other user's profile")
	}
	return nil
}

// mapToStudentProfileResponse converts database types to Goa response
func mapToStudentProfileResponse(profile profilesdb.Profile, studentProfile profilesdb.StudentProfile) *profiles.StudentProfileResponse {
	return &profiles.StudentProfileResponse{
		UserID:     profile.UserID,
		FirstName:  profile.FirstName,
		LastName:   profile.LastName,
		Email:      profile.Email,
		Phone:      getStringPtr(profile.Phone),
		AvatarURL:  getStringPtr(profile.AvatarUrl),
		Bio:        getStringPtr(profile.Bio),
		GradeLevel: studentProfile.GradeLevel,
		Major:      getStringPtr(studentProfile.Major),
		CreatedAt:  timestampToMillis(profile.CreatedAt),
		UpdatedAt:  timestampToMillisPtr(profile.UpdatedAt),
		IsActive:   profile.IsActive,
	}
}

// mapToTeacherProfileResponse converts database types to Goa response
func mapToTeacherProfileResponse(profile profilesdb.Profile, teacherProfile profilesdb.TeacherProfile) *profiles.TeacherProfileResponse {
	return &profiles.TeacherProfileResponse{
		UserID:    profile.UserID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Phone:     getStringPtr(profile.Phone),
		AvatarURL: getStringPtr(profile.AvatarUrl),
		Bio:       getStringPtr(profile.Bio),
		Position:  teacherProfile.Position,
		CreatedAt: timestampToMillis(profile.CreatedAt),
		UpdatedAt: timestampToMillisPtr(profile.UpdatedAt),
		IsActive:  profile.IsActive,
	}
}

// mapToProfileResponse converts database profile to basic response
func mapToProfileResponse(profile profilesdb.Profile) *profiles.ProfileResponse {
	return &profiles.ProfileResponse{
		UserID:    profile.UserID,
		Role:      profile.Role,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     profile.Email,
		Phone:     getStringPtr(profile.Phone),
		AvatarURL: getStringPtr(profile.AvatarUrl),
		Bio:       getStringPtr(profile.Bio),
		CreatedAt: timestampToMillis(profile.CreatedAt),
		UpdatedAt: timestampToMillisPtr(profile.UpdatedAt),
		IsActive:  profile.IsActive,
	}
}

// mapToCompleteProfileResponse converts complete profile data to response
func mapToCompleteProfileResponse(completeProfile any, role string) *profiles.CompleteProfileResponse {
	switch role {
	case "student":
		if sp, ok := completeProfile.(profilesdb.GetCompleteStudentProfileRow); ok {
			return &profiles.CompleteProfileResponse{
				UserID:     sp.UserID,
				Role:       sp.Role,
				FirstName:  sp.FirstName,
				LastName:   sp.LastName,
				Email:      sp.Email,
				Phone:      getStringPtr(sp.Phone),
				AvatarURL:  getStringPtr(sp.AvatarUrl),
				Bio:        getStringPtr(sp.Bio),
				CreatedAt:  timestampToMillis(sp.CreatedAt),
				UpdatedAt:  timestampToMillisPtr(sp.UpdatedAt),
				IsActive:   sp.IsActive,
				GradeLevel: &sp.GradeLevel,
				Major:      getStringPtr(sp.Major),
			}
		}
	case "teacher":
		if tp, ok := completeProfile.(profilesdb.GetCompleteTeacherProfileRow); ok {
			return &profiles.CompleteProfileResponse{
				UserID:    tp.UserID,
				Role:      tp.Role,
				FirstName: tp.FirstName,
				LastName:  tp.LastName,
				Email:     tp.Email,
				Phone:     getStringPtr(tp.Phone),
				AvatarURL: getStringPtr(tp.AvatarUrl),
				Bio:       getStringPtr(tp.Bio),
				CreatedAt: timestampToMillis(tp.CreatedAt),
				UpdatedAt: timestampToMillisPtr(tp.UpdatedAt),
				IsActive:  tp.IsActive,
				Position:  &tp.Position,
			}
		}
	}
	return nil
}

// mapToPublicProfileResponse converts profile to public response
func mapToPublicProfileResponse(publicProfile profilesdb.GetPublicProfileByUserIdRow) *profiles.PublicProfileResponse {
	return &profiles.PublicProfileResponse{
		UserID:    publicProfile.UserID,
		Role:      publicProfile.Role,
		FirstName: publicProfile.FirstName,
		LastName:  publicProfile.LastName,
		AvatarURL: getStringPtr(publicProfile.AvatarUrl),
		Bio:       getStringPtr(publicProfile.Bio),
		IsActive:  publicProfile.IsActive,
	}
}

// Helper function to safely get string pointer values from pgtype.Text
func getStringPtr(text pgtype.Text) *string {
	if !text.Valid {
		return nil
	}
	return &text.String
}

// Helper function to convert pgtype.Timestamptz to int64 (Unix timestamp)
func timestampToMillis(timestamp pgtype.Timestamptz) int64 {
	if !timestamp.Valid {
		return 0
	}
	return timestamp.Time.UnixMilli()
}

// Helper function to convert pgtype.Timestamptz to *int64 (Unix timestamp)
func timestampToMillisPtr(timestamp pgtype.Timestamptz) *int64 {
	if !timestamp.Valid {
		return nil
	}
	timestampPointer := timestamp.Time.UnixMilli()
	return &timestampPointer
}

// Helper function to convert string to pgtype.Text
func stringToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

// Helper function to convert string to pgtype.Text (non-pointer)
func stringToPgTextDirect(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}
