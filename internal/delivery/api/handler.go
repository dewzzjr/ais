package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dewzzjr/ais/internal/delivery/api/payload"
	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/pkg/errs"
	"github.com/dewzzjr/ais/pkg/request"
	"github.com/dewzzjr/ais/pkg/response"
	"github.com/gorilla/mux"
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

func (d *delivery) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		response.Error(w, errs.Wrap(http.StatusBadRequest, errors.New("missing id")))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, errs.Wrap(http.StatusBadRequest, errors.New("id is not integer")))
		return
	}
	result, err := d.Article.Get(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Send(w, http.StatusOK,
		payload.ArticleResponse{}.FromModel(result),
	)
}
