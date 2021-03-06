package app

import (
	//"log"
	"net/http"
	"time"

	"github.com/Moriartii/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8002",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  50 * time.Second,
	}

	if err := (srv.ListenAndServe()); err != nil {
		//log.Fatal(err)
		panic(err)
	}
}
