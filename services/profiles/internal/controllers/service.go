package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/internal/repositories"
)

type profilessrvc struct {
	profileRepo        ports.ProfileRepository
	studentProfileRepo ports.StudentProfileRepository
	teacherProfileRepo ports.TeacherProfileRepository
}

func NewProfilesService(
	repoManager *repositories.RepositoryManager,
) profiles.Service {
	return &profilessrvc{
		profileRepo:        repoManager.ProfileRepo,
		studentProfileRepo: repoManager.StudentProfileRepo,
		teacherProfileRepo: repoManager.TeacherProfileRepo,
	}
}
