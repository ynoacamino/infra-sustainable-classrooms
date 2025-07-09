package ports

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TransactionManager maneja las transacciones de base de datos
type TransactionManager interface {
	WithTx(ctx context.Context, fn func(tx pgx.Tx) error) error
}
