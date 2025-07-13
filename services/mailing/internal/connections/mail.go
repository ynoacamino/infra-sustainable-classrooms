package connections

import (
	"fmt"

	"github.com/wneessen/go-mail"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/mailing/config"
)

// ConnectSMTP creates a new SMTP client with the provided configuration
func ConnectSMTP(cfg *config.Config) (*mail.Client, error) {
	smtpConfig := cfg.GetSMTPConfig()

	client, err := mail.NewClient(smtpConfig.Host,
		mail.WithPort(smtpConfig.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(smtpConfig.Username),
		mail.WithPassword(smtpConfig.Password),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create SMTP client: %w", err)
	}

	return client, nil
}
