package main

import (
	"mini-project-acp12/config"
	"mini-project-acp12/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
