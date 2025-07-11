package ports

import (
	"context"

	profilesdb "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/database"
)

type StudentProfileRepository interface {
	CreateStudentProfile(ctx context.Context, params profilesdb.CreateStudentProfileParams) (profilesdb.StudentProfile, error)
	GetStudentProfileByProfileId(ctx context.Context, profileID int64) (profilesdb.StudentProfile, error)
	GetStudentProfileByUserId(ctx context.Context, userID int64) (profilesdb.StudentProfile, error)
	UpdateStudentProfile(ctx context.Context, params profilesdb.UpdateStudentProfileParams) (profilesdb.StudentProfile, error)
	GetCompleteStudentProfile(ctx context.Context, userID int64) (profilesdb.GetCompleteStudentProfileRow, error)
}
