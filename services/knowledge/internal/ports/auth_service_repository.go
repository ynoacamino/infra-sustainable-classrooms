package ports

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
)

type AuthServiceRepository interface {
	ValidateUser(ctx context.Context, in *auth.ValidateUserPayload) (*auth.UserValidationResponse, error)
	GetUserByID(ctx context.Context, in *auth.GetUserByIDPayload) (*auth.User, error)
}
