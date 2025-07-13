package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

// MockTestsRepository is a mock implementation of TestsRepository
type MockTestsRepository struct {
	mock.Mock
}

func (m *MockTestsRepository) CreateTest(ctx context.Context, arg codelabdb.CreateTestParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockTestsRepository) GetTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).([]codelabdb.Test), args.Error(1)
}

func (m *MockTestsRepository) GetPublicTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).([]codelabdb.Test), args.Error(1)
}

func (m *MockTestsRepository) GetHiddenTestsByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Test, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).([]codelabdb.Test), args.Error(1)
}

func (m *MockTestsRepository) UpdateTest(ctx context.Context, arg codelabdb.UpdateTestParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockTestsRepository) DeleteTest(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTestsRepository) DeleteTestsByExercise(ctx context.Context, exerciseID int64) error {
	args := m.Called(ctx, exerciseID)
	return args.Error(0)
}
