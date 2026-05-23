package main

import (
	"log"
	"market-service/internal/app"
	"market-service/internal/config"
	"market-service/internal/handlers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := app.New(cfg)
	defer app.Stop()

	rest := handlers.NewRestController(app.GetGinRouter(), app.GetContainer())
	rest.Setup()

	app.Run()
}
