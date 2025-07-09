package ports

import (
	"context"

	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
)

// SessionRepository define las operaciones de persistencia para sesiones
type SessionRepository interface {
	CreateSession(ctx context.Context, params authdb.CreateSessionParams) (authdb.Session, error)
	GetSessionByToken(ctx context.Context, sessionToken string) (authdb.GetSessionByTokenRow, error)
	GetUserSessions(ctx context.Context, userID int64) ([]authdb.Session, error)
	UpdateSessionAccess(ctx context.Context, sessionToken string) error
	RefreshSession(ctx context.Context, params authdb.RefreshSessionParams) (authdb.Session, error)
	DeactivateSession(ctx context.Context, sessionToken string) error
	DeactivateUserSessions(ctx context.Context, userID int64) error
	DeactivateAllUserSessionsExcept(ctx context.Context, params authdb.DeactivateAllUserSessionsExceptParams) error
	CleanupExpiredSessions(ctx context.Context) error
	GetSessionStats(ctx context.Context) (authdb.GetSessionStatsRow, error)
}
