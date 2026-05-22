package services

import "auth-service/internal/repos"

type TokenService struct {
	tokenRepo *repos.TokenRepo
}

func NewTokenService(tokenRepo *repos.TokenRepo) *TokenService {
	return &TokenService{
		tokenRepo: tokenRepo,
	}
}
