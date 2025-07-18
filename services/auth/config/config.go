package config

import (
	"context"
	"fmt"
	"os"
	"time"

	goaLog "goa.design/clue/log"
)

// Environment variable keys
const (
	DATABASE_URL = "DATABASE_URL"
	MAX_CONNS    = "MAX_CONNS"
	MIN_CONNS    = "MIN_CONNS"

	HTTP_PORT = "HTTP_PORT"
	GRPC_PORT = "GRPC_PORT"
	DBG       = "DBG"
	APP_ENV   = "APP_ENV"

	// Auth specific
	TOTP_ISSUER      = "TOTP_ISSUER"
	SESSION_DURATION = "SESSION_DURATION"

	// Security
	BACKUP_CODES_COUNT = "BACKUP_CODES_COUNT"
)

type DBConfig struct {
	ConnectionString string
	Ctx              context.Context
	MaxConns         int
	MinConns         int
}

type Config struct {
	// Server configuration
	DatabaseURL string
	MaxConns    int
	MinConns    int

	HTTPPort    string
	GRPCPort    string
	Debug       bool
	Environment string
	Ctx         context.Context

	TOTPIssuer      string
	SessionDuration time.Duration

	// Security
	BackupCodesCount int
}

func NewConfig() (*Config, error) {
	// Determine environment and env file
	environment := getEnvOrDefault(APP_ENV, "development")

	// Required configurations
	databaseURL := os.Getenv(DATABASE_URL)
	if databaseURL == "" {
		return nil, fmt.Errorf("environment variable %s not found", DATABASE_URL)
	}

	max_conns := parseIntOrDefault(MAX_CONNS, 500)
	min_conns := parseIntOrDefault(MIN_CONNS, 50)
	if max_conns < min_conns {
		return nil, fmt.Errorf("max connections (%d) cannot be less than min connections (%d)", max_conns, min_conns)
	}

	// Optional configurations with defaults
	httpPort := getEnvOrDefault(HTTP_PORT, "8080")
	grpcPort := getEnvOrDefault(GRPC_PORT, "9090")
	totpIssuer := getEnvOrDefault(TOTP_ISSUER, "Auth Service")

	// Parse boolean and numeric values
	debug := parseBoolOrDefault(DBG, false)
	backupCodesCount := parseIntOrDefault(BACKUP_CODES_COUNT, 10)

	// Parse durations
	sessionDuration := parseDurationOrDefault(SESSION_DURATION, "24h")

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
	goaLog.Print(ctx, goaLog.KV{K: "totp-issuer", V: totpIssuer})

	return &Config{
		DatabaseURL:      databaseURL,
		HTTPPort:         httpPort,
		GRPCPort:         grpcPort,
		Debug:            debug,
		Environment:      environment,
		Ctx:              ctx,
		TOTPIssuer:       totpIssuer,
		SessionDuration:  sessionDuration,
		BackupCodesCount: backupCodesCount,
		MaxConns:         max_conns,
		MinConns:         min_conns,
	}, nil
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) GetDatabaseConfig() *DBConfig {
	return &DBConfig{
		ConnectionString: c.DatabaseURL,
		Ctx:              c.Ctx,
		MaxConns:         c.MaxConns,
		MinConns:         c.MinConns,
	}
}
