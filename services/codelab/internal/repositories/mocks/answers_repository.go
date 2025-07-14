package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	codelabdb "github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/database"
)

// MockAnswersRepository is a mock implementation of AnswersRepository
type MockAnswersRepository struct {
	mock.Mock
}

func (m *MockAnswersRepository) CheckIfAnswerExists(ctx context.Context, arg codelabdb.CheckIfAnswerExistsParams) (int32, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(int32), args.Error(1)
}

func (m *MockAnswersRepository) CreateAnswer(ctx context.Context, arg codelabdb.CreateAnswerParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockAnswersRepository) GetAnswerByUserAndExercise(ctx context.Context, arg codelabdb.GetAnswerByUserAndExerciseParams) (codelabdb.Answer, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(codelabdb.Answer), args.Error(1)
}

func (m *MockAnswersRepository) ListAnswersByExercise(ctx context.Context, exerciseID int64) ([]codelabdb.Answer, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).([]codelabdb.Answer), args.Error(1)
}

func (m *MockAnswersRepository) ListAnswersByUser(ctx context.Context, userID int64) ([]codelabdb.Answer, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]codelabdb.Answer), args.Error(1)
}

func (m *MockAnswersRepository) UpdateAnswerCompleted(ctx context.Context, arg codelabdb.UpdateAnswerCompletedParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockAnswersRepository) CountCompletedAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockAnswersRepository) CountTotalAnswersByExercise(ctx context.Context, exerciseID int64) (int64, error) {
	args := m.Called(ctx, exerciseID)
	return args.Get(0).(int64), args.Error(1)
}
