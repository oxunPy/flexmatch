package main

import (
	"log"
	"market-service/internal/config"
	"market-service/internal/database"
	"market-service/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "market-service")
	})

	routes.Setup(pool, router)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal()
	}
}
