package articles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/internal/repository"
	"github.com/dewzzjr/ais/pkg/collection"
	"github.com/dewzzjr/ais/pkg/errs"
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
	cache[model.Article](c,
		u.CacheArticle,
		fmt.Sprintf(RedisKey, result.ID),
		result,
	)
	return result, nil
}

func (u *usecase) Get(c context.Context, id int64) (*model.Article, error) {
	r, err := u.CacheArticle.Get(c,
		fmt.Sprintf(RedisKey, id),
	)
	if err == nil {
		return &r, nil
	}
	results, err := u.Article.GetArticlesByID(c, id)
	if err != nil {
		return nil, err
	}
	result := collection.First(results)
	if result == nil {
		return nil, errs.Wrap(http.StatusNotFound, errors.New("article not found"))
	}
	cache[model.Article](c,
		u.CacheArticle,
		fmt.Sprintf(RedisKey, result.ID),
		result,
	)
	return result, nil
}

func cache[T any](c context.Context, r repository.Cache[T], key string, result *T) {
	if err := r.Set(c,
		key,
		pointer.Val(result),
	); err != nil {
		log.Println("Cache.Set", key, err)
	}
}
