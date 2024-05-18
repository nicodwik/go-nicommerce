package main

import (
	"go-nicommerce/config"
	"go-nicommerce/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
