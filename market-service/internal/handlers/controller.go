package handlers

import (
	"market-service/internal/app"
	"market-service/internal/net"
)

type RestController struct {
	router *net.GinRouter
	*app.AppContainer
}

func NewRestController(
	router *net.GinRouter,
	container *app.AppContainer,
) *RestController {
	return &RestController{
		router:       router,
		AppContainer: container,
	}
}

func (rest *RestController) Setup() {
	var weapon = rest.router.Group("weapon")
	{
		weapon.POST("/create", rest.CreateWeaponHandler())
		weapon.GET("/list", rest.GetAllWeaponsHandler())
		weapon.POST("/buy", rest.BuyWeaponHandler())
		weapon.POST("/sell", rest.SellWeaponHandler())
	}

	var skin = rest.router.Group("skin")
	{
		skin.POST("/create", rest.CreateSkinHandler())
		skin.GET("/list", rest.GetAllSkinsHandler())
		skin.POST("/buy", rest.BuySkinHandler())
		skin.POST("/sell", rest.SellSkinHandler())
	}

	var armor = rest.router.Group("armor")
	{
		armor.POST("/create", rest.CreateArmorHandler())
		armor.GET("/list", rest.GetAllArmorsHandler())
		armor.POST("/buy", rest.BuyArmorHandler())
		armor.POST("/sell", rest.SellArmorHandler())
	}
}
