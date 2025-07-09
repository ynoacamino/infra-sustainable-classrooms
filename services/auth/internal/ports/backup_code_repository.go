package ports

import (
	"context"

	authdb "github.com/ynoacamino/infrastructure/services/auth/gen/database"
)

// BackupCodeRepository define las operaciones de persistencia para códigos de respaldo
type BackupCodeRepository interface {
	CreateBackupCodes(ctx context.Context, params []authdb.CreateBackupCodesParams) (int64, error)
	GetUserBackupCodes(ctx context.Context, userID int64) ([]authdb.BackupCode, error)
	GetBackupCodeByHash(ctx context.Context, params authdb.GetBackupCodeByHashParams) (authdb.BackupCode, error)
	UseBackupCode(ctx context.Context, params authdb.UseBackupCodeParams) (authdb.BackupCode, error)
	GetUsedBackupCodes(ctx context.Context, userID int64) ([]authdb.BackupCode, error)
	CountAvailableBackupCodes(ctx context.Context, userID int64) (int64, error)
	DeleteUserBackupCodes(ctx context.Context, userID int64) error
	DeleteUsedBackupCodes(ctx context.Context, userID int64) error
}
