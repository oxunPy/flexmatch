package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"market-service/internal/models"
	"market-service/internal/repos"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BuyWeaponRequest struct {
	PlayerID int64  `json:"player_id"`
	WeaponID string `json:"weapon_id"`
}

type BuyWeaponResponse struct {
	Success  bool   `json:"success"`
	Info     string `json:"info"`
	WeaponID string `json:"weapon_id"`
}

func BuyWeaponHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req BuyWeaponRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid data request",
			})
			return
		}

		playerWeapon, err := repos.BuyWeapon(pool, req.WeaponID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error buying weapon: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, BuyWeaponResponse{
			Success:  true,
			Info:     "successfully bought new weapon",
			WeaponID: playerWeapon.WeaponID,
		})
	}
}

type SellWeaponRequest struct {
	PlayerID int64  `json:"player_id"`
	WeaponID string `json:"weapon_id"`
}

type SellWeaponResponse struct {
	Success  bool   `json:"success"`
	Info     string `json:"info"`
	WeaponID string `json:"weapon_id"`
}

func SellWeaponHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SellWeaponRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid data request: %s\n", err.Error()),
			})
			return
		}

		success, err := repos.SellWeapon(pool, req.WeaponID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error selling weapon: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, SellWeaponResponse{
			Success:  success,
			Info:     "weapon sold successfully",
			WeaponID: req.WeaponID,
		})
	}
}

type CreateWeaponRequest struct {
	Name string            `json:"name" binding:"required"`
	Desc string            `json:"desc" binding:"required"`
	Type models.WeaponType `json:"type"`
	Cost float64           `json:"cost" binding:"required"`
	Attr json.RawMessage   `json:"attr"`
}

type CreateWeaponResponse struct {
	Success bool          `json:"success"`
	Info    string        `json:"info"`
	Weapon  models.Weapon `json:"weapon"`
}

func CreateWeaponHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateWeaponRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid req data: %s\n", err.Error()),
			})
			return
		}

		var attr models.WeaponAttributes
		if len(req.Attr) > 0 {
			if err := json.Unmarshal(req.Attr, &attr); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid weapon attr: %s\n", err.Error()),
				})
				return
			}
		}

		weapon, err := repos.CreateWeapon(pool, models.Weapon{
			Name: req.Name,
			Desc: req.Desc,
			Type: req.Type,
			Cost: req.Cost,
			Attr: attr,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error storing weapon: %s\n", err.Error()),
			})
			return
		}

		c.JSON(http.StatusCreated, CreateWeaponResponse{
			Success: true,
			Info:    "weapon created successfully",
			Weapon:  *weapon,
		})
	}
}

type GetAllWeaponsResponse struct {
	Success bool             `json:"success"`
	List    []*models.Weapon `json:"list"`
}

func GetAllWeaponsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		list, err := repos.GetAllWeapons(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error get all weapons: %s\n", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, GetAllWeaponsResponse{
			Success: true,
			List:    list,
		})
	}
}
