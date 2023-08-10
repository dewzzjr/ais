package api

import (
	"encoding/json"
	"net/http"

	"github.com/dewzzjr/ais/internal/model"
)

func (d *delivery) CreateArticles(w http.ResponseWriter, r *http.Request) {
	if err := d.Article.Insert(r.Context(), model.Article{
		// TODO: construct from json payload
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (d *delivery) FetchArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := d.Article.Fetch(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(map[string]interface{}{
		"data": articles,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
