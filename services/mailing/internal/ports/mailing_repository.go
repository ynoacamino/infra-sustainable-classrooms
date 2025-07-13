package ports

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/gen/mailing"
)

// MailingRepository defines the interface for email operations
type MailingRepository interface {
	// SendEmail sends an email message via SMTP
	SendEmail(ctx context.Context, email *mailing.EmailMessage) (*mailing.EmailResponse, error)
}
