package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

// MockSubmissionRepository is a mock implementation of SubmissionRepository
type MockSubmissionRepository struct {
	mock.Mock
}

func (m *MockSubmissionRepository) CreateSubmission(ctx context.Context, params knowledgedb.CreateSubmissionParams) (knowledgedb.TestSubmission, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(knowledgedb.TestSubmission), args.Error(1)
}

func (m *MockSubmissionRepository) GetSubmissionById(ctx context.Context, id int64) (knowledgedb.TestSubmission, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(knowledgedb.TestSubmission), args.Error(1)
}

func (m *MockSubmissionRepository) GetUserSubmissions(ctx context.Context, userID int64) ([]knowledgedb.GetUserSubmissionsRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]knowledgedb.GetUserSubmissionsRow), args.Error(1)
}

func (m *MockSubmissionRepository) CheckUserCompletedTest(ctx context.Context, params knowledgedb.CheckUserCompletedTestParams) (bool, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(bool), args.Error(1)
}
