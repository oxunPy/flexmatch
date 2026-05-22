package services

import "payment-service/internal/repos"

type WalletService struct {
	walletRepo *repos.WalletRepo
}

func NewWalletService(walletRepo *repos.WalletRepo) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}
