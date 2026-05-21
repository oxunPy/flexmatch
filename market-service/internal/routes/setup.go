package routes

import (
	"market-service/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(pool *pgxpool.Pool, router *gin.Engine) {
	var weapon = router.Group("weapon")
	{
		weapon.POST("/create", handlers.CreateWeaponHandler(pool))
		weapon.GET("/list", handlers.GetAllWeaponsHandler(pool))
		weapon.POST("/buy", handlers.BuyWeaponHandler(pool))
		weapon.POST("/sell", handlers.SellWeaponHandler(pool))
	}

	var skin = router.Group("skin")
	{
		skin.POST("/create", handlers.CreateSkinHandler(pool))
		skin.GET("/list", handlers.GetAllSkinsHandler(pool))
		skin.POST("/buy", handlers.BuySkinHandler(pool))
		skin.POST("/sell", handlers.SellSkinHandler(pool))
	}

	var armor = router.Group("armor")
	{
		armor.POST("/create", handlers.CreateArmorHandler(pool))
		armor.GET("/list", handlers.GetAllArmorsHandler(pool))
		armor.POST("/buy", handlers.BuyArmorHandler(pool))
		armor.POST("/sell", handlers.SellArmorHandler(pool))
	}
}
