package grpc

import (
	"context"
	"payment-service/internal/services"
	v1 "protos-service/protos/gen/payment/protobuf"
)

type PaymentServiceApi struct {
	v1.UnimplementedPaymentServiceServer
	service *services.PaymentService
}

func RegisterPaymentServiceApi(s *GrpcServer, service *services.PaymentService) {
	v1.RegisterPaymentServiceServer(s.server, &PaymentServiceApi{
		service: service,
	})
}

func (s *PaymentServiceApi) CreatePayment(ctx context.Context, req *v1.CreatePaymentRequest) (*v1.CreatePaymentResponse, error) {
	return nil, nil
}

func (s *PaymentServiceApi) GetPayments(ctx context.Context, req *v1.GetPaymentsRequest) (*v1.GetPaymentsResponse, error) {
	return nil, nil
}
