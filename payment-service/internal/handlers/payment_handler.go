package handlers

import (
	"fmt"
	"net/http"
	"payment-service/internal/models"
	"payment-service/internal/repos"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreatePaymentRequest struct {
	Amount   float64            `json:"amount"`
	Type     models.PaymentType `json:"type"`
	PlayerID int64              `json:"player_id"`
	WalletID int64              `json:"wallet_id"`
	ItemID   string             `json:"item_id"`
}

type CreatePaymentResponse struct {
	Success bool           `json:"success"`
	Info    string         `json:"info"`
	Payment models.Payment `json:"payment"`
}

func CreatePaymentHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePaymentRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data type",
			})
			return
		}

		pay, err := repos.CreatePayment(pool, models.Payment{
			ItemID:   req.ItemID,
			PlayerID: req.PlayerID,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Errorf("error create payment"),
			})
			return
		}

		c.JSON(http.StatusCreated, CreatePaymentResponse{
			Success: true,
			Info:    "payment created successfully",
			Payment: *pay,
		})
	}
}

type GetAllPaymentsResponse struct {
	Success bool              `json:"success"`
	List    []*models.Payment `json:"list"`
}

func GetAllPaymentsHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		pays, err := repos.GetAllPaymentList(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("err get all payments: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, GetAllPaymentsResponse{
			Success: true,
			List:    pays,
		})
	}
}
