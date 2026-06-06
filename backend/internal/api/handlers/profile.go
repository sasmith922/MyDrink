package handlers

import (
	"database/sql"
	"net/http"

	"drakemaye/backend/internal/models"
	"drakemaye/backend/internal/services"
	"drakemaye/backend/internal/storage"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct{ db *sql.DB }

func NewProfileHandler(db *sql.DB) ProfileHandler { return ProfileHandler{db: db} }

func (h ProfileHandler) Get(c *gin.Context) {
	// TODO: Replace hardcoded user ID with authenticated user context.
	user, err := storage.GetUser(h.db, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	goals := models.Goal{
		WeeklyCalorieGoal:       user.WeeklyCalorieGoal,
		WeeklyStandardDrinkGoal: user.WeeklyStandardDrinkGoal,
		NewDrinksToTryGoal:      user.NewDrinksToTryGoal,
	}
	c.JSON(http.StatusOK, models.ProfileResponse{User: user, Goals: goals, Badges: services.BuildBadges(user)})
}

func (h ProfileHandler) UpdateGoals(c *gin.Context) {
	var goals models.Goal
	if err := c.ShouldBindJSON(&goals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if goals.WeeklyCalorieGoal < 0 || goals.WeeklyStandardDrinkGoal < 0 || goals.NewDrinksToTryGoal < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "goal values must be non-negative"})
		return
	}
	// TODO: Replace hardcoded user ID with authenticated user context.
	if err := storage.UpdateGoals(h.db, 1, goals); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.Get(c)
}
