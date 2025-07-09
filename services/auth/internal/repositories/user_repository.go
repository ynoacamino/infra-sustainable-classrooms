package repositories

import (
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/ports"
)

// UserRepository implementa ports.UserRepository usando embedding directo
type UserRepository struct {
	*authdb.Queries
}

// NewUserRepository crea una nueva instancia del repositorio de usuarios
func NewUserRepository(db authdb.DBTX) ports.UserRepository {
	return &UserRepository{
		Queries: authdb.New(db),
	}
}

// NewUserRepositoryWithTx crea una nueva instancia del repositorio de usuarios con transacción
func NewUserRepositoryWithTx(tx authdb.DBTX) ports.UserRepository {
	return &UserRepository{
		Queries: authdb.New(tx),
	}
}
