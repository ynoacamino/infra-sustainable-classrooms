package repositories

import (
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/ports"
)

// SessionRepository implementa ports.SessionRepository usando embedding
type SessionRepository struct {
	*authdb.Queries
}

// NewSessionRepository crea una nueva instancia del repositorio de sesiones
func NewSessionRepository(db authdb.DBTX) ports.SessionRepository {
	return &SessionRepository{
		Queries: authdb.New(db),
	}
}

// NewSessionRepositoryWithTx crea una nueva instancia del repositorio de sesiones con transacción
func NewSessionRepositoryWithTx(tx authdb.DBTX) ports.SessionRepository {
	return &SessionRepository{
		Queries: authdb.New(tx),
	}
}
