package cache

import (
	"github.com/dewzzjr/ais/internal/config"
	"github.com/redis/go-redis/v9"
)

type client[T any] struct {
	*redis.Client
	Config config.Redis
}

func New[T any](cli *redis.Client, cfg config.Redis) *client[T] {
	return &client[T]{cli, cfg}
}
