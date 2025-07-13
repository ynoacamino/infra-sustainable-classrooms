package mocks

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// MockVideoRepository is a mock implementation of VideoRepository
type MockVideoRepository struct {
	mock.Mock
}

func (m *MockVideoRepository) CreateVideo(ctx context.Context, params videolearningdb.CreateVideoParams) (videolearningdb.Video, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(videolearningdb.Video), args.Error(1)
}

func (m *MockVideoRepository) GetVideoByID(ctx context.Context, id int64) (videolearningdb.GetVideoByIDRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(videolearningdb.GetVideoByIDRow), args.Error(1)
}

func (m *MockVideoRepository) DeleteVideo(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockVideoRepository) SearchVideos(ctx context.Context, params videolearningdb.SearchVideosParams) ([]videolearningdb.SearchVideosRow, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]videolearningdb.SearchVideosRow), args.Error(1)
}

func (m *MockVideoRepository) GetVideosByCategory(ctx context.Context, id int64) ([]videolearningdb.GetVideosByCategoryRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]videolearningdb.GetVideosByCategoryRow), args.Error(1)
}

func (m *MockVideoRepository) GetVideosByUser(ctx context.Context, params videolearningdb.GetVideosByUserParams) ([]videolearningdb.GetVideosByUserRow, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]videolearningdb.GetVideosByUserRow), args.Error(1)
}

func (m *MockVideoRepository) GetSimilarVideos(ctx context.Context, id int64) ([]videolearningdb.GetSimilarVideosRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]videolearningdb.GetSimilarVideosRow), args.Error(1)
}

func (m *MockVideoRepository) GetRecentVideos(ctx context.Context, interval pgtype.Interval) ([]videolearningdb.GetRecentVideosRow, error) {
	args := m.Called(ctx, interval)
	return args.Get(0).([]videolearningdb.GetRecentVideosRow), args.Error(1)
}

func (m *MockVideoRepository) IncrementVideoViews(ctx context.Context, params videolearningdb.IncrementVideoViewsParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockVideoRepository) IncrementVideoLikes(ctx context.Context, params videolearningdb.IncrementVideoLikesParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockVideoRepository) AssignTagToVideo(ctx context.Context, params videolearningdb.AssignTagToVideoParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockVideoRepository) RemoveTagFromVideo(ctx context.Context, params videolearningdb.RemoveTagFromVideoParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}
