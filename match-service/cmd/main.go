package main

import (
	"log"
	"match-service/internal/config"
	"match-service/internal/database"
	"match-service/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config\n", err)
	}

	pool, err := database.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("failed to connect database\n", err)
	}
	defer pool.Close()

	router := gin.Default()
	_ = router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "match-service")
	})

	routes.Setup(pool, router)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal()
	}

}
