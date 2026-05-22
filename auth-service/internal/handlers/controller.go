package handlers

import (
	"auth-service/internal/app"
	"auth-service/internal/net"
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
	auth := rest.router.Group("/auth")
	{
		auth.POST("/register", rest.CreatePlayerHandler())
		auth.POST("/login", rest.LoginPlayerHandler())
		auth.GET("/me", rest.GetMe())
	}
}
