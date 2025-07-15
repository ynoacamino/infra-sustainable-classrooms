package controllers

import (
	"context"

	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/gen/mailing"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/internal/repositories"
	"goa.design/clue/log"
)

// mailingController implements the mailing service interface
type mailingController struct {
	repos *repositories.RepositoryManager
}

// NewMailing creates a new mailing service controller
func NewMailing(repos *repositories.RepositoryManager) mailing.Service {
	return &mailingController{
		repos: repos,
	}
}

// SendEmail handles sending an email message
func (s *mailingController) SendEmail(ctx context.Context, p *mailing.SendEmailPayload) (*mailing.EmailResponse, error) {
	log.Printf(ctx, "Sending email to: %v, subject: %s", p.Email.To, p.Email.Subject)

	result, err := s.repos.Mailing.SendEmail(ctx, p.Email)
	if err != nil {
		log.Errorf(ctx, err, "Failed to send email")
		return nil, mailing.EmailSendFailed(err.Error())
	}

	if !result.Success {
		log.Printf(ctx, "Email sending failed: %s", result.Message)
		return result, nil
	}

	log.Printf(ctx, "Email sent successfully: %s", result.Message)
	return result, nil
}
