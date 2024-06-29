package main

import (
	"User/internal/storage"
	"log"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()
}
