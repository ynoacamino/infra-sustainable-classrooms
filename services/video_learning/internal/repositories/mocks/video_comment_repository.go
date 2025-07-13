package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// MockVideoCommentRepository is a mock implementation of VideoCommentRepository
type MockVideoCommentRepository struct {
	mock.Mock
}

func (m *MockVideoCommentRepository) CreateComment(ctx context.Context, params videolearningdb.CreateCommentParams) (videolearningdb.VideoComment, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(videolearningdb.VideoComment), args.Error(1)
}

func (m *MockVideoCommentRepository) GetCommentsForVideo(ctx context.Context, params videolearningdb.GetCommentsForVideoParams) ([]videolearningdb.VideoComment, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]videolearningdb.VideoComment), args.Error(1)
}

func (m *MockVideoCommentRepository) DeleteComment(ctx context.Context, params videolearningdb.DeleteCommentParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}
