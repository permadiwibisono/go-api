package main

import (
	"fmt"
	"go-api/middlewares"
	"go-api/utils"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var defaultRoute = "v1"

func Handler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"version":     defaultRoute,
			"message":     "Hello world",
			"status_code": 200,
		}
		utils.WriteJson(w, data)
	})
	return r
}

func main() {
	port := 3000
	r := chi.NewRouter()
	r.Use(middlewares.Logger)
	r.Route(fmt.Sprintf("/%s", defaultRoute), func(v1 chi.Router) {
		v1.Mount("/", Handler())
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, fmt.Sprintf("/%s", defaultRoute), 301)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"error": map[string]interface{}{
				"message":     "Not Found",
				"status_code": 404,
			},
		}
		w.WriteHeader(http.StatusNotFound)
		utils.WriteJson(w, data)
	})
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Listening on port :%d", port)
	srv.ListenAndServe()
}
