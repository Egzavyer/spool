package main

import (
	"fmt"
	"log"
	"net/http"
	"spool/internal/server"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

	server.ApplyMigrations()

	http.HandleFunc("/hello", hello)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
