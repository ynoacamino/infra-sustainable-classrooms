package repositories

import (
	"github.com/wneessen/go-mail"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/internal/ports"
)

// RepositoryManager manages all repositories for the mailing service
type RepositoryManager struct {
	Mailing ports.MailingRepository
}

// NewRepositoryManager creates a new repository manager
func NewRepositoryManager(smtpClient *mail.Client, fromAddress string, prod bool) *RepositoryManager {
	return &RepositoryManager{
		Mailing: NewMailingRepository(smtpClient, fromAddress, prod),
	}
}

// Close closes any resources held by the repository manager
func (rm *RepositoryManager) Close() {
	// Since go-mail client doesn't need explicit closing, this is a no-op
	// but we keep it for consistency with other services
}
