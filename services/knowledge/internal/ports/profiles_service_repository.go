package ports

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

type ProfilesServiceRepository interface {
	GetCompleteProfile(ctx context.Context, in *profiles.GetCompleteProfilePayload) (*profiles.CompleteProfileResponse, error)
	GetPublicProfileByID(ctx context.Context, in *profiles.GetPublicProfileByIDPayload) (*profiles.PublicProfileResponse, error)
	ValidateUserRole(ctx context.Context, in *profiles.ValidateUserRolePayload) (*profiles.RoleValidationResponse, error)
}
