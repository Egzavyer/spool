package main

import (
	"log"
	"spool/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server.ApplyMigrations()

	dbpool := server.ConnectDB()
	defer dbpool.Close()

	s := server.NewServer(dbpool)
	log.Println("Server started on :8080")
	s.ListenAndServe()
}
