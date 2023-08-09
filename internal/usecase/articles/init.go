package articles

import "github.com/dewzzjr/ais/internal/repository"

type usecase struct {
	repository.Article
}

func New(r repository.Article) *usecase {
	return &usecase{r}
}
