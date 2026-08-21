package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"spool/gen/job/v1/jobv1connect"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func NewServer() *http.Server {
	jobServiceHandler := &JobServiceHandler{}

	mux := http.NewServeMux()
	path, handler := jobv1connect.NewJobServiceHandler(jobServiceHandler,
		// Validation via Protovalidate
		connect.WithInterceptors(validate.NewInterceptor()))

	mux.Handle(path, handler)
	p := new(http.Protocols)
	p.SetHTTP1(true)

	// Use h2c so we can serve HTTP/2 without TLS.
	p.SetUnencryptedHTTP2(true)

	return &http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: p,
	}
}
