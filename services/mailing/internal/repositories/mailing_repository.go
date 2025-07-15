package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/wneessen/go-mail"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/gen/mailing"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/internal/ports"
)

// mailingRepository implements the MailingRepository interface
type mailingRepository struct {
	client *mail.Client
	from   string
	prod   bool
}

// NewMailingRepository creates a new mailing repository
func NewMailingRepository(client *mail.Client, from string, prod bool) ports.MailingRepository {
	return &mailingRepository{
		client: client,
		from:   from,
		prod:   prod,
	}
}

// SendEmail sends an email message via SMTP
func (r *mailingRepository) SendEmail(ctx context.Context, email *mailing.EmailMessage) (*mailing.EmailResponse, error) {
	if !r.prod {
		fmt.Printf("Sending email to: %v, subject: %s\n", email.To, email.Subject)
		return nil, fmt.Errorf("development mode, not sending email")
	}

	// Create new message
	msg := mail.NewMsg()

	// Set from address
	if err := msg.From(r.from); err != nil {
		return nil, fmt.Errorf("failed to set from address: %w", err)
	}

	// Set recipients
	if err := msg.To(email.To...); err != nil {
		return nil, fmt.Errorf("failed to set to addresses: %w", err)
	}

	// Set CC if provided
	if len(email.Cc) > 0 {
		if err := msg.Cc(email.Cc...); err != nil {
			return nil, fmt.Errorf("failed to set cc addresses: %w", err)
		}
	}

	// Set BCC if provided
	if len(email.Bcc) > 0 {
		if err := msg.Bcc(email.Bcc...); err != nil {
			return nil, fmt.Errorf("failed to set bcc addresses: %w", err)
		}
	}

	// Set subject
	msg.Subject(email.Subject)

	// Set body based on content type
	if email.IsHTML {
		msg.SetBodyString(mail.TypeTextHTML, email.Body)
	} else {
		msg.SetBodyString(mail.TypeTextPlain, email.Body)
	}

	// Send the email
	if err := r.client.DialAndSend(msg); err != nil {
		return &mailing.EmailResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to send email: %v", err),
		}, nil
	}

	// Get message ID if available
	messageID := msg.GetGenHeader(mail.HeaderMessageID)
	var msgIDString string
	if len(messageID) > 0 {
		msgIDString = strings.Join(messageID, ",")
	}

	return &mailing.EmailResponse{
		Success:   true,
		Message:   "Email sent successfully",
		MessageID: &msgIDString,
	}, nil
}
