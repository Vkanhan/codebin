package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectToDB() (*sql.DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable not set")
	}

	var db *sql.DB
	var err error

	// Retry connecting to the database
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dbURL)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second) 
	}

	return nil, fmt.Errorf("could not connect to database after multiple attempts: %v", err)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	server := &http.Server{
		Handler: routes(),
		Addr:    ":" + portString,
	}

	log.Printf("Listening to port: %v", portString)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}