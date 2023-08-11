package payload

import (
	"time"

	"github.com/dewzzjr/ais/internal/model"
)

type ArticleResponse struct {
	Author    string     `json:"author"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"created_at"`
}

func (ArticleResponse) FromModel(a model.Article) ArticleResponse {
	return ArticleResponse{
		Author:    a.Author,
		Title:     a.Title,
		Body:      a.Body,
		CreatedAt: &a.CreatedAt,
	}
}

func (ArticleResponse) FromModels(as []model.Article) (result []ArticleResponse) {
	result = make([]ArticleResponse, 0)
	for _, v := range as {
		result = append(result, ArticleResponse{}.FromModel(v))
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
