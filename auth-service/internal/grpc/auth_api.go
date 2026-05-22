package grpc

import (
	"auth-service/internal/services"
	v1 "auth-service/protos/gen/sso/protobuf"
	"context"
)

type AuthServiceApi struct {
	v1.UnimplementedAuthServer
	auth *services.AuthService
}

func RegisterAuthServiceApi(s *GrpcServer, auth *services.AuthService) {
	v1.RegisterAuthServer(s.server, &AuthServiceApi{
		auth: auth,
	})
}

func (s *AuthServiceApi) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	return nil, nil
}

func (s *AuthServiceApi) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return nil, nil
}

func (s *AuthServiceApi) GetMe(ctx context.Context, req *v1.GetMeRequest) (*v1.GetMeResponse, error) {
	return nil, nil
}
