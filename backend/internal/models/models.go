package models

import "time"

type User struct {
	ID                           int64   `json:"id"`
	Username                     string  `json:"username"`
	FavoriteDrink                string  `json:"favoriteDrink"`
	WeeklyCalorieGoal            int     `json:"weeklyCalorieGoal"`
	WeeklyStandardDrinkGoal      float64 `json:"weeklyStandardDrinkGoal"`
	NewDrinksToTryGoal           int     `json:"newDrinksToTryGoal"`
	LifetimeDrinksLogged         int     `json:"lifetimeDrinksLogged"`
	LifetimeCalories             int     `json:"lifetimeCalories"`
	LifetimeStandardDrinks       float64 `json:"lifetimeStandardDrinks"`
	GoalsMetCurrentWeek          bool    `json:"goalsMetCurrentWeek"`
	DistinctDrinksLoggedLifetime int     `json:"distinctDrinksLoggedLifetime"`
}

type Drink struct {
	ID                   int64   `json:"id"`
	Name                 string  `json:"name"`
	Type                 string  `json:"type"`
	DefaultServingSizeOz float64 `json:"defaultServingSizeOz"`
	ABVPercent           float64 `json:"abvPercent"`
	Calories             int     `json:"calories"`
	Description          string  `json:"description"`
}

type DrinkLog struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"userId"`
	DrinkID        *int64    `json:"drinkId,omitempty"`
	DrinkName      string    `json:"drinkName"`
	DrinkType      string    `json:"drinkType"`
	ServingSizeOz  float64   `json:"servingSizeOz"`
	ABVPercent     float64   `json:"abvPercent"`
	Calories       int       `json:"calories"`
	StandardDrinks float64   `json:"standardDrinks"`
	Location       string    `json:"location"`
	Notes          string    `json:"notes"`
	Rating         int       `json:"rating"`
	CreatedAt      time.Time `json:"createdAt"`
}

type FeedPost struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"userId"`
	UserName       string    `json:"userName"`
	DrinkLogID     int64     `json:"drinkLogId"`
	DrinkName      string    `json:"drinkName"`
	DrinkType      string    `json:"drinkType"`
	ABVPercent     float64   `json:"abvPercent"`
	Calories       int       `json:"calories"`
	StandardDrinks float64   `json:"standardDrinks"`
	Location       string    `json:"location"`
	Notes          string    `json:"notes"`
	Timestamp      time.Time `json:"timestamp"`
	LikeCount      int       `json:"likeCount"`
}

type Goal struct {
	WeeklyCalorieGoal       int     `json:"weeklyCalorieGoal"`
	WeeklyStandardDrinkGoal float64 `json:"weeklyStandardDrinkGoal"`
	NewDrinksToTryGoal      int     `json:"newDrinksToTryGoal"`
}

type Badge struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Earned      bool   `json:"earned"`
}

type StatsSummary struct {
	RangeLabel           string  `json:"rangeLabel"`
	TotalDrinks          int     `json:"totalDrinks"`
	TotalCalories        int     `json:"totalCalories"`
	TotalStandardDrinks  float64 `json:"totalStandardDrinks"`
	AverageABV           float64 `json:"averageAbv"`
	MostCommonDrinkType  string  `json:"mostCommonDrinkType"`
	FavoriteHighestRated string  `json:"favoriteHighestRated"`
}

type ProfileResponse struct {
	User   User    `json:"user"`
	Goals  Goal    `json:"goals"`
	Badges []Badge `json:"badges"`
}
