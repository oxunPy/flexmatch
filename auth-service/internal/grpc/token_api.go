package grpc

import (
	"auth-service/internal/services"
	v1 "auth-service/protos/gen/token/protobuf"
	"context"
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
