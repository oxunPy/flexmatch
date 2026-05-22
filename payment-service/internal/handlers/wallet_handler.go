package handlers

import (
	"fmt"
	"net/http"
	"payment-service/internal/models"
	"payment-service/internal/repos"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateWalletResponse struct {
	Success bool          `json:"success"`
	Info    string        `json:"info"`
	Wallet  models.Wallet `json:"wallet"`
}

func CreateWalletHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		playerID, err := strconv.ParseInt(c.Query("player_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "player_id param is required",
			})
			return
		}
		wlt, err := repos.CreateWallet(pool, playerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("err creating new wallet: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, CreateWalletResponse{
			Success: true,
			Info:    "Created new wallet successfully",
			Wallet:  *wlt,
		})
	}
}

type GetMyWalletsResponse struct {
	Success  bool             `json:"success"`
	PlayerID int64            `json:"player_id"`
	Wallets  []*models.Wallet `json:"wallets"`
}

func GetMyWalletsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		playerID, err := strconv.ParseInt(c.Query("player_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "player_id param is required",
			})
			return
		}

		wallets, err := repos.GetMyWallets(pool, playerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("err get my wallets: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, GetMyWalletsResponse{
			Success:  true,
			PlayerID: playerID,
			Wallets:  wallets,
		})
	}
}

type GetAllWalletsResponse struct {
	Success       bool                       `json:"success"`
	Info          string                     `json:"info"`
	PlayerWallets map[int64][]*models.Wallet `json:"player_wallets"`
}

func GetAllWalletsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		wallets, err := repos.GetWalletsList(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("err get my wallets: %s", err.Error()),
			})
			return
		}

		pWallets := make(map[int64][]*models.Wallet)
		for i := range wallets {
			wlt := wallets[i]
			pWallets[wlt.PlayerID] = append(pWallets[wlt.PlayerID], wlt)
		}

		c.JSON(http.StatusOK, GetAllWalletsResponse{
			Success:       true,
			Info:          "All player with wallets",
			PlayerWallets: pWallets,
		})
	}
}
