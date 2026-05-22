package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:\n", err)
	}

	app := app.New(cfg)
	defer app.Stop()

	rest := handlers.NewRestController(app.GetGinRouter(), app.GetContainer())
	rest.Setup()

	app.Run()
}
