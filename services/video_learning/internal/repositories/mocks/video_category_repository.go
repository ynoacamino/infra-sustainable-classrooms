package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// MockVideoCategoryRepository is a mock implementation of VideoCategoryRepository
type MockVideoCategoryRepository struct {
	mock.Mock
}

func (m *MockVideoCategoryRepository) GetOrCreateCategory(ctx context.Context, name string) (videolearningdb.VideoCategory, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(videolearningdb.VideoCategory), args.Error(1)
}

func (m *MockVideoCategoryRepository) GetAllCategories(ctx context.Context) ([]videolearningdb.VideoCategory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]videolearningdb.VideoCategory), args.Error(1)
}

func (m *MockVideoCategoryRepository) GetCategoryById(ctx context.Context, id int64) (videolearningdb.VideoCategory, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(videolearningdb.VideoCategory), args.Error(1)
}
