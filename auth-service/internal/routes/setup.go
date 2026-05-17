package routes

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(router *gin.Engine, pool *pgxpool.Pool, cfg *config.Config) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.CreatePlayerHandler(pool))
		auth.POST("/login", handlers.LoginPlayerHandler(pool, cfg))
		auth.GET("/me", handlers.GetMe())
	}
}
