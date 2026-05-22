package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/security"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPlayerRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password" binding:"required"`
}

func (rest *RestController) CreatePlayerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RegisterPlayerRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username, Email, Password are required",
			})
			return
		}

		if len(request.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Password length should be at least 6",
			})
			return
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})
			return
		}

		player, err := rest.PlayerRepo.CreateNewPlayer(models.Player{
			Username:  request.Username,
			Email:     request.Email,
			Firstname: request.Firstname,
			Lastname:  request.Lastname,
			Password:  string(hashedPass),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Error storing player: %v\n", err),
			})
			return
		}

		c.JSON(http.StatusCreated, player)
	}
}

type LoginPlayerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (rest *RestController) LoginPlayerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request LoginPlayerRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username or Email and Password are required",
			})
			return
		}

		player, err := rest.PlayerRepo.GetPlayer(request.Username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Player not found",
			})
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(request.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Password is not matched",
			})
			return
		}

		exp := time.Now().Add(24 * time.Hour).Unix()
		claims := jwt.MapClaims{
			"user_id":  player.ID,
			"username": player.Username,
			"email":    player.Email,
			"exp":      exp,
		}
		token, err := security.NewToken(&claims, rest.Cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("JWT err: %v\n", err),
			})
			return
		}

		pt, err := rest.TokenRepo.CreatePlayerToken(models.PlayerToken{
			Token:     *token,
			PlayerID:  player.ID,
			Player:    *player,
			ExpiredAt: time.Unix(exp, 0),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Error creating player token: %v\n", err),
			})
			return
		}

		c.JSON(http.StatusOK, pt)
	}
}

func (rest *RestController) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
