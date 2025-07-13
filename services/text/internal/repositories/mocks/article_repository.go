package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	textdb "github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/database"
)

type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) CreateArticle(ctx context.Context, params textdb.CreateArticleParams) (error) {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockArticleRepository) GetArticle(ctx context.Context, id int64) (textdb.Article, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(textdb.Article), args.Error(1)
}

func (m *MockArticleRepository) ListArticlesBySection(ctx context.Context, sectionID int64) ([]textdb.Article, error) {
	args := m.Called(ctx, sectionID)
	return args.Get(0).([]textdb.Article), args.Error(1)
}

func (m *MockArticleRepository) DeleteArticle(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockArticleRepository) UpdateArticle(ctx context.Context, params textdb.UpdateArticleParams) (error) {
	args := m.Called(ctx, params)
	return args.Error(0)
}