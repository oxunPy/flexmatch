package routes

import (
	"payment-service/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(router *gin.Engine, pool *pgxpool.Pool) {
	var payment = router.Group("/payment")
	{
		payment.GET("/all", handlers.GetAllPaymentsHandler(pool))
		payment.POST("/create", handlers.CreatePaymentHandler(pool))
	}

	var wallet = router.Group("/wallet")
	{
		wallet.POST("/create", handlers.CreateWalletHandler(pool))
		wallet.GET("/my", handlers.GetMyWalletsHandler(pool))
		wallet.GET("/list", handlers.GetAllWalletsHandler(pool))
	}
}
