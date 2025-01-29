package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
		"./ui/partials/navigation.html",
	}
	tmp, err := template.ParseFiles(files...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", err))
		return
	}
	err = tmp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", err))
		return
	}
}

func gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("Not found: %v", err))
		return
	}
	fmt.Fprintf(w, "Display specific snippet with ID: %d", id)
}

func gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a new gist"))
}
