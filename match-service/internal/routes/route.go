package routes

import (
	"match-service/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(pool *pgxpool.Pool, router *gin.Engine) {
	g := router.Group("/match")
	{
		g.POST("/create", handlers.CreateMatchHandler(pool))
		g.DELETE("/delete", handlers.DeleteMatchHandler(pool))
		g.GET("/list", handlers.GetAllMatchesHandler(pool))
		g.POST("/join", handlers.JoinMatchHandler(pool))
	}
}
