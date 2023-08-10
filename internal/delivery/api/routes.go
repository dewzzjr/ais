package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dewzzjr/ais/pkg/middleware"
	"github.com/gorilla/mux"
)

func (d *delivery) Route() *mux.Router {
	r := mux.NewRouter()
	r.Use()
	r.HandleFunc("/articles", d.FetchArticles).Methods(http.MethodGet)
	r.HandleFunc("/articles", d.CreateArticles).Methods(http.MethodPost)
	d.Server.Handler = r
	return r
}

func (d *delivery) Start() {
	r := d.Route()
	r.Use(
		middleware.Logger,
		mux.CORSMethodMiddleware(r),
	)
	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-stopC
		log.Printf("caught signal %v: terminating\n", s)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer wg.Done()
		defer cancel()
		if err := d.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	log.Println("server starting...")
	if err := d.Server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	wg.Wait()
	log.Println("server stopping...")
}
