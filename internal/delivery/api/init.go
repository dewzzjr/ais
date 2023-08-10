package api

import (
	"net/http"

	"github.com/dewzzjr/ais/internal/config"
	"github.com/dewzzjr/ais/internal/repository/sqldb"
	"github.com/dewzzjr/ais/internal/service/mysqlsvc"
	"github.com/dewzzjr/ais/internal/usecase"
	"github.com/dewzzjr/ais/internal/usecase/articles"
)

type delivery struct {
	usecase.Article
	*http.Server
	Config config.API
}

func New(u usecase.Article, c config.API) *delivery {
	return &delivery{
		Article: u,
		Server: &http.Server{
			Addr: c.Address,
		},
		Config: c,
	}
}

func Run() bool {
	cfg := config.Instance()
	db := mysqlsvc.New(cfg.Database)
	r := sqldb.New(db)
	u := articles.New(r)
	d := New(u, cfg.API)
	d.Start()
	return true
}
