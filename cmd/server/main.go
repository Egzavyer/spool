package main

import (
	"log"
	"spool/config"
	"spool/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	server.ApplyMigrations()

	dbpool := server.ConnectDB()
	defer dbpool.Close()

	s := server.NewServer(dbpool, cfg.Server.Addr)
	log.Printf("Server started on localhost%v\n", cfg.Server.Addr)
	s.ListenAndServe()
}
