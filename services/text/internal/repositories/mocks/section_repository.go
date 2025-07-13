package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type MockSectionRepository struct {
	mock.Mock
}

func (m *MockSectionRepository) CreateSection(ctx context.Context, params textdb.CreateSectionParams) (error) {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockSectionRepository) GetSection(ctx context.Context, id int64) (textdb.Section, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(textdb.Section), args.Error(1)
}

func (m *MockSectionRepository) ListSectionsByCourse(ctx context.Context, courseID int64) ([]textdb.Section, error) {
	args := m.Called(ctx, courseID)
	return args.Get(0).([]textdb.Section), args.Error(1)
}

func (m *MockSectionRepository) DeleteSection(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSectionRepository) UpdateSection(ctx context.Context, params textdb.UpdateSectionParams) (error) {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockSectionRepository) GetNextOrderForCourse(ctx context.Context, courseID int64) (int32, error) {
	args := m.Called(ctx, courseID)
	return args.Get(0).(int32), args.Error(1)
}