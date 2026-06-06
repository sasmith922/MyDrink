package storage

import (
	"database/sql"
	"time"

	"drakemaye/backend/internal/models"
)

func ListFeedPosts(db *sql.DB) ([]models.FeedPost, error) {
	rows, err := db.Query(`SELECT id, user_id, user_name, drink_log_id, drink_name, drink_type, abv_percent, calories, standard_drinks, location, notes, timestamp, like_count FROM feed_posts ORDER BY datetime(timestamp) DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.FeedPost{}
	for rows.Next() {
		var p models.FeedPost
		var ts string
		if err := rows.Scan(&p.ID, &p.UserID, &p.UserName, &p.DrinkLogID, &p.DrinkName, &p.DrinkType, &p.ABVPercent, &p.Calories, &p.StandardDrinks, &p.Location, &p.Notes, &ts, &p.LikeCount); err != nil {
			return nil, err
		}
		parsed, _ := time.Parse(time.RFC3339Nano, ts)
		if parsed.IsZero() {
			parsed, _ = time.Parse("2006-01-02 15:04:05", ts)
		}
		p.Timestamp = parsed
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func LikeFeedPost(db *sql.DB, id int64) (models.FeedPost, error) {
	if _, err := db.Exec(`UPDATE feed_posts SET like_count = like_count + 1 WHERE id = ?`, id); err != nil {
		return models.FeedPost{}, err
	}
	return GetFeedPost(db, id)
}

func GetFeedPost(db *sql.DB, id int64) (models.FeedPost, error) {
	var p models.FeedPost
	var ts string
	err := db.QueryRow(`SELECT id, user_id, user_name, drink_log_id, drink_name, drink_type, abv_percent, calories, standard_drinks, location, notes, timestamp, like_count FROM feed_posts WHERE id = ?`, id).
		Scan(&p.ID, &p.UserID, &p.UserName, &p.DrinkLogID, &p.DrinkName, &p.DrinkType, &p.ABVPercent, &p.Calories, &p.StandardDrinks, &p.Location, &p.Notes, &ts, &p.LikeCount)
	if err != nil {
		return models.FeedPost{}, err
	}
	parsed, _ := time.Parse(time.RFC3339Nano, ts)
	if parsed.IsZero() {
		parsed, _ = time.Parse("2006-01-02 15:04:05", ts)
	}
	p.Timestamp = parsed
	return p, nil
}
