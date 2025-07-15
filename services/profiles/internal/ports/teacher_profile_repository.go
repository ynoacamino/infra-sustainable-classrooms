package ports

import (
	"context"

	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

type TeacherProfileRepository interface {
	CreateTeacherProfile(ctx context.Context, params profilesdb.CreateTeacherProfileParams) (profilesdb.TeacherProfile, error)
	GetTeacherProfileByProfileId(ctx context.Context, profileID int64) (profilesdb.TeacherProfile, error)
	GetTeacherProfileByUserId(ctx context.Context, userID int64) (profilesdb.TeacherProfile, error)
	UpdateTeacherProfile(ctx context.Context, params profilesdb.UpdateTeacherProfileParams) (profilesdb.TeacherProfile, error)
	GetCompleteTeacherProfile(ctx context.Context, userID int64) (profilesdb.GetCompleteTeacherProfileRow, error)
}
