package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
)

// MockAuthServiceRepository is a mock implementation of AuthServiceRepository
type MockAuthServiceRepository struct {
	mock.Mock
}

func (m *MockAuthServiceRepository) ValidateUser(ctx context.Context, in *auth.ValidateUserPayload) (*auth.UserValidationResponse, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.UserValidationResponse), args.Error(1)
}

func (m *MockAuthServiceRepository) GetUserByID(ctx context.Context, in *auth.GetUserByIDPayload) (*auth.User, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.User), args.Error(1)
}
