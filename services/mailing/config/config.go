package config

import (
	"context"
	"fmt"
	"os"

	goaLog "goa.design/clue/log"
)

// Environment variable keys
const (
	HTTP_PORT = "HTTP_PORT"
	GRPC_PORT = "GRPC_PORT"
	DBG       = "DBG"
	APP_ENV   = "APP_ENV"

	// SMTP configuration
	SMTP_HOST     = "SMTP_HOST"
	SMTP_PORT     = "SMTP_PORT"
	SMTP_USERNAME = "SMTP_USERNAME"
	SMTP_PASSWORD = "SMTP_PASSWORD"
	SMTP_FROM     = "SMTP_FROM"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Config struct {
	// Server configuration
	HTTPPort    string
	GRPCPort    string
	Debug       bool
	Environment string
	Ctx         context.Context

	// SMTP configuration
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string
}

func NewConfig() (*Config, error) {
	// Determine environment
	environment := getEnvOrDefault(APP_ENV, "development")

	// Required configurations
	smtpHost := os.Getenv(SMTP_HOST)
	if smtpHost == "" {
		return nil, fmt.Errorf("environment variable %s not found", SMTP_HOST)
	}

	smtpUsername := os.Getenv(SMTP_USERNAME)
	if smtpUsername == "" {
		return nil, fmt.Errorf("environment variable %s not found", SMTP_USERNAME)
	}

	smtpPassword := os.Getenv(SMTP_PASSWORD)
	if smtpPassword == "" {
		return nil, fmt.Errorf("environment variable %s not found", SMTP_PASSWORD)
	}

	smtpFrom := os.Getenv(SMTP_FROM)
	if smtpFrom == "" {
		return nil, fmt.Errorf("environment variable %s not found", SMTP_FROM)
	}

	// Optional configurations with defaults
	httpPort := getEnvOrDefault(HTTP_PORT, "8080")
	grpcPort := getEnvOrDefault(GRPC_PORT, "9090")
	smtpPort := parseIntOrDefault(SMTP_PORT, 587)

	// Parse boolean values
	debug := parseBoolOrDefault(DBG, false)

	// Setup logging context
	format := goaLog.FormatTerminal
	ctx := goaLog.Context(context.Background(), goaLog.WithFormat(format))

	if debug {
		ctx = goaLog.Context(ctx, goaLog.WithDebug())
		goaLog.Debugf(ctx, "debug logs enabled")
	}

	// Log important configuration
	goaLog.Print(ctx, goaLog.KV{K: "http-port", V: httpPort})
	goaLog.Print(ctx, goaLog.KV{K: "grpc-port", V: grpcPort})
	goaLog.Print(ctx, goaLog.KV{K: "environment", V: environment})
	goaLog.Print(ctx, goaLog.KV{K: "smtp-host", V: smtpHost})
	goaLog.Print(ctx, goaLog.KV{K: "smtp-port", V: smtpPort})

	return &Config{
		HTTPPort:     httpPort,
		GRPCPort:     grpcPort,
		Debug:        debug,
		Environment:  environment,
		Ctx:          ctx,
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
		SMTPFrom:     smtpFrom,
	}, nil
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) GetSMTPConfig() *SMTPConfig {
	return &SMTPConfig{
		Host:     c.SMTPHost,
		Port:     c.SMTPPort,
		Username: c.SMTPUsername,
		Password: c.SMTPPassword,
		From:     c.SMTPFrom,
	}
}
