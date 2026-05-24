package main

import (
	"file-service/internal/app"
	"file-service/internal/config"
	"file-service/internal/handlers"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("err loading config")
	}

	app := app.New(cfg)
	defer app.Stop()

	rest := handlers.NewRestController(app.GetGinRouter(), app.GetContainer())
	rest.Setup()

	app.Run()

}
