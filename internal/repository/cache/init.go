package cache

import (
	"github.com/redis/go-redis/v9"
)

type client[T any] struct {
	*redis.Client
}

func New[T any](cli *redis.Client) *client[T] {
	return &client[T]{cli}
}
