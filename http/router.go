package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func router() http.Handler {
	r := mux.NewRouter()

	return r
}

// ListenAndServe starts a new web http server
func ListenAndServe(addr string) error {

	server := http.Server{
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60*time.Second,
		Addr: addr,
		Handler: router(),
	}

	return server.ListenAndServe()
}