package services

import "drakemaye/backend/internal/models"

func BuildBadges(user models.User) []models.Badge {
	return []models.Badge{
		{Name: "First Log", Description: "Logged your first drink", Earned: user.LifetimeDrinksLogged >= 1},
		{Name: "Tried 10 Drinks", Description: "Logged 10 unique drinks", Earned: user.DistinctDrinksLoggedLifetime >= 10},
		{Name: "Low Calorie Week", Description: "Stayed under weekly calorie goal", Earned: user.GoalsMetCurrentWeek},
		{Name: "Explorer", Description: "Tried 3+ drink types", Earned: user.DistinctDrinksLoggedLifetime >= 3},
	}
}
