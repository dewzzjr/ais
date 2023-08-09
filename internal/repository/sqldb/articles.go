package sqldb

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
)

func (r *repo) FetchArticles(c context.Context) (result []model.Article, err error) {
	find := r.db.WithContext(c).Find(&result)
	return result, find.Error
}

func (r *repo) InsertArticle(c context.Context, payload model.Article) error {
	return r.db.WithContext(c).Create(&payload).Error
}
