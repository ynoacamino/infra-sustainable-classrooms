package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/gen/stats"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/repositories"
)

type statssrvc struct {
	textServiceRepo     ports.TextServiceRepository
	profilesServiceRepo ports.ProfilesServiceRepository
}

func NewStats(
	repoManager *repositories.RepositoryManager,
) stats.Service {
	return &statssrvc{
		textServiceRepo:     repoManager.TextServiceRepo,
		profilesServiceRepo: repoManager.ProfilesServiceRepo,
	}
}
