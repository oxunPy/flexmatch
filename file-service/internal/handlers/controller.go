package handlers

import (
	"file-service/internal/app"
	"file-service/internal/net"
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
	api := rest.router.Group("/files")
	{
		api.POST("/upload", rest.Upload)
		api.GET("/:id", rest.GetInfo)
		api.GET("/:id/download", rest.Download)
		api.DELETE("/:id", rest.Delete)
	}
}
