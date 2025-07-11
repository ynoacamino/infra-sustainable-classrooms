package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

// MockProfileRepository is a mock implementation of ProfileRepository
type MockProfileRepository struct {
	mock.Mock
}

func (m *MockProfileRepository) CreateProfile(ctx context.Context, params profilesdb.CreateProfileParams) (profilesdb.Profile, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(profilesdb.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetProfileByUserId(ctx context.Context, userID int64) (profilesdb.Profile, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(profilesdb.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetProfileByEmail(ctx context.Context, email string) (profilesdb.Profile, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(profilesdb.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetPublicProfileByUserId(ctx context.Context, userID int64) (profilesdb.GetPublicProfileByUserIdRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(profilesdb.GetPublicProfileByUserIdRow), args.Error(1)
}

func (m *MockProfileRepository) UpdateProfile(ctx context.Context, params profilesdb.UpdateProfileParams) (profilesdb.Profile, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(profilesdb.Profile), args.Error(1)
}

func (m *MockProfileRepository) DeactivateProfile(ctx context.Context, userID int64) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockProfileRepository) CheckProfileExists(ctx context.Context, userID int64) (bool, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockProfileRepository) GetProfileRole(ctx context.Context, userID int64) (string, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(string), args.Error(1)
}
