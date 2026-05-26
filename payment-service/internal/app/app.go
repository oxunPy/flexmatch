package app

import (
	"log"

	"payment-service/internal/config"
	"payment-service/internal/grpc"
	"payment-service/internal/net"
)

type App struct {
	router    *net.GinRouter
	grpc      *grpc.GrpcServer
	container *AppContainer
}

func New(cfg *config.Config) *App {
	ct := NewContainer(cfg)
	router := net.NewRouter(cfg.Port)

	server := grpc.NewServer(cfg.GPort)
	grpc.RegisterPaymentServiceApi(server, ct.PaymentService)
	grpc.RegisterWalletServiceApi(server, ct.WalletService)

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
	go func() {
		if err := a.grpc.Run(); err != nil {
			log.Println("grpc server stopped:", err)
		}
	}()

	if err := a.router.Run(); err != nil {
		log.Println("http server stopped:", err)
	}
}

func (a *App) Stop() {
	a.grpc.Stop()
	a.container.Stop()
}
