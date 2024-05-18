package main

import (
	"go-nicommerce/config"
	"go-nicommerce/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
