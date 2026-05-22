package main

import (
	"log"
	"payment-service/internal/app"
	"payment-service/internal/config"
	"payment-service/internal/handlers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("err load config", err)
	}

	app := app.New(cfg)
	defer app.Stop()

	rest := handlers.NewRestController(app.GetGinRouter(), app.GetContainer())
	rest.Setup()

	app.Run()
}
