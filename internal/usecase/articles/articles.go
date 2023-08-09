package articles

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

func (u *usecase) Fetch(c context.Context) ([]model.Article, error) {
	return u.Article.FetchArticles(c)
}

func (u *usecase) Insert(c context.Context, payload model.Article) error {
	// TODO: cache process
	return u.Article.InsertArticle(c, payload)
}
