package app

import (
	"payment-service/internal/config"
	"payment-service/internal/database"
	"payment-service/internal/repos"
	"payment-service/internal/services"
)

type AppContainer struct {
	Cfg     *config.Config
	Storage *database.PostgresStorage

	// repos
	PaymentRepo *repos.PaymentRepo
	WalletRepo  *repos.WalletRepo

	// services
	WalletService  *services.WalletService
	PaymentService *services.PaymentService
}

func NewContainer(cfg *config.Config) *AppContainer {
	storage, _ := database.NewStorage(cfg.DatabaseURL)

	walletRepo := repos.NewWalletRepo(storage)
	paymentRepo := repos.NewPaymentRepo(storage)

	walletService := services.NewWalletService(walletRepo)
	paymentService := services.NewPaymentService(paymentRepo)

	return &AppContainer{
		cfg,
		storage,
		paymentRepo,
		walletRepo,
		walletService,
		paymentService,
	}
}

func (ct *AppContainer) Stop() {
	ct.Storage.Close()
}
