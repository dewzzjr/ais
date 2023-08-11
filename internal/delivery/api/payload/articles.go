package payload

import (
	"time"

	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/pkg/pointer"
)

type ArticleResponse struct {
	ID        int64      `json:"id"`
	Author    string     `json:"author"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"created_at"`
}

func (ArticleResponse) FromModel(a *model.Article) ArticleResponse {
	if a == nil {
		return ArticleResponse{}
	}
	return ArticleResponse{
		ID:        int64(a.ID),
		Author:    a.Author,
		Title:     a.Title,
		Body:      a.Body,
		CreatedAt: &a.CreatedAt,
	}
}

func (ArticleResponse) FromModels(as []model.Article) (result []ArticleResponse) {
	result = make([]ArticleResponse, 0)
	for _, v := range as {
		result = append(result, ArticleResponse{}.FromModel(pointer.New(v)))
	}
	return
}

type ArticleRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (a ArticleRequest) ToModel() model.Article {
	return model.Article{
		Author: a.Author,
		Title:  a.Title,
		Body:   a.Body,
	}
}
