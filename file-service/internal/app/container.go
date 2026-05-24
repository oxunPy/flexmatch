package app

import (
	"file-service/internal/config"
	"file-service/internal/database"
	"file-service/internal/repos"
	"file-service/internal/services"
	"file-service/internal/storage"
)

type AppContainer struct {
	Cfg         *config.Config
	DbStorage   *database.PostgresStorage
	FileStorage *storage.LocalStorage
	// repos
	FileRepo *repos.FileRepo
	// services
	FileService *services.FileService
}

func NewContainer(cfg *config.Config) *AppContainer {
	db, _ := database.NewStorage(cfg.DatabaseURL)
	local, _ := storage.NewLocalStorage(cfg.StoragePath, cfg.BaseURL)

	fileRepo := repos.NewFileRepo(db)

	fileService := services.NewFileService(local, fileRepo)

	return &AppContainer{
		cfg,
		db,
		local,
		fileRepo,
		fileService,
	}
}

func (ct *AppContainer) Stop() {
	ct.DbStorage.Close()
}
