package grpc

import (
	"context"
	"payment-service/internal/services"
	v1 "payment-service/protos/gen/wallet/protobuf"
)

type WalletServiceApi struct {
	v1.UnimplementedWalletServiceServer
	service *services.WalletService
}

func RegisterWalletServiceApi(s *GrpcServer, service *services.WalletService) {
	v1.RegisterWalletServiceServer(s.server, &WalletServiceApi{
		service: service,
	})
}

func (s *WalletServiceApi) CheckBalance(context.Context, *v1.BalanceRequest) (*v1.BalanceResponse, error) {
	return nil, nil
}
func (s *WalletServiceApi) GetWallets(context.Context, *v1.WalletsRequest) (*v1.WalletsResponse, error) {
	return nil, nil
}
