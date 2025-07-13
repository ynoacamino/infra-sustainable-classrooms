package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type MockCourseRepository struct {
	mock.Mock
}

func (m *MockCourseRepository) CreateCourse(ctx context.Context, params textdb.CreateCourseParams) (error) {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockCourseRepository) GetCourse(ctx context.Context, id int64) (textdb.Course, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(textdb.Course), args.Error(1)
}

func (m *MockCourseRepository) ListCourses(ctx context.Context) ([]textdb.Course, error) {
	args := m.Called(ctx)
	return args.Get(0).([]textdb.Course), args.Error(1)
}

func (m *MockCourseRepository) DeleteCourse(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCourseRepository) UpdateCourse(ctx context.Context, params textdb.UpdateCourseParams) (error) {
	args := m.Called(ctx, params)
	return args.Error(0)
}
