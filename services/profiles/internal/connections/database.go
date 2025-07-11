package connections

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/config"
)

type ConnectDBConfig interface {
	GetDatabaseConfig() *config.DBConfig
}

func ConnectDB(cfg ConnectDBConfig) (*pgxpool.Pool, error) {
	dbConfig := cfg.GetDatabaseConfig()
	ctx := dbConfig.Ctx

	poolConfig, err := pgxpool.ParseConfig(dbConfig.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Configure connection pool based on environment
	poolConfig.MinConns = int32(dbConfig.MinConns)
	poolConfig.MaxConns = int32(dbConfig.MaxConns)

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
