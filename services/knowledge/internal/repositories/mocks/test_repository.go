package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	knowledgedb "github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/database"
)

// MockTestRepository is a mock implementation of TestRepository
type MockTestRepository struct {
	mock.Mock
}

func (m *MockTestRepository) CreateTest(ctx context.Context, params knowledgedb.CreateTestParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockTestRepository) GetTestById(ctx context.Context, id int64) (knowledgedb.Test, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(knowledgedb.Test), args.Error(1)
}

func (m *MockTestRepository) GetMyTests(ctx context.Context, createdBy int64) ([]knowledgedb.Test, error) {
	args := m.Called(ctx, createdBy)
	return args.Get(0).([]knowledgedb.Test), args.Error(1)
}

func (m *MockTestRepository) GetAvailableTests(ctx context.Context, createdBy int64) ([]knowledgedb.GetAvailableTestsRow, error) {
	args := m.Called(ctx, createdBy)
	return args.Get(0).([]knowledgedb.GetAvailableTestsRow), args.Error(1)
}

func (m *MockTestRepository) UpdateTest(ctx context.Context, params knowledgedb.UpdateTestParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockTestRepository) DeleteTest(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
