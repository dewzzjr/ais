//go:generate mockgen -source=articles.go -destination=mocks/articles.go -package=mocks
package repository

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

type Article interface {
	FetchArticles(context.Context, model.Filter) ([]model.Article, error)
	InsertArticle(context.Context, model.Article) (*model.Article, error)
	GetArticlesByID(context.Context, ...int64) ([]model.Article, error)
}
