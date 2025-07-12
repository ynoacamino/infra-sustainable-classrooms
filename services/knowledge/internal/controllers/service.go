package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/knowledge"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/repositories"
)

type knowledgesvrc struct {
	authServiceRepo     ports.AuthServiceRepository
	profilesServiceRepo ports.ProfilesServiceRepository
	questionRepo        ports.QuestionRepository
	submissionRepo      ports.SubmissionRepository
	testRepo            ports.TestRepository
	answerRepo          ports.AnswerRepository
}

func NewKnowledge(
	repoManager *repositories.RepositoryManager,
) knowledge.Service {
	return &knowledgesvrc{
		authServiceRepo:     repoManager.AuthServiceRepo,
		profilesServiceRepo: repoManager.ProfilesServiceRepo,
		questionRepo:        repoManager.QuestionRepo,
		submissionRepo:      repoManager.SubmissionRepo,
		testRepo:            repoManager.TestRepo,
		answerRepo:          repoManager.AnswerRepo,
	}
}
