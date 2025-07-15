package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

// MockTeacherProfileRepository is a mock implementation of TeacherProfileRepository
type MockTeacherProfileRepository struct {
	mock.Mock
}

func (m *MockTeacherProfileRepository) CreateTeacherProfile(ctx context.Context, params profilesdb.CreateTeacherProfileParams) (profilesdb.TeacherProfile, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(profilesdb.TeacherProfile), args.Error(1)
}

func (m *MockTeacherProfileRepository) GetTeacherProfileByProfileId(ctx context.Context, profileID int64) (profilesdb.TeacherProfile, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).(profilesdb.TeacherProfile), args.Error(1)
}

func (m *MockTeacherProfileRepository) GetTeacherProfileByUserId(ctx context.Context, userID int64) (profilesdb.TeacherProfile, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(profilesdb.TeacherProfile), args.Error(1)
}

func (m *MockTeacherProfileRepository) UpdateTeacherProfile(ctx context.Context, params profilesdb.UpdateTeacherProfileParams) (profilesdb.TeacherProfile, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(profilesdb.TeacherProfile), args.Error(1)
}

func (m *MockTeacherProfileRepository) GetCompleteTeacherProfile(ctx context.Context, userID int64) (profilesdb.GetCompleteTeacherProfileRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(profilesdb.GetCompleteTeacherProfileRow), args.Error(1)
}
