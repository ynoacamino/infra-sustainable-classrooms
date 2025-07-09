package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
)

// MockSessionRepository is a mock implementation of SessionRepository
type MockSessionRepository struct {
	mock.Mock
}

func (m *MockSessionRepository) CreateSession(ctx context.Context, params authdb.CreateSessionParams) (authdb.Session, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.Session), args.Error(1)
}

func (m *MockSessionRepository) GetSessionByToken(ctx context.Context, sessionToken string) (authdb.GetSessionByTokenRow, error) {
	args := m.Called(ctx, sessionToken)
	return args.Get(0).(authdb.GetSessionByTokenRow), args.Error(1)
}

func (m *MockSessionRepository) GetUserSessions(ctx context.Context, userID int64) ([]authdb.Session, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]authdb.Session), args.Error(1)
}

func (m *MockSessionRepository) UpdateSessionAccess(ctx context.Context, sessionToken string) error {
	args := m.Called(ctx, sessionToken)
	return args.Error(0)
}

func (m *MockSessionRepository) RefreshSession(ctx context.Context, params authdb.RefreshSessionParams) (authdb.Session, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.Session), args.Error(1)
}

func (m *MockSessionRepository) DeactivateSession(ctx context.Context, sessionToken string) error {
	args := m.Called(ctx, sessionToken)
	return args.Error(0)
}

func (m *MockSessionRepository) DeactivateUserSessions(ctx context.Context, userID int64) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockSessionRepository) DeactivateAllUserSessionsExcept(ctx context.Context, params authdb.DeactivateAllUserSessionsExceptParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockSessionRepository) CleanupExpiredSessions(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockSessionRepository) GetSessionStats(ctx context.Context) (authdb.GetSessionStatsRow, error) {
	args := m.Called(ctx)
	return args.Get(0).(authdb.GetSessionStatsRow), args.Error(1)
}
