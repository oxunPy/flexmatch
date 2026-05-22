package handlers

import (
	"payment-service/internal/app"
	"payment-service/internal/net"
)

type RestController struct {
	router *net.GinRouter
	*app.AppContainer
}

func NewRestController(
	router *net.GinRouter,
	container *app.AppContainer,
) *RestController {
	return &RestController{
		router:       router,
		AppContainer: container,
	}
}

func (rest *RestController) Setup() {
	var payment = rest.router.Group("/payment")
	{
		payment.GET("/all", rest.GetAllPaymentsHandler())
		payment.POST("/create", rest.CreatePaymentHandler())
	}

	var wallet = rest.router.Group("/wallet.proto")
	{
		wallet.POST("/create", rest.CreateWalletHandler())
		wallet.GET("/my", rest.GetMyWalletsHandler())
		wallet.GET("/list", rest.GetAllWalletsHandler())
	}
}
