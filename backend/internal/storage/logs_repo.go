package storage

import (
	"database/sql"
	"time"

	"drakemaye/backend/internal/models"
)

func ListLogs(db *sql.DB) ([]models.DrinkLog, error) {
	rows, err := db.Query(`SELECT id, user_id, drink_id, drink_name, drink_type, serving_size_oz, abv_percent, calories, standard_drinks, location, notes, rating, created_at FROM drink_logs ORDER BY datetime(created_at) DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []models.DrinkLog{}
	for rows.Next() {
		var log models.DrinkLog
		var drinkID sql.NullInt64
		var createdAt string
		if err := rows.Scan(&log.ID, &log.UserID, &drinkID, &log.DrinkName, &log.DrinkType, &log.ServingSizeOz, &log.ABVPercent, &log.Calories, &log.StandardDrinks, &log.Location, &log.Notes, &log.Rating, &createdAt); err != nil {
			return nil, err
		}
		if drinkID.Valid {
			log.DrinkID = &drinkID.Int64
		}
		parsed, _ := time.Parse(time.RFC3339Nano, createdAt)
		if parsed.IsZero() {
			parsed, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		}
		log.CreatedAt = parsed
		logs = append(logs, log)
	}
	return logs, rows.Err()
}

func ListLogsSince(db *sql.DB, since time.Time) ([]models.DrinkLog, error) {
	rows, err := db.Query(`SELECT id, user_id, drink_id, drink_name, drink_type, serving_size_oz, abv_percent, calories, standard_drinks, location, notes, rating, created_at FROM drink_logs WHERE datetime(created_at) >= datetime(?) ORDER BY datetime(created_at) DESC`, since.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []models.DrinkLog{}
	for rows.Next() {
		var log models.DrinkLog
		var drinkID sql.NullInt64
		var createdAt string
		if err := rows.Scan(&log.ID, &log.UserID, &drinkID, &log.DrinkName, &log.DrinkType, &log.ServingSizeOz, &log.ABVPercent, &log.Calories, &log.StandardDrinks, &log.Location, &log.Notes, &log.Rating, &createdAt); err != nil {
			return nil, err
		}
		if drinkID.Valid {
			log.DrinkID = &drinkID.Int64
		}
		parsed, _ := time.Parse(time.RFC3339Nano, createdAt)
		if parsed.IsZero() {
			parsed, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		}
		log.CreatedAt = parsed
		logs = append(logs, log)
	}
	return logs, rows.Err()
}

func GetLog(db *sql.DB, id int64) (models.DrinkLog, error) {
	var log models.DrinkLog
	var drinkID sql.NullInt64
	var createdAt string
	err := db.QueryRow(`SELECT id, user_id, drink_id, drink_name, drink_type, serving_size_oz, abv_percent, calories, standard_drinks, location, notes, rating, created_at FROM drink_logs WHERE id = ?`, id).
		Scan(&log.ID, &log.UserID, &drinkID, &log.DrinkName, &log.DrinkType, &log.ServingSizeOz, &log.ABVPercent, &log.Calories, &log.StandardDrinks, &log.Location, &log.Notes, &log.Rating, &createdAt)
	if err != nil {
		return models.DrinkLog{}, err
	}
	if drinkID.Valid {
		log.DrinkID = &drinkID.Int64
	}
	parsed, _ := time.Parse(time.RFC3339Nano, createdAt)
	if parsed.IsZero() {
		parsed, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	}
	log.CreatedAt = parsed
	return log, nil
}

func CreateLogAndFeedPost(db *sql.DB, log models.DrinkLog, userName string) (models.DrinkLog, error) {
	tx, err := db.Begin()
	if err != nil {
		return models.DrinkLog{}, err
	}
	defer tx.Rollback()

	res, err := tx.Exec(`INSERT INTO drink_logs (user_id, drink_id, drink_name, drink_type, serving_size_oz, abv_percent, calories, standard_drinks, location, notes, rating) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		log.UserID, log.DrinkID, log.DrinkName, log.DrinkType, log.ServingSizeOz, log.ABVPercent, log.Calories, log.StandardDrinks, log.Location, log.Notes, log.Rating)
	if err != nil {
		return models.DrinkLog{}, err
	}
	logID, err := res.LastInsertId()
	if err != nil {
		return models.DrinkLog{}, err
	}
	log.ID = logID
	log.CreatedAt = time.Now().UTC()

	_, err = tx.Exec(`INSERT INTO feed_posts (user_id, user_name, drink_log_id, drink_name, drink_type, abv_percent, calories, standard_drinks, location, notes, like_count) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0)`,
		log.UserID, userName, log.ID, log.DrinkName, log.DrinkType, log.ABVPercent, log.Calories, log.StandardDrinks, log.Location, log.Notes)
	if err != nil {
		return models.DrinkLog{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.DrinkLog{}, err
	}
	return log, nil
}
