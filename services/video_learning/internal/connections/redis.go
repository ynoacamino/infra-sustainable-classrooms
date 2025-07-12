package connections

import (
	"github.com/redis/go-redis/v9"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/config"
)

type ConnectRedisConfig interface {
	GetRedisConfig() *config.RedisConfig
}

func ConnectRedis(cfg ConnectRedisConfig) (*redis.Client, error) {
	redisConfig := cfg.GetRedisConfig()

	// TODO should we add security to this?
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Endpoint,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(redisConfig.Ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
