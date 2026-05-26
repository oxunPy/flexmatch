package app

import (
	"log"

	"market-service/internal/config"
	"market-service/internal/grpc"
	"market-service/internal/net"
)

type App struct {
	router        *net.GinRouter
	grpcServer    *grpc.GrpcServer
	fileClient    *grpc.FileClient
	paymentClient *grpc.PaymentClient
	container     *AppContainer
}

func New(cfg *config.Config) *App {
	ct := NewContainer(cfg)
	router := net.NewRouter(cfg.HttpPort)

	server := grpc.NewServer(cfg.GrpcPort)
	grpc.RegisterSkinServiceApi(server, ct.SkinService)
	grpc.RegisterArmorServiceApi(server, ct.ArmorService)
	grpc.RegisterWeaponServiceApi(server, ct.WeaponService)

	return &App{
		container:  ct,
		grpcServer: server,
		router:     router,
	}
}

func (a *App) GetContainer() *AppContainer {
	return a.container
}

func (a *App) GetGinRouter() *net.GinRouter {
	return a.router
}

func (a *App) Run() {
	go func() {
		if err := a.grpcServer.Run(); err != nil {
			log.Println("grpc server stopped:", err)
		}
	}()

	if err := a.router.Run(); err != nil {
		log.Println("http server stopped:", err)
	}
}

func (a *App) Stop() {
	a.grpcServer.Stop()
	a.container.Stop()
}
