package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockCacheRepository is a mock implementation of CacheRepository
type MockCacheRepository struct {
	mock.Mock
}

func (m *MockCacheRepository) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockCacheRepository) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockCacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
}

func (m *MockCacheRepository) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *MockCacheRepository) Increment(ctx context.Context, key string, value int) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockCacheRepository) GetInt(ctx context.Context, key string) (int64, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCacheRepository) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	args := m.Called(ctx, key, value)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockCacheRepository) SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error) {
	args := m.Called(ctx, key, value, expiration)
	return args.Bool(0), args.Error(1)
}

func (m *MockCacheRepository) Scan(ctx context.Context, pattern string) ([]string, error) {
	args := m.Called(ctx, pattern)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockCacheRepository) MGet(ctx context.Context, keys []string) ([]string, error) {
	args := m.Called(ctx, keys)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockCacheRepository) MSet(ctx context.Context, keyValues map[string]any, expiration time.Duration) error {
	args := m.Called(ctx, keyValues, expiration)
	return args.Error(0)
}
