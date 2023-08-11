package redissvc

import (
	"log"

	"github.com/dewzzjr/ais/internal/config"
	"github.com/redis/go-redis/v9"
)

func New(cfg config.Redis) *redis.Client {
	log.Printf("connecting redis [%s]...\n", cfg.Address)
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})
}
