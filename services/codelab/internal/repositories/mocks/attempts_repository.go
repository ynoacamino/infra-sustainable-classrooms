package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

// MockAttemptsRepository is a mock implementation of AttemptsRepository
type MockAttemptsRepository struct {
	mock.Mock
}

func (m *MockAttemptsRepository) CreateAttempt(ctx context.Context, arg codelabdb.CreateAttemptParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockAttemptsRepository) GetAttempt(ctx context.Context, id int64) (codelabdb.Attempt, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(codelabdb.Attempt), args.Error(1)
}

func (m *MockAttemptsRepository) GetAttemptsByAnswer(ctx context.Context, answerID int64) ([]codelabdb.Attempt, error) {
	args := m.Called(ctx, answerID)
	return args.Get(0).([]codelabdb.Attempt), args.Error(1)
}

func (m *MockAttemptsRepository) GetAttemptsByUserAndExercise(ctx context.Context, arg codelabdb.GetAttemptsByUserAndExerciseParams) ([]codelabdb.Attempt, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).([]codelabdb.Attempt), args.Error(1)
}

func (m *MockAttemptsRepository) GetLatestAttemptByAnswer(ctx context.Context, answerID int64) (codelabdb.Attempt, error) {
	args := m.Called(ctx, answerID)
	return args.Get(0).(codelabdb.Attempt), args.Error(1)
}

func (m *MockAttemptsRepository) CountAttemptsByAnswer(ctx context.Context, answerID int64) (int64, error) {
	args := m.Called(ctx, answerID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAttemptsRepository) CountSuccessfulAttemptsByAnswer(ctx context.Context, answerID int64) (int64, error) {
	args := m.Called(ctx, answerID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAttemptsRepository) CountTotalAttemptsByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAttemptsRepository) CountSuccessfulAttemptsByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAttemptsRepository) GetAttemptsWithAnswerInfo(ctx context.Context, exerciseID int64) ([]codelabdb.GetAttemptsWithAnswerInfoRow, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).([]codelabdb.GetAttemptsWithAnswerInfoRow), args.Error(1)
}

func (m *MockAttemptsRepository) GetUserAttemptsForExercise(ctx context.Context, arg codelabdb.GetUserAttemptsForExerciseParams) ([]codelabdb.Attempt, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).([]codelabdb.Attempt), args.Error(1)
}
