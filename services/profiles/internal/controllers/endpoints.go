package controllers

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

func (s *profilessrvc) CreateStudentProfile(ctx context.Context, payload *profiles.CreateStudentProfilePayload) (res *profiles.StudentProfileResponse, err error) {
	// Validate session
	userInfo, err := s.authServiceRepo.ValidateUser(ctx, &auth.ValidateUserPayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, profiles.Unauthorized("invalid session token")
	}

	// Check if profile already exists
	exists, err := s.profileRepo.CheckProfileExists(ctx, userInfo.User.ID)
	if err != nil {
		return nil, profiles.InvalidInput("error checking profile existence")
	}
	if exists {
		return nil, profiles.ProfileAlreadyExists("profile already exists for this user")
	}

	// Create profile
	profileParams := profilesdb.CreateProfileParams{
		UserID:    userInfo.User.ID,
		Role:      "student",
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     stringToPgText(payload.Phone),
		AvatarUrl: stringToPgText(payload.AvatarURL),
		Bio:       stringToPgText(payload.Bio),
	}

	createdProfile, err := s.profileRepo.CreateProfile(ctx, profileParams)
	if err != nil {
		return nil, profiles.InvalidInput("failed to create profile")
	}

	// Create student profile
	studentParams := profilesdb.CreateStudentProfileParams{
		ProfileID:  createdProfile.ID,
		GradeLevel: payload.GradeLevel,
		Major:      stringToPgText(payload.Major),
	}

	createdStudentProfile, err := s.studentProfileRepo.CreateStudentProfile(ctx, studentParams)
	if err != nil {
		return nil, profiles.InvalidInput("failed to create student profile")
	}

	// Map to response
	return mapToStudentProfileResponse(createdProfile, createdStudentProfile), nil
}

func (s *profilessrvc) CreateTeacherProfile(ctx context.Context, payload *profiles.CreateTeacherProfilePayload) (res *profiles.TeacherProfileResponse, err error) {
	// Validate session
	userInfo, err := s.authServiceRepo.ValidateUser(ctx, &auth.ValidateUserPayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, profiles.Unauthorized("invalid session token")
	}

	// Check if profile already exists
	exists, err := s.profileRepo.CheckProfileExists(ctx, userInfo.User.ID)
	if err != nil {
		return nil, profiles.InvalidInput("error checking profile existence")
	}
	if exists {
		return nil, profiles.ProfileAlreadyExists("profile already exists for this user")
	}

	// Create profile
	profileParams := profilesdb.CreateProfileParams{
		UserID:    userInfo.User.ID,
		Role:      "teacher",
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     stringToPgText(payload.Phone),
		AvatarUrl: stringToPgText(payload.AvatarURL),
		Bio:       stringToPgText(payload.Bio),
	}

	createdProfile, err := s.profileRepo.CreateProfile(ctx, profileParams)
	if err != nil {
		return nil, profiles.InvalidInput("failed to create profile")
	}

	// Create teacher profile
	teacherParams := profilesdb.CreateTeacherProfileParams{
		ProfileID: createdProfile.ID,
		Position:  payload.Position,
	}

	createdTeacherProfile, err := s.teacherProfileRepo.CreateTeacherProfile(ctx, teacherParams)
	if err != nil {
		return nil, profiles.InvalidInput("failed to create teacher profile")
	}

	// Map to response
	return mapToTeacherProfileResponse(createdProfile, createdTeacherProfile), nil
}

func (s *profilessrvc) GetCompleteProfile(ctx context.Context, payload *profiles.GetCompleteProfilePayload) (res *profiles.CompleteProfileResponse, err error) {
	// Validate session
	userInfo, err := s.authServiceRepo.ValidateUser(ctx, &auth.ValidateUserPayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, profiles.Unauthorized("invalid session token")
	}

	// Get profile to check role
	profile, err := s.profileRepo.GetProfileByUserId(ctx, userInfo.User.ID)
	if err != nil {
		return nil, profiles.ProfileNotFound("profile not found")
	}

	// Get complete profile based on role
	switch profile.Role {
	case "student":
		completeProfile, err := s.studentProfileRepo.GetCompleteStudentProfile(ctx, userInfo.User.ID)
		if err != nil {
			return nil, profiles.ProfileNotFound("student profile not found")
		}
		return mapToCompleteProfileResponse(completeProfile, "student"), nil
	case "teacher":
		completeProfile, err := s.teacherProfileRepo.GetCompleteTeacherProfile(ctx, userInfo.User.ID)
		if err != nil {
			return nil, profiles.ProfileNotFound("teacher profile not found")
		}
		return mapToCompleteProfileResponse(completeProfile, "teacher"), nil
	default:
		return nil, profiles.InvalidRole("invalid role")
	}
}

func (s *profilessrvc) GetPublicProfileByID(ctx context.Context, payload *profiles.GetPublicProfileByIDPayload) (res *profiles.PublicProfileResponse, err error) {
	// Public endpoint - no session validation needed
	// Get public profile
	publicProfile, err := s.profileRepo.GetPublicProfileByUserId(ctx, payload.UserID)
	if err != nil {
		return nil, profiles.ProfileNotFound("profile not found")
	}

	// Map to response
	return mapToPublicProfileResponse(publicProfile), nil
}

func (s *profilessrvc) UpdateProfile(ctx context.Context, payload *profiles.UpdateProfilePayload) (res *profiles.ProfileResponse, err error) {
	// Validate session
	userInfo, err := s.authServiceRepo.ValidateUser(ctx, &auth.ValidateUserPayload{
		SessionToken: payload.SessionToken,
	})
	if err != nil {
		return nil, profiles.Unauthorized("invalid session token")
	}

	// User can only update their own profile
	userID := userInfo.User.ID

	// Check if profile exists
	existingProfile, err := s.profileRepo.GetProfileByUserId(ctx, userID)
	if err != nil {
		return nil, profiles.ProfileNotFound("profile not found")
	}

	// Update profile
	updateParams := profilesdb.UpdateProfileParams{
		UserID:    userID,
		FirstName: getStringOrDefault(payload.FirstName, existingProfile.FirstName),
		LastName:  getStringOrDefault(payload.LastName, existingProfile.LastName),
		Email:     getStringOrDefault(payload.Email, existingProfile.Email),
		Phone:     getStringPtrOrDefault(payload.Phone, existingProfile.Phone),
		AvatarUrl: getStringPtrOrDefault(payload.AvatarURL, existingProfile.AvatarUrl),
		Bio:       getStringPtrOrDefault(payload.Bio, existingProfile.Bio),
	}

	updatedProfile, err := s.profileRepo.UpdateProfile(ctx, updateParams)
	if err != nil {
		return nil, profiles.InvalidInput("failed to update profile")
	}

	// Map to response
	return mapToProfileResponse(updatedProfile), nil
}

func (s *profilessrvc) ValidateUserRole(ctx context.Context, payload *profiles.ValidateUserRolePayload) (res *profiles.RoleValidationResponse, err error) {
	// This endpoint is for microservice communication, no session validation needed
	// Get profile by user ID
	profile, err := s.profileRepo.GetProfileByUserId(ctx, payload.UserID)
	if err != nil {
		// Return empty role if profile not found
		return &profiles.RoleValidationResponse{
			UserID: payload.UserID,
			Role:   "",
		}, nil
	}

	// Return role validation response
	return &profiles.RoleValidationResponse{
		UserID: payload.UserID,
		Role:   profile.Role,
	}, nil
}

// Helper functions for updates
func getStringOrDefault(new *string, existing string) string {
	if new != nil {
		return *new
	}
	return existing
}

func getStringPtrOrDefault(new *string, existing pgtype.Text) pgtype.Text {
	if new != nil {
		return stringToPgText(new)
	}
	return existing
}
