package repositories

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infrastructure/services/auth/internal/ports"
)

// RepositoryManager coordina todos los repositorios del dominio auth
type RepositoryManager struct {
	UserRepo       ports.UserRepository
	SessionRepo    ports.SessionRepository
	BackupCodeRepo ports.BackupCodeRepository
	TxManager      ports.TransactionManager
	pool           *pgxpool.Pool
}

// NewRepositoryManager crea una nueva instancia del manejador de repositorios
func NewRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		UserRepo:       NewUserRepository(pool),
		SessionRepo:    NewSessionRepository(pool),
		BackupCodeRepo: NewBackupCodeRepository(pool),
		TxManager:      NewTransactionManager(pool),
		pool:           pool,
	}
}

// Close cierra la conexión del pool
func (rm *RepositoryManager) Close() {
	rm.pool.Close()
}

func (rm *RepositoryManager) WithTransaction(tx pgx.Tx) *TransactionalRepositories {
	return &TransactionalRepositories{
		UserRepo:       rm.NewUserRepositoryWithTx(tx),
		SessionRepo:    rm.NewSessionRepositoryWithTx(tx),
		BackupCodeRepo: rm.NewBackupCodeRepositoryWithTx(tx),
	}
}

type TransactionalRepositories struct {
	UserRepo       ports.UserRepository
	SessionRepo    ports.SessionRepository
	BackupCodeRepo ports.BackupCodeRepository
}

func (rm *RepositoryManager) NewUserRepositoryWithTx(tx pgx.Tx) ports.UserRepository {
	return NewUserRepositoryWithTx(tx)
}

func (rm *RepositoryManager) NewSessionRepositoryWithTx(tx pgx.Tx) ports.SessionRepository {
	return NewSessionRepositoryWithTx(tx)
}

func (rm *RepositoryManager) NewBackupCodeRepositoryWithTx(tx pgx.Tx) ports.BackupCodeRepository {
	return NewBackupCodeRepositoryWithTx(tx)
}
