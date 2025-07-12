package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// transactionManager implementa TransactionManager
type TransactionManager struct {
	pool *pgxpool.Pool
}

// NewTransactionManager crea una nueva instancia del manejador de transacciones
func NewTransactionManager(pool *pgxpool.Pool) TransactionManager {
	return TransactionManager{
		pool: pool,
	}
}

func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
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
