package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/gen/codelab"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/codelab/internal/repositories"
)

type codelabsvrc struct {
	profilesServiceRepo ports.ProfilesServiceRepository
	answersRepo         ports.AnswersRepository
	exercisesRepo       ports.ExercisesRepository
	testsRepo           ports.TestsRepository
	attemptsRepo        ports.AttemptsRepository
}

func NewCodelab(repoManager *repositories.RepositoryManager) codelab.Service {
	return &codelabsvrc{
		profilesServiceRepo: repoManager.ProfilesServiceRepo,
		answersRepo:         repoManager.AnswersRepo,
		exercisesRepo:       repoManager.ExercisesRepo,
		testsRepo:           repoManager.TestsRepo,
		attemptsRepo:        repoManager.AttemptsRepo,
	}
}
