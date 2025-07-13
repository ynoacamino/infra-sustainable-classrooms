package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// MockVideoTagRepository is a mock implementation of VideoTagRepository
type MockVideoTagRepository struct {
	mock.Mock
}

func (m *MockVideoTagRepository) GetOrCreateTag(ctx context.Context, name string) (videolearningdb.VideoTag, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(videolearningdb.VideoTag), args.Error(1)
}

func (m *MockVideoTagRepository) GetAllTags(ctx context.Context) ([]videolearningdb.VideoTag, error) {
	args := m.Called(ctx)
	return args.Get(0).([]videolearningdb.VideoTag), args.Error(1)
}

func (m *MockVideoTagRepository) GetTagsByVideoID(ctx context.Context, videoID int64) ([]videolearningdb.VideoTag, error) {
	args := m.Called(ctx, videoID)
	return args.Get(0).([]videolearningdb.VideoTag), args.Error(1)
}

func (m *MockVideoTagRepository) GetTagByName(ctx context.Context, name string) (videolearningdb.VideoTag, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(videolearningdb.VideoTag), args.Error(1)
}
