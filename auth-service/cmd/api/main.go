package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:\n", err)
	}
	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect database:\n", err)
	}

	defer pool.Close()

	router := gin.Default()
	_ = router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "auth-service")
	})

	routes.Setup(router, pool, cfg)

	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal()
	}
}
