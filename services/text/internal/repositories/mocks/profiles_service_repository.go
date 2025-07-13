package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

type MockProfilesServiceRepository struct {
	mock.Mock
}

func (m *MockProfilesServiceRepository) GetCompleteProfile(ctx context.Context, in *profiles.GetCompleteProfilePayload) (*profiles.CompleteProfileResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*profiles.CompleteProfileResponse), args.Error(1)
}
func (m *MockProfilesServiceRepository) GetPublicProfileByID(ctx context.Context, in *profiles.GetPublicProfileByIDPayload) (*profiles.PublicProfileResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*profiles.PublicProfileResponse), args.Error(1)
}
func (m *MockProfilesServiceRepository) ValidateUserRole(ctx context.Context, in *profiles.ValidateUserRolePayload) (*profiles.RoleValidationResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*profiles.RoleValidationResponse), args.Error(1)
}