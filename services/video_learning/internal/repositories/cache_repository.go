package repositories

import (
	"context"
	"encoding/json"
	"strconv"
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

// Scan returns all keys matching the pattern
func (r *cacheRepository) Scan(ctx context.Context, pattern string) ([]string, error) {
	var keys []string
	iter := r.client.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return keys, nil
}

// MGet gets multiple keys at once
func (r *cacheRepository) MGet(ctx context.Context, keys []string) ([]string, error) {
	if len(keys) == 0 {
		return []string{}, nil
	}

	result, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	values := make([]string, len(result))
	for i, val := range result {
		if val != nil {
			values[i] = val.(string)
		} else {
			values[i] = ""
		}
	}

	return values, nil
}

// MSet sets multiple keys at once
func (r *cacheRepository) MSet(ctx context.Context, keyValues map[string]any, expiration time.Duration) error {
	pipe := r.client.Pipeline()

	for key, value := range keyValues {
		valueBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		pipe.Set(ctx, key, string(valueBytes), expiration)
	}

	_, err := pipe.Exec(ctx)
	return err
}

// IncrBy increments a key by the given value
func (r *cacheRepository) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return r.client.IncrBy(ctx, key, value).Result()
}

// GetInt gets an integer value from cache
func (r *cacheRepository) GetInt(ctx context.Context, key string) (int64, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(result, 10, 64)
}
