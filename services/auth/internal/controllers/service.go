package controllers

import (
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/config"
	auth "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/ports"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/repositories"
)

type authsrvc struct {
	userRepo       ports.UserRepository
	sessionRepo    ports.SessionRepository
	backupCodeRepo ports.BackupCodeRepository
	txManager      ports.TransactionManager

	cfg         *config.Config
	repoManager *repositories.RepositoryManager
}

func NewAuth(repoManager *repositories.RepositoryManager, cfg *config.Config) auth.Service {
	return &authsrvc{
		userRepo:       repoManager.UserRepo,
		sessionRepo:    repoManager.SessionRepo,
		backupCodeRepo: repoManager.BackupCodeRepo,
		txManager:      repoManager.TxManager,
		repoManager:    repoManager,
		cfg:            cfg,
	}
}
