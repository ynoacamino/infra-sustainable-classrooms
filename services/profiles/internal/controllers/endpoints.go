package controllers

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

func (s *profilessrvc) CreateStudentProfile(context.Context, *profiles.CreateStudentProfilePayload) (res *profiles.StudentProfileResponse, err error) {

	return
}

func (s *profilessrvc) CreateTeacherProfile(context.Context, *profiles.CreateTeacherProfilePayload) (res *profiles.TeacherProfileResponse, err error) {
	// only teachers can create teacher profiles
	return
}

func (s *profilessrvc) GetCompleteProfile(context.Context, *profiles.GetCompleteProfilePayload) (res *profiles.CompleteProfileResponse, err error) {
	// all users can get their complete profile
	return
}

func (s *profilessrvc) GetPublicProfileByID(ctx context.Context, payload *profiles.GetPublicProfileByIDPayload) (res *profiles.PublicProfileResponse, err error) {
	// all users can get public profiles by user ID
	return
}

func (s *profilessrvc) UpdateProfile(ctx context.Context, payload *profiles.UpdateProfilePayload) (res *profiles.ProfileResponse, err error) {
	// all users can update their own profile
	return
}

func (s *profilessrvc) ValidateUserRole(ctx context.Context, payload *profiles.ValidateUserRolePayload) (res *profiles.RoleValidationResponse, err error) {
	// endpoint for microservice to validate user roles
	return
}
