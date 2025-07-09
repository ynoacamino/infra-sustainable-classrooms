package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, params authdb.CreateUserParams) (authdb.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int64) (authdb.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(authdb.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByIdentifier(ctx context.Context, identifier string) (authdb.User, error) {
	args := m.Called(ctx, identifier)
	return args.Get(0).(authdb.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserTOTPSecret(ctx context.Context, params authdb.UpdateUserTOTPSecretParams) (authdb.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.User), args.Error(1)
}

func (m *MockUserRepository) VerifyUser(ctx context.Context, id int64) (authdb.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(authdb.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUserLastLogin(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUserMetadata(ctx context.Context, params authdb.UpdateUserMetadataParams) (authdb.User, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserStats(ctx context.Context) (authdb.GetUserStatsRow, error) {
	args := m.Called(ctx)
	return args.Get(0).(authdb.GetUserStatsRow), args.Error(1)
}
