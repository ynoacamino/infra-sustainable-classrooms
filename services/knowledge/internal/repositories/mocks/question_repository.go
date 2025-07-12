package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

// MockQuestionRepository is a mock implementation of QuestionRepository
type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) CreateQuestion(ctx context.Context, params knowledgedb.CreateQuestionParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetQuestionById(ctx context.Context, id int64) (knowledgedb.Question, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(knowledgedb.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetQuestionsByTestId(ctx context.Context, testID int64) ([]knowledgedb.Question, error) {
	args := m.Called(ctx, testID)
	return args.Get(0).([]knowledgedb.Question), args.Error(1)
}

func (m *MockQuestionRepository) UpdateQuestion(ctx context.Context, params knowledgedb.UpdateQuestionParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockQuestionRepository) DeleteQuestion(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
