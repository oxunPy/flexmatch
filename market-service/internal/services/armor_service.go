package services

import "market-service/internal/repos"

type ArmorService struct {
	armorRepo *repos.ArmorRepo
}

func NewArmorService(armorRepo *repos.ArmorRepo) *ArmorService {
	return &ArmorService{
		armorRepo: armorRepo,
	}
}
