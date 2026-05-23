package services

import "market-service/internal/repos"

type SkinService struct {
	skinRepo *repos.SkinRepo
}

func NewSkinService(skinRepo *repos.SkinRepo) *SkinService {
	return &SkinService{
		skinRepo: skinRepo,
	}
}
