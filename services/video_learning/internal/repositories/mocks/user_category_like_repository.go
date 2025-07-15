package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	videolearningdb "github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/gen/database"
)

// MockUserCategoryLikeRepository is a mock implementation of UserCategoryLikeRepository
type MockUserCategoryLikeRepository struct {
	mock.Mock
}

func (m *MockUserCategoryLikeRepository) GetUserCategoryLikes(ctx context.Context, userID int64) ([]videolearningdb.GetUserCategoryLikesRow, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]videolearningdb.GetUserCategoryLikesRow), args.Error(1)
}

func (m *MockUserCategoryLikeRepository) UpsertUserCategoryLike(ctx context.Context, params videolearningdb.UpsertUserCategoryLikeParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockUserCategoryLikeRepository) DeleteAllUserCategoryLikes(ctx context.Context, userID int64) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockUserCategoryLikeRepository) UpsertUserVideoLike(ctx context.Context, params videolearningdb.UpsertUserVideoLikeParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockUserCategoryLikeRepository) GetUserVideoLike(ctx context.Context, params videolearningdb.GetUserVideoLikeParams) (videolearningdb.UserVideoLike, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(videolearningdb.UserVideoLike), args.Error(1)
}

func (m *MockUserCategoryLikeRepository) IncrementUserCategoryLike(ctx context.Context, params videolearningdb.IncrementUserCategoryLikeParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}
