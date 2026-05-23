package app

import (
	"market-service/internal/config"
	"market-service/internal/database"
	"market-service/internal/repos"
	"market-service/internal/services"
)

type AppContainer struct {
	Cfg     *config.Config
	Storage *database.PostgresStorage

	// repos
	ArmorRepo  *repos.ArmorRepo
	SkinRepo   *repos.SkinRepo
	WeaponRepo *repos.WeaponRepo

	// services
	ArmorService  *services.ArmorService
	SkinService   *services.SkinService
	WeaponService *services.WeaponService
}

func NewContainer(cfg *config.Config) *AppContainer {
	storage, _ := database.NewStorage(cfg.DatabaseURL)

	armorRepo := repos.NewArmorRepo(storage)
	skinRepo := repos.NewSkinRepo(storage)
	weaponRepo := repos.NewWeaponRepo(storage)

	armorService := services.NewArmorService(armorRepo)
	skinService := services.NewSkinService(skinRepo)
	weaponService := services.NewWeaponService(weaponRepo)

	return &AppContainer{
		cfg,
		storage,
		armorRepo,
		skinRepo,
		weaponRepo,
		armorService,
		skinService,
		weaponService,
	}
}

func (ct *AppContainer) Stop() {
	ct.Storage.Close()
}
