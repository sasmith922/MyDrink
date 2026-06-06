package storage

import "database/sql"

func RunMigrations(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT NOT NULL,
			favorite_drink TEXT NOT NULL DEFAULT '',
			weekly_calorie_goal INTEGER NOT NULL DEFAULT 0,
			weekly_standard_drink_goal REAL NOT NULL DEFAULT 0,
			new_drinks_to_try_goal INTEGER NOT NULL DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS drinks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			type TEXT NOT NULL,
			default_serving_size_oz REAL NOT NULL,
			abv_percent REAL NOT NULL,
			calories INTEGER NOT NULL,
			description TEXT NOT NULL DEFAULT ''
		);`,
		`CREATE TABLE IF NOT EXISTS drink_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			drink_id INTEGER,
			drink_name TEXT NOT NULL,
			drink_type TEXT NOT NULL,
			serving_size_oz REAL NOT NULL,
			abv_percent REAL NOT NULL,
			calories INTEGER NOT NULL,
			standard_drinks REAL NOT NULL,
			location TEXT NOT NULL DEFAULT '',
			notes TEXT NOT NULL DEFAULT '',
			rating INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
		`CREATE TABLE IF NOT EXISTS feed_posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			user_name TEXT NOT NULL,
			drink_log_id INTEGER NOT NULL,
			drink_name TEXT NOT NULL,
			drink_type TEXT NOT NULL,
			abv_percent REAL NOT NULL,
			calories INTEGER NOT NULL,
			standard_drinks REAL NOT NULL,
			location TEXT NOT NULL DEFAULT '',
			notes TEXT NOT NULL DEFAULT '',
			timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			like_count INTEGER NOT NULL DEFAULT 0,
			FOREIGN KEY(drink_log_id) REFERENCES drink_logs(id)
		);`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}
