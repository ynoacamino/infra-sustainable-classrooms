package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/gen/stats"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/stats/internal/ports"
)

type statssrvc struct {
	textServiceRepo     ports.TextServiceRepository
	profilesServiceRepo ports.ProfilesServiceRepository
}

func NewStatsService(
	textServiceRepo ports.TextServiceRepository,
	profilesServiceRepo ports.ProfilesServiceRepository,
) stats.Service {
	return &statssrvc{
		textServiceRepo:     textServiceRepo,
		profilesServiceRepo: profilesServiceRepo,
	}
}