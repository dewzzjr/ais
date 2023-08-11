package api

import (
	"net/http"

	"github.com/dewzzjr/ais/internal/delivery/api/payload"
	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/pkg/errs"
	"github.com/dewzzjr/ais/pkg/request"
	"github.com/dewzzjr/ais/pkg/response"
)

func (d *delivery) CreateArticles(w http.ResponseWriter, r *http.Request) {
	var req payload.ArticleRequest
	if err := request.Read(r, &req); err != nil {
		response.Error(w, errs.Wrap(http.StatusBadRequest, err))
		return
	}
	result, err := d.Article.Insert(r.Context(), req.ToModel())
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Send(w, http.StatusCreated,
		payload.ArticleResponse{}.FromModel(result),
	)
}

func (d *delivery) FetchArticles(w http.ResponseWriter, r *http.Request) {
	var filter model.Filter
	if err := request.Read(r, &filter, request.Query); err != nil {
		response.Error(w, errs.Wrap(http.StatusBadRequest, err))
		return
	}
	articles, err := d.Article.Fetch(r.Context(), filter)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Send(w, http.StatusOK, map[string]interface{}{
		"data": payload.ArticleResponse{}.FromModels(articles),
	})
}
