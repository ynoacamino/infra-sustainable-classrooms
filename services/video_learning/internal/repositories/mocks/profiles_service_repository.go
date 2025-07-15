package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

// MockProfilesServiceRepository is a mock implementation of ProfilesServiceRepository
type MockProfilesServiceRepository struct {
	mock.Mock
}

func (m *MockProfilesServiceRepository) GetCompleteProfile(ctx context.Context, payload *profiles.GetCompleteProfilePayload) (*profiles.CompleteProfileResponse, error) {
	args := m.Called(ctx, payload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*profiles.CompleteProfileResponse), args.Error(1)
}

func (m *MockProfilesServiceRepository) GetPublicProfileByID(ctx context.Context, payload *profiles.GetPublicProfileByIDPayload) (*profiles.PublicProfileResponse, error) {
	args := m.Called(ctx, payload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*profiles.PublicProfileResponse), args.Error(1)
}

func (m *MockProfilesServiceRepository) ValidateUserRole(ctx context.Context, payload *profiles.ValidateUserRolePayload) (*profiles.RoleValidationResponse, error) {
	args := m.Called(ctx, payload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*profiles.RoleValidationResponse), args.Error(1)
}
