package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

// MockStudentProfileRepository is a mock implementation of StudentProfileRepository
type MockStudentProfileRepository struct {
	mock.Mock
}

func (m *MockStudentProfileRepository) CreateStudentProfile(ctx context.Context, params profilesdb.CreateStudentProfileParams) (profilesdb.StudentProfile, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(profilesdb.StudentProfile), args.Error(1)
}

func (m *MockStudentProfileRepository) GetStudentProfileByProfileId(ctx context.Context, profileID int64) (profilesdb.StudentProfile, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).(profilesdb.StudentProfile), args.Error(1)
}

func (m *MockStudentProfileRepository) GetStudentProfileByUserId(ctx context.Context, userID int64) (profilesdb.StudentProfile, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(profilesdb.StudentProfile), args.Error(1)
}

func (m *MockStudentProfileRepository) UpdateStudentProfile(ctx context.Context, params profilesdb.UpdateStudentProfileParams) (profilesdb.StudentProfile, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(profilesdb.StudentProfile), args.Error(1)
}

func (m *MockStudentProfileRepository) GetCompleteStudentProfile(ctx context.Context, userID int64) (profilesdb.GetCompleteStudentProfileRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(profilesdb.GetCompleteStudentProfileRow), args.Error(1)
}
