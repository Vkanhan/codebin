package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
	}
	tmp, err := template.ParseFiles(files ...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return 
	}
	err = tmp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println("error executing the template")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil && id < 1 {
		http.Error(w, "not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Display specific snippet with ID: %d", id)
}

func gistCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a new gist"))
}
