package mocks

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/repositories"
)

// MockBackupCodeRepository is a mock implementation of BackupCodeRepository
type MockBackupCodeRepository struct {
	mock.Mock
}

func (m *MockBackupCodeRepository) CreateBackupCodes(ctx context.Context, params []authdb.CreateBackupCodesParams) (int64, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBackupCodeRepository) GetUserBackupCodes(ctx context.Context, userID int64) ([]authdb.BackupCode, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]authdb.BackupCode), args.Error(1)
}

func (m *MockBackupCodeRepository) GetBackupCodeByHash(ctx context.Context, params authdb.GetBackupCodeByHashParams) (authdb.BackupCode, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.BackupCode), args.Error(1)
}

func (m *MockBackupCodeRepository) UseBackupCode(ctx context.Context, params authdb.UseBackupCodeParams) (authdb.BackupCode, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(authdb.BackupCode), args.Error(1)
}

func (m *MockBackupCodeRepository) GetUsedBackupCodes(ctx context.Context, userID int64) ([]authdb.BackupCode, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]authdb.BackupCode), args.Error(1)
}

func (m *MockBackupCodeRepository) CountAvailableBackupCodes(ctx context.Context, userID int64) (int64, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBackupCodeRepository) DeleteUserBackupCodes(ctx context.Context, userID int64) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockBackupCodeRepository) DeleteUsedBackupCodes(ctx context.Context, userID int64) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

// MockTransactionManager is a mock implementation of TransactionManager
type MockTransactionManager struct {
	mock.Mock
}

func (m *MockTransactionManager) WithTx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}

// MockRepositoryManager is a mock implementation of RepositoryManager
type MockRepositoryManager struct {
	mock.Mock
	UserRepo       *MockUserRepository
	SessionRepo    *MockSessionRepository
	BackupCodeRepo *MockBackupCodeRepository
}

func (m *MockRepositoryManager) WithTransaction(tx pgx.Tx) *repositories.TransactionalRepositories {
	args := m.Called(tx)
	return args.Get(0).(*repositories.TransactionalRepositories)
}
