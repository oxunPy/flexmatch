package routes

import (
	"market-service/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(pool *pgxpool.Pool, router *gin.Engine) {
	var weapon = router.Group("weapon")
	{
		weapon.POST("/create", handlers.CreateWeaponHandler())
		weapon.GET("/list", handlers.GetAllWeaponsHandler())
		weapon.POST("/buy", handlers.BuyWeaponHandler())
		weapon.POST("/sell", handlers.SellWeaponHandler())
	}

	var skin = router.Group("skin")
	{
		skin.POST("/create", handlers.CreateSkinHandler())
		skin.GET("/list", handlers.GetAllSkinsHandler())
		skin.POST("/buy", handlers.BuySkinHandler())
		skin.POST("/sell", handlers.SellSkinHandler())
	}

	var armor = router.Group("armor")
	{
		armor.POST("/create", handlers.CreateArmorHandler())
		armor.GET("/list", handlers.GetAllArmorsHandler())
		armor.POST("/buy", handlers.BuyArmorHandler())
		armor.POST("/sell", handlers.SellArmorHandler())
	}
}
