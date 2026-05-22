package net

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	port   int
	router *gin.Engine
}

func NewRouter(port int) *GinRouter {
	router := gin.Default()
	_ = router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "auth-service")
	})

	return &GinRouter{
		port:   port,
		router: router,
	}
}

func (r *GinRouter) Run() error {
	return r.router.Run(":" + strconv.Itoa(r.port))
}

func (r *GinRouter) Group(path string) *gin.RouterGroup {
	return r.router.Group(path)
}

func (r *GinRouter) POST(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.router.POST(path, handler)
}

func (r *GinRouter) GET(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.router.GET(path, handler)
}

func (r *GinRouter) PUT(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.router.PUT(path, handler)
}

func (r *GinRouter) DELETE(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.router.DELETE(path, handler)
}

func (r *GinRouter) PATCH(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.router.PATCH(path, handler)
}
