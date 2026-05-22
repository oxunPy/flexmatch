package app

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/repos"
	"auth-service/internal/services"
)

type AppContainer struct {
	Cfg     *config.Config
	Storage *database.PostgresStorage

	// repos
	PlayerRepo *repos.PlayerRepo
	TokenRepo  *repos.TokenRepo

	// services
	AuthService  *services.AuthService
	TokenService *services.TokenService
}

func NewContainer(cfg *config.Config) *AppContainer {
	storage, _ := database.NewStorage(cfg.DatabaseURL)

	playerRepo := repos.NewPlayerRepo(storage)
	tokenRepo := repos.NewTokenRepo(storage)

	authService := services.NewAuthService(playerRepo)
	tokenService := services.NewTokenService(tokenRepo)

	return &AppContainer{
		cfg,
		storage,
		playerRepo,
		tokenRepo,
		authService,
		tokenService,
	}
}

func (ct *AppContainer) Stop() {
	ct.Storage.Close()
}
