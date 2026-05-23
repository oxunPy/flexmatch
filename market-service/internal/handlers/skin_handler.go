package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"market-service/internal/models"

	"github.com/gin-gonic/gin"
)

type BuySkinRequest struct {
	PlayerID int64  `json:"player_id"`
	SkinID   string `json:"skin_id"`
}

type BuySkinResponse struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
	SkinID  string `json:"skin_id"`
}

func (rest *RestController) BuySkinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req BuySkinRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid data request",
			})
			return
		}

		playerSkin, err := rest.SkinRepo.BuySkin(req.SkinID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error buying skin: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, BuySkinResponse{
			Success: true,
			Info:    "successfully bought new skin",
			SkinID:  playerSkin.SkinID,
		})
	}
}

type SellSkinRequest struct {
	PlayerID int64  `json:"player_id"`
	SkinID   string `json:"skin_id"`
}

type SellSkinResponse struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
	SkinID  string `json:"skin_id"`
}

func (rest *RestController) SellSkinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SellSkinRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid data request: %s\n", err.Error()),
			})
			return
		}

		success, err := rest.SkinRepo.SellSkin(req.SkinID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error selling skin: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, SellSkinResponse{
			Success: success,
			Info:    "skin sold successfully",
			SkinID:  req.SkinID,
		})
	}
}

type CreateSkinRequest struct {
	Name string          `json:"name" binding:"required"`
	Cost float64         `json:"cost" binding:"required"`
	Attr json.RawMessage `json:"attr"`
}

type CreateSkinResponse struct {
	Success bool        `json:"success"`
	Info    string      `json:"info"`
	Skin    models.Skin `json:"skin"`
}

func (rest *RestController) CreateSkinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateSkinRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid req data: %s\n", err.Error()),
			})
			return
		}

		var attr models.SkinAttributes
		if len(req.Attr) > 0 {
			if err := json.Unmarshal(req.Attr, &attr); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid skin attr: %s\n", err.Error()),
				})
				return
			}
		}

		skin, err := rest.SkinRepo.CreateSkin(models.Skin{
			Name: req.Name,
			Cost: req.Cost,
			Attr: attr,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error storing skin: %s\n", err.Error()),
			})
			return
		}

		c.JSON(http.StatusCreated, CreateSkinResponse{
			Success: true,
			Info:    "skin created successfully",
			Skin:    *skin,
		})
	}
}

type GetAllSkinsResponse struct {
	Success bool           `json:"success"`
	List    []*models.Skin `json:"list"`
}

func (rest *RestController) GetAllSkinsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		list, err := rest.SkinRepo.GetAllSkins()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error get all skins: %s\n", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, GetAllSkinsResponse{
			Success: true,
			List:    list,
		})
	}
}
