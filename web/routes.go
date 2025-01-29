package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	router := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	router.Get("/", home)
	router.Get("/gist/view", gistView)
	router.Post("/gist/create", gistCreate)

	return router
}
