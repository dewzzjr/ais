package articles

import (
	"context"
	"fmt"
	"log"

	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/pkg/collection"
	"github.com/dewzzjr/ais/pkg/pointer"
)

const RedisKey = "articles:%d"

func (u *usecase) Fetch(c context.Context, filter model.Filter) ([]model.Article, error) {
	return u.Article.FetchArticles(c, filter)
}

func (u *usecase) Insert(c context.Context, payload model.Article) (*model.Article, error) {
	if err := payload.Validate(); err != nil {
		return nil, err
	}
	result, err := u.Article.InsertArticle(c, payload)
	if err != nil {
		return nil, err
	}
	if err := u.CacheArticle.Set(c,
		fmt.Sprintf(RedisKey, result.ID),
		pointer.Val(result),
	); err != nil {
		log.Println("CacheArticle.Set", err)
	}
	return result, nil
}

func (u *usecase) Get(c context.Context, id int64) (*model.Article, error) {
	result, err := u.CacheArticle.Get(c,
		fmt.Sprintf(RedisKey, id),
	)
	if err == nil {
		return &result, nil
	}
	results, err := u.Article.GetArticlesByID(c, id)
	if err != nil {
		return nil, err
	}
	return collection.First(results), nil
}
