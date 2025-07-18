package config

import (
	"context"
	"fmt"
	"os"

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

	PROFILES_GRPC_ADDRESS = "PROFILES_GRPC_ADDRESS"

	MINIO_ENDPOINT   = "MINIO_ENDPOINT"
	MINIO_ACCESS_KEY = "MINIO_ACCESS_KEY"
	MINIO_SECRET_KEY = "MINIO_SECRET_KEY"

	REDIS_ENDPOINT = "REDIS_ENDPOINT"

	// Aggregation configuration
	AGGREGATION_INTERVAL_SECONDS = "AGGREGATION_INTERVAL_SECONDS"
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

type MinioConfig struct {
	Ctx       context.Context
	Endpoint  string
	AccessKey string
	SecretKey string
}

type RedisConfig struct {
	Ctx      context.Context
	Endpoint string
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
	ProfilesGRPCAddress string

	// MinIO configuration
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string

	// Redis configuration
	RedisEndpoint string

	// Aggregation configuration
	AggregationIntervalSeconds int
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
	profilesGRPCAddress := getEnvOrDefault(PROFILES_GRPC_ADDRESS, fmt.Sprintf("profiles-service:%s", grpcPort))

	// Parse boolean and numeric values
	debug := parseBoolOrDefault(DBG, false)

	// Setup logging context
	format := goaLog.FormatTerminal

	ctx := goaLog.Context(context.Background(), goaLog.WithFormat(format))

	if debug {
		ctx = goaLog.Context(ctx, goaLog.WithDebug())
		goaLog.Debugf(ctx, "debug logs enabled")
	}

	// MinIO configuration
	minioEndpoint := getEnvOrDefault(MINIO_ENDPOINT, "minio:9000")
	minioAccessKey := getEnvOrDefault(MINIO_ACCESS_KEY, "minioadmin")
	minioSecretKey := getEnvOrDefault(MINIO_SECRET_KEY, "minioadmin")

	if minioEndpoint == "" || minioAccessKey == "" || minioSecretKey == "" {
		return nil, fmt.Errorf("MinIO configuration is incomplete")
	}

	// Redis configuration
	redisEndpoint := getEnvOrDefault(REDIS_ENDPOINT, "redis:6379")
	if redisEndpoint == "" {
		return nil, fmt.Errorf("redis configuration is incomplete")
	}

	// Aggregation configuration
	aggregationIntervalSeconds := parseIntOrDefault(AGGREGATION_INTERVAL_SECONDS, 300)

	// Log important configuration
	goaLog.Print(ctx, goaLog.KV{K: "http-port", V: httpPort})
	goaLog.Print(ctx, goaLog.KV{K: "grpc-port", V: grpcPort})
	goaLog.Print(ctx, goaLog.KV{K: "environment", V: environment})

	return &Config{
		DatabaseURL:                databaseURL,
		HTTPPort:                   httpPort,
		GRPCPort:                   grpcPort,
		Debug:                      debug,
		Environment:                environment,
		Ctx:                        ctx,
		MaxConns:                   max_conns,
		MinConns:                   min_conns,
		ProfilesGRPCAddress:        profilesGRPCAddress,
		MinioEndpoint:              minioEndpoint,
		MinioAccessKey:             minioAccessKey,
		MinioSecretKey:             minioSecretKey,
		RedisEndpoint:              redisEndpoint,
		AggregationIntervalSeconds: aggregationIntervalSeconds,
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

func (c *Config) GetMinioConfig() *MinioConfig {
	return &MinioConfig{
		Ctx:       c.Ctx,
		Endpoint:  c.MinioEndpoint,
		AccessKey: c.MinioAccessKey,
		SecretKey: c.MinioSecretKey,
	}
}

func (c *Config) GetRedisConfig() *RedisConfig {
	return &RedisConfig{
		Ctx:      c.Ctx,
		Endpoint: c.RedisEndpoint,
	}
}
