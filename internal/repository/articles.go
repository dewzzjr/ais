//go:generate mockgen -source=articles.go -destination=mocks/articles.go -package=mocks
package repository

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

type Article interface {
	FetchArticles(context.Context) ([]model.Article, error)
	InsertArticle(context.Context, model.Article) error
}
