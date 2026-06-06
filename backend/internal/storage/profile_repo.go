package storage

import (
	"database/sql"
	"fmt"
	"time"

	"drakemaye/backend/internal/models"
)

func GetUser(db *sql.DB, userID int64) (models.User, error) {
	var u models.User
	err := db.QueryRow(`SELECT id, username, favorite_drink, weekly_calorie_goal, weekly_standard_drink_goal, new_drinks_to_try_goal FROM users WHERE id = ?`, userID).
		Scan(&u.ID, &u.Username, &u.FavoriteDrink, &u.WeeklyCalorieGoal, &u.WeeklyStandardDrinkGoal, &u.NewDrinksToTryGoal)
	if err != nil {
		return models.User{}, err
	}

	if err := db.QueryRow(`SELECT COUNT(*), COALESCE(SUM(calories), 0), COALESCE(SUM(standard_drinks), 0), COUNT(DISTINCT drink_name) FROM drink_logs WHERE user_id = ?`, userID).
		Scan(&u.LifetimeDrinksLogged, &u.LifetimeCalories, &u.LifetimeStandardDrinks, &u.DistinctDrinksLoggedLifetime); err != nil {
		return models.User{}, err
	}

	startOfWeek := time.Now().AddDate(0, 0, -7).Format("2006-01-02 15:04:05")
	var weekCalories int
	if err := db.QueryRow(`SELECT COALESCE(SUM(calories), 0) FROM drink_logs WHERE user_id = ? AND datetime(created_at) >= datetime(?)`, userID, startOfWeek).Scan(&weekCalories); err != nil {
		return models.User{}, err
	}
	u.GoalsMetCurrentWeek = u.WeeklyCalorieGoal > 0 && weekCalories <= u.WeeklyCalorieGoal
	return u, nil
}

func UpdateGoals(db *sql.DB, userID int64, goal models.Goal) error {
	_, err := db.Exec(`UPDATE users SET weekly_calorie_goal = ?, weekly_standard_drink_goal = ?, new_drinks_to_try_goal = ? WHERE id = ?`,
		goal.WeeklyCalorieGoal, goal.WeeklyStandardDrinkGoal, goal.NewDrinksToTryGoal, userID)
	return err
}

func GetUserName(db *sql.DB, userID int64) (string, error) {
	var username string
	if err := db.QueryRow(`SELECT username FROM users WHERE id = ?`, userID).Scan(&username); err != nil {
		return "", err
	}
	return username, nil
}

func SeedData(db *sql.DB) error {
	var userCount int
	if err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&userCount); err != nil {
		return err
	}
	if userCount > 0 {
		return nil
	}

	if _, err := db.Exec(`INSERT INTO users (id, username, favorite_drink, weekly_calorie_goal, weekly_standard_drink_goal, new_drinks_to_try_goal) VALUES (1, 'DrakeFan92', 'Hazy IPA', 2200, 10, 2)`); err != nil {
		return err
	}

	drinkSeeds := []models.Drink{
		{Name: "Hazy IPA", Type: "Beer", DefaultServingSizeOz: 12, ABVPercent: 6.5, Calories: 210, Description: "Juicy hop-forward IPA."},
		{Name: "Cabernet Sauvignon", Type: "Wine", DefaultServingSizeOz: 5, ABVPercent: 13.5, Calories: 125, Description: "Full-bodied red wine."},
		{Name: "Whiskey Soda", Type: "Cocktail", DefaultServingSizeOz: 8, ABVPercent: 10, Calories: 140, Description: "Whiskey and club soda."},
		{Name: "Light Lager", Type: "Beer", DefaultServingSizeOz: 12, ABVPercent: 4.2, Calories: 102, Description: "Crisp low-calorie lager."},
	}
	for _, d := range drinkSeeds {
		if _, err := CreateDrink(db, d); err != nil {
			return fmt.Errorf("seed drink %s: %w", d.Name, err)
		}
	}

	if _, err := db.Exec(`INSERT INTO drink_logs (user_id, drink_name, drink_type, serving_size_oz, abv_percent, calories, standard_drinks, location, notes, rating, created_at)
		VALUES
		(1, 'Hazy IPA', 'Beer', 12, 6.5, 210, 3.90, 'Boston', 'Game day pour', 5, datetime('now', '-2 day')),
		(1, 'Whiskey Soda', 'Cocktail', 8, 10, 140, 4.00, 'Cambridge', 'After dinner', 4, datetime('now', '-1 day'))`); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO feed_posts (user_id, user_name, drink_log_id, drink_name, drink_type, abv_percent, calories, standard_drinks, location, notes, timestamp, like_count)
		VALUES
		(1, 'DrakeFan92', 1, 'Hazy IPA', 'Beer', 6.5, 210, 3.90, 'Boston', 'Game day pour', datetime('now', '-2 day'), 4),
		(1, 'DrakeFan92', 2, 'Whiskey Soda', 'Cocktail', 10, 140, 4.00, 'Cambridge', 'After dinner', datetime('now', '-1 day'), 2)`); err != nil {
		return err
	}

	return nil
}
