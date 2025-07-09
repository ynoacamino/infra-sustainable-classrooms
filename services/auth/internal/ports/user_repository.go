package ports

import (
	"context"

	authdb "github.com/ynoacamino/infrastructure/services/auth/gen/database"
)

// UserRepository define las operaciones de persistencia para usuarios
type UserRepository interface {
	CreateUser(ctx context.Context, params authdb.CreateUserParams) (authdb.User, error)
	GetUserByID(ctx context.Context, id int64) (authdb.User, error)
	GetUserByIdentifier(ctx context.Context, identifier string) (authdb.User, error)
	UpdateUserTOTPSecret(ctx context.Context, params authdb.UpdateUserTOTPSecretParams) (authdb.User, error)
	VerifyUser(ctx context.Context, id int64) (authdb.User, error)
	UpdateUserLastLogin(ctx context.Context, id int64) error
	UpdateUserMetadata(ctx context.Context, params authdb.UpdateUserMetadataParams) (authdb.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserStats(ctx context.Context) (authdb.GetUserStatsRow, error)
}
