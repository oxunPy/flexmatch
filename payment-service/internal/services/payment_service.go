package services

import "payment-service/internal/repos"

type PaymentService struct {
	paymentRepo *repos.PaymentRepo
}

func NewPaymentService(paymentRepo *repos.PaymentRepo) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}
