package server

import (
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
)

func ApplyMigrations() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Applying migrations")
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

	log.Println("Migrations applied successfully")
}
