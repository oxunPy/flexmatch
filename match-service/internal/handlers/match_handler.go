package handlers

import (
	"errors"
	"fmt"
	"match-service/internal/models"
	"match-service/internal/repos"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateMatchRequest struct {
	Title string           `json:"title"`
	Date  string           `json:"date"`
	Type  models.MatchType `json:"type"`
}

type CreateMatchResponse struct {
	Success bool         `json:"success"`
	Data    models.Match `json:"data"`
}

func CreateMatchHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateMatchRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data type",
			})
			return
		}

		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid time format",
			})
			return
		}

		match, err := repos.CreateMatch(pool, req.Title, date, req.Type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Error saving match: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusCreated, CreateMatchResponse{
			Success: true,
			Data:    *match,
		})
	}
}

type DeleteMatchRequest struct {
	MatchID string
}

type DeleteMatchResponse struct {
	Success bool   `json:"success"`
	Text    string `json:"text"`
	MatchID string `json:"match_id"`
}

func DeleteMatchHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		matchID := c.Query("match_id")
		if matchID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "match_id param is required",
			})
			return
		}

		success, err := repos.DeleteMatch(pool, matchID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf("match with id: %s not found", matchID),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error while deleting match: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, DeleteMatchResponse{
			Success: success,
			Text:    "Match deleted successfully",
			MatchID: matchID,
		})
	}
}

type ListMatchResponse struct {
	Success bool            `json:"success"`
	List    []*models.Match `json:"list"`
}

func GetAllMatchesHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		matches, err := repos.GetAllMatches(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("err getting all matches: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, ListMatchResponse{
			Success: true,
			List:    matches,
		})
	}
}

type JoinMatchRequest struct {
	MatchID  string `json:"match_id"`
	PlayerID int64  `json:"player_id"`
}

type JoinMatchResponse struct {
	Success  bool   `json:"success"`
	MatchID  string `json:"match_id"`
	PlayerID int64  `json:"player_id"`
}

func JoinMatchHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req JoinMatchRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "PlayerID and MatchID required",
			})
			return
		}
		if req.MatchID == "" || req.PlayerID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data types",
			})

			return
		}

		mp, err := repos.JoinMatch(pool, req.MatchID, req.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Match or Player not found: %s", err.Error()),
			})
		}

		c.JSON(http.StatusCreated, JoinMatchResponse{
			Success:  true,
			MatchID:  mp.MatchID,
			PlayerID: mp.PlayerID,
		})
	}
}
