//go:generate mockgen -source=articles.go -destination=mocks/articles.go -package=mocks
package usecase

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

type Article interface {
	Fetch(context.Context, model.Filter) ([]model.Article, error)
	Insert(context.Context, model.Article) (*model.Article, error)
	Get(context.Context, int64) (*model.Article, error)
}
