package sqldb

import (
	"context"

	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/pkg/pointer"
	"gorm.io/gorm/clause"
)

func (r *repo) FetchArticles(c context.Context, filter model.Filter) (result []model.Article, err error) {
	find := r.db.WithContext(c).
		Model(&model.Article{}).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "created_at"},
			Desc:   true,
		})
	if pointer.Val(filter.Query) != "" {
		find.Where("MATCH(title, body) AGAINST(? IN NATURAL LANGUAGE MODE)", filter.Query)
	}
	find.Find(&result)
	return result, find.Error
}

func (r *repo) InsertArticle(c context.Context, payload model.Article) error {
	return r.db.WithContext(c).Create(&payload).Error
}
