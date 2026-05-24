package app

import (
	"file-service/internal/config"
	"file-service/internal/grpc"
	"file-service/internal/net"
)

type App struct {
	router    *net.GinRouter
	grpc      *grpc.GrpcServer
	container *AppContainer
}

func New(cfg *config.Config) *App {
	ct := NewContainer(cfg)
	router := net.NewRouter(cfg.HttpPort)

	server := grpc.NewServer(cfg.GrpcPort)
	grpc.RegisterFileServiceApi(server, ct.FileService)

	return &App{
		container: ct,
		grpc:      server,
		router:    router,
	}
}

func (a *App) GetContainer() *AppContainer {
	return a.container
}

func (a *App) GetGinRouter() *net.GinRouter {
	return a.router
}

func (a *App) Run() {
	a.grpc.Run()
	a.router.Run()
}

func (a *App) Stop() {
	a.grpc.Stop()
	a.container.Stop()
}
