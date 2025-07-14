package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

// MockExercisesRepository is a mock implementation of ExercisesRepository
type MockExercisesRepository struct {
	mock.Mock
}

func (m *MockExercisesRepository) CreateExercise(ctx context.Context, arg codelabdb.CreateExerciseParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockExercisesRepository) GetExerciseById(ctx context.Context, id int64) (codelabdb.Exercise, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(codelabdb.Exercise), args.Error(1)
}

func (m *MockExercisesRepository) GetExerciseToResolveById(ctx context.Context, id int64) (codelabdb.GetExerciseToResolveByIdRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(codelabdb.GetExerciseToResolveByIdRow), args.Error(1)
}

func (m *MockExercisesRepository) ListExercises(ctx context.Context) ([]codelabdb.Exercise, error) {
	args := m.Called(ctx)
	return args.Get(0).([]codelabdb.Exercise), args.Error(1)
}

func (m *MockExercisesRepository) ListExercisesToResolve(ctx context.Context) ([]codelabdb.ListExercisesToResolveRow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]codelabdb.ListExercisesToResolveRow), args.Error(1)
}

func (m *MockExercisesRepository) UpdateExercise(ctx context.Context, arg codelabdb.UpdateExerciseParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockExercisesRepository) DeleteExercise(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
