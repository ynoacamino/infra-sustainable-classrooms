package repositories

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/ports"
)

// backupCodeRepository implementa BackupCodeRepository
type backupCodeRepository struct {
	*authdb.Queries
}

// NewBackupCodeRepository crea una nueva instancia del repositorio de códigos de respaldo
func NewBackupCodeRepository(pool *pgxpool.Pool) ports.BackupCodeRepository {
	return &backupCodeRepository{
		Queries: authdb.New(pool),
	}
}

// NewBackupCodeRepositoryWithTx crea una nueva instancia del repositorio de códigos de respaldo con transacción
func NewBackupCodeRepositoryWithTx(tx pgx.Tx) ports.BackupCodeRepository {
	return &backupCodeRepository{
		Queries: authdb.New(tx),
	}
}
