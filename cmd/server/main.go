package main

import (
	"log"
	"spool/internal/server"
)

func main() {

	server.ApplyMigrations()

	s := server.NewServer()
	log.Println("Server started on :8080")
	s.ListenAndServe()

}
