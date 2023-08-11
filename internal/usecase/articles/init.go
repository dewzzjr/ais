package articles

import "github.com/dewzzjr/ais/internal/repository"

type usecase struct {
	repository.Article
	repository.CacheArticle
}

func New(r repository.Article, c repository.CacheArticle) *usecase {
	return &usecase{r, c}
}
