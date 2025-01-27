package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Codebin")

	router := chi.NewRouter()

	router.Get("/", home)
	router.Get("/gist/view", gistView)
	router.Post("/gist/create", gistCreate)

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Listening to port: %v", portString)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
