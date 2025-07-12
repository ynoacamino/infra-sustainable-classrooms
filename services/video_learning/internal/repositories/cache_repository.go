package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/internal/ports"
)

type cacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) ports.CacheRepository {
	return &cacheRepository{
		client: client,
	}
}

func (r *cacheRepository) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, string(valueBytes), expiration).Err()
}

func (r *cacheRepository) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *cacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *cacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r *cacheRepository) SetNX(ctx context.Context, key string, value any, expiration time.Duration) (bool, error) {
	// Convert value to JSON string for storage
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	result, err := r.client.SetNX(ctx, key, string(valueBytes), expiration).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}
