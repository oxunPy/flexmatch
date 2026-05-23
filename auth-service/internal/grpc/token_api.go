package grpc

import (
	"auth-service/internal/services"
	"context"
	v1 "protos-service/protos/gen/token/protobuf"
)

type TokenServiceApi struct {
	v1.UnimplementedTokenServer
	token *services.TokenService
}

func RegisterTokenServiceApi(s *GrpcServer, token *services.TokenService) {
	v1.RegisterTokenServer(s.server, &TokenServiceApi{
		token: token,
	})
}

func (s *TokenServiceApi) CheckToken(ctx context.Context, req *v1.TokenRequest) (*v1.TokenResponse, error) {
	return nil, nil
}
