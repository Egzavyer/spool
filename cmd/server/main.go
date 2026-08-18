package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

	log.Println("Applying migrations")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	pg_url := os.Getenv("PG_URL")
	m, err := migrate.New(
		"file://db/migrations",
		pg_url)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("Database is already up to date, no migrations applied")
		} else {
			log.Fatalf("Failed to apply migrations: %s", err)
		}
	}

	http.HandleFunc("/hello", hello)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
