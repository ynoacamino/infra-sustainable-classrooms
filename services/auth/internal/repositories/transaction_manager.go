package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/ports"
)

// transactionManager implementa TransactionManager
type transactionManager struct {
	pool *pgxpool.Pool
}

// NewTransactionManager crea una nueva instancia del manejador de transacciones
func NewTransactionManager(pool *pgxpool.Pool) ports.TransactionManager {
	return &transactionManager{
		pool: pool,
	}
}

func (tm *transactionManager) WithTx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := tm.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := fn(tx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
