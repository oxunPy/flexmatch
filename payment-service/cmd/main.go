package main

import (
	"log"
	"net/http"
	"payment-service/internal/config"
	"payment-service/internal/database"
	"payment-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("err load config", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("err connect database", err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "payment-service")
	})

	routes.Setup(router, pool)

	if err := router.Run(); err != nil {
		log.Fatal()
	}
}
