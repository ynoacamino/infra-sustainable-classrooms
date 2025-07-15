package ports

import (
	"context"

	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, params profilesdb.CreateProfileParams) (profilesdb.Profile, error)
	GetProfileByUserId(ctx context.Context, userID int64) (profilesdb.Profile, error)
	GetProfileByEmail(ctx context.Context, email string) (profilesdb.Profile, error)
	GetPublicProfileByUserId(ctx context.Context, userID int64) (profilesdb.GetPublicProfileByUserIdRow, error)
	UpdateProfile(ctx context.Context, params profilesdb.UpdateProfileParams) (profilesdb.Profile, error)
	DeactivateProfile(ctx context.Context, userID int64) error
	CheckProfileExists(ctx context.Context, userID int64) (bool, error)
	GetProfileRole(ctx context.Context, userID int64) (string, error)
}
