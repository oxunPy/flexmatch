package handlers

import (
	"encoding/json"
	"fmt"
	"market-service/internal/models"
	"market-service/internal/repos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BuyArmorRequest struct {
	PlayerID int64  `json:"player_id"`
	ArmorID  string `json:"armor_id"`
}

type BuyArmorResponse struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
	ArmorID string `json:"armor_id"`
}

func BuyArmorHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req BuyArmorRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data request",
			})
			return
		}

		pa, err := repos.BuyArmor(pool, req.ArmorID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error buying armor: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, BuyArmorResponse{
			Success: true,
			Info:    "Successfully bought new armor",
			ArmorID: pa.ArmorID,
		})
	}
}

type SellArmorRequest struct {
	PlayerID int64  `json:"player_id"`
	ArmorID  string `json:"armor_id"`
}

type SellArmorResponse struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
	ArmorID string `json:"armor_id"`
}

func SellArmorHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SellArmorRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid data request: %s\n", err.Error()),
			})
			return
		}

		succ, err := repos.SellArmor(pool, req.ArmorID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error selling armor: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, SellArmorResponse{
			Success: succ,
			Info:    "armor selled successfully",
			ArmorID: req.ArmorID,
		})
	}
}

type CreateArmorRequest struct {
	Name string          `json:"name" binding:"required"`
	Desc string          `json:"desc" binding:"required"`
	Cost float64         `json:"cost" binding:"required"`
	Attr json.RawMessage `json:"attr"`
}

type CreateArmorResponse struct {
	Success bool         `json:"success"`
	Info    string       `json:"info"`
	Armor   models.Armor `json:"armor"`
}

func CreateArmorHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateArmorRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid req data: %s\n", err.Error()),
			})
			return
		}

		var attr models.ArmorAttributes
		if len(req.Attr) > 0 {
			if err := json.Unmarshal(req.Attr, &attr); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("invalid armor attr: %s\n", err.Error()),
				})
				return
			}
		}

		armor, err := repos.CreateArmor(pool, models.Armor{
			Name: req.Name,
			Desc: req.Desc,
			Cost: req.Cost,
			Attr: attr,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error storing armor: %s\n", err.Error()),
			})
			return
		}

		c.JSON(http.StatusCreated, CreateArmorResponse{
			Success: true,
			Info:    "Armor created successfully!",
			Armor:   *armor,
		})

	}
}

type GetAllArmorsResponse struct {
	Success bool            `json:"success"`
	List    []*models.Armor `json:"list"`
}

func GetAllArmorsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		list, err := repos.GetAllArmors(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error get all armors: %s\n", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, GetAllArmorsResponse{
			Success: true,
			List:    list,
		})
	}
}
