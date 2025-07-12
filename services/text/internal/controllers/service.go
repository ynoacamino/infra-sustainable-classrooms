package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/gen/text"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/text/internal/repositories"
)

type textsrvc struct {
	articleRepo ports.ArticleRepository
	sectionRepo ports.SectionRepository
	courseRepo  ports.CourseRepository
	profilesServiceRepo ports.ProfilesServiceRepository
}

func NewText(
	repoManager *repositories.RepositoryManager,
) text.Service {
	return &textsrvc{
		articleRepo:         repoManager.ArticleRepo,
		sectionRepo:         repoManager.SectionRepo,
		courseRepo:          repoManager.CourseRepo,
		profilesServiceRepo: repoManager.ProfilesServiceRepo,
	}
}
