package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hoyeah69/bookstore_items-api/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 500 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
