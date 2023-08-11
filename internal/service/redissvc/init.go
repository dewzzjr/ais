package redissvc

import (
	"github.com/dewzzjr/ais/internal/config"
	"github.com/redis/go-redis/v9"
)

func New(cfg config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})
}
