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

	AUTH_GRPC_ADDRESS = "AUTH_GRPC_ADDRESS"
)

type DBConfig struct {
	ConnectionString string
	Ctx              context.Context
	MaxConns         int
	MinConns         int
}

type ConnectGRPCConfig struct {
	GrpcAddress string
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

	// gRPC configuration
	AuthGRPCAddress string

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

	// gRPC configuration
	authGRPCAddress := getEnvOrDefault(AUTH_GRPC_ADDRESS, fmt.Sprintf("localhost:%s", grpcPort))

	// Parse boolean and numeric values
	debug := parseBoolOrDefault(DBG, false)

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

	return &Config{
		DatabaseURL:     databaseURL,
		HTTPPort:        httpPort,
		GRPCPort:        grpcPort,
		Debug:           debug,
		Environment:     environment,
		Ctx:             ctx,
		MaxConns:        max_conns,
		MinConns:        min_conns,
		AuthGRPCAddress: authGRPCAddress,
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
