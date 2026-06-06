package storage

import (
	"database/sql"

	"drakemaye/backend/internal/models"
)

func ListDrinks(db *sql.DB) ([]models.Drink, error) {
	rows, err := db.Query(`SELECT id, name, type, default_serving_size_oz, abv_percent, calories, description FROM drinks ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drinks []models.Drink
	for rows.Next() {
		var d models.Drink
		if err := rows.Scan(&d.ID, &d.Name, &d.Type, &d.DefaultServingSizeOz, &d.ABVPercent, &d.Calories, &d.Description); err != nil {
			return nil, err
		}
		drinks = append(drinks, d)
	}
	return drinks, rows.Err()
}

func GetDrink(db *sql.DB, id int64) (models.Drink, error) {
	var d models.Drink
	err := db.QueryRow(`SELECT id, name, type, default_serving_size_oz, abv_percent, calories, description FROM drinks WHERE id = ?`, id).
		Scan(&d.ID, &d.Name, &d.Type, &d.DefaultServingSizeOz, &d.ABVPercent, &d.Calories, &d.Description)
	return d, err
}

func CreateDrink(db *sql.DB, d models.Drink) (models.Drink, error) {
	res, err := db.Exec(`INSERT INTO drinks (name, type, default_serving_size_oz, abv_percent, calories, description) VALUES (?, ?, ?, ?, ?, ?)`,
		d.Name, d.Type, d.DefaultServingSizeOz, d.ABVPercent, d.Calories, d.Description)
	if err != nil {
		return models.Drink{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return models.Drink{}, err
	}
	d.ID = id
	return d, nil
}
