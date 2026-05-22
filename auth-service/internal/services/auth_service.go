package services

import "auth-service/internal/repos"

type AuthService struct {
	playerRepo *repos.PlayerRepo
}

func NewAuthService(playerRepo *repos.PlayerRepo) *AuthService {
	return &AuthService{
		playerRepo: playerRepo,
	}
}
