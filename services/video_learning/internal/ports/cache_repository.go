package ports

import (
	"context"
	"time"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error)

	// New methods for batch operations and scanning
	Scan(ctx context.Context, pattern string) ([]string, error)
	MGet(ctx context.Context, keys []string) ([]string, error)
	MSet(ctx context.Context, keyValues map[string]any, expiration time.Duration) error
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	GetInt(ctx context.Context, key string) (int64, error)
}
