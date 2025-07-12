package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

// MockAnswerRepository is a mock implementation of AnswerRepository
type MockAnswerRepository struct {
	mock.Mock
}

func (m *MockAnswerRepository) CreateAnswerSubmission(ctx context.Context, params knowledgedb.CreateAnswerSubmissionParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockAnswerRepository) GetAnswersBySubmission(ctx context.Context, submissionID int64) ([]knowledgedb.GetAnswersBySubmissionRow, error) {
	args := m.Called(ctx, submissionID)
	return args.Get(0).([]knowledgedb.GetAnswersBySubmissionRow), args.Error(1)
}
