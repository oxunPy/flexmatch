package services

import "market-service/internal/repos"

type WeaponService struct {
	weaponRepo *repos.WeaponRepo
}

func NewWeaponService(weaponRepo *repos.WeaponRepo) *WeaponService {
	return &WeaponService{
		weaponRepo: weaponRepo,
	}
}
