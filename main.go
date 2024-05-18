package main

import (
	"fmt"
	"go-nicommerce/config"
	"go-nicommerce/env"
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
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", env.Find("APP_PORT", "8000"))))
}
