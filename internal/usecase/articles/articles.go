package articles

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

func (u *usecase) Fetch(c context.Context, filter model.Filter) ([]model.Article, error) {
	return u.Article.FetchArticles(c, filter)
}

func (u *usecase) Insert(c context.Context, payload model.Article) (*model.Article, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}
	// TODO: cache process
	return u.Article.InsertArticle(c, payload)
}
