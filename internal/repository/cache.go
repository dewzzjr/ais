//go:generate mockgen --build_flags=--mod=mod -destination=mocks/cache.go -package=mocks . CacheArticle
package repository

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

type CacheArticle Cache[model.Article]

type Cache[T any] interface {
	Get(context.Context, string) (T, error)
	Del(context.Context, ...string) error
	Set(context.Context, string, T) error
}
