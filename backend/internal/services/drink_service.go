package services

import (
	"errors"
	"strings"

	"drakemaye/backend/internal/models"
	"drakemaye/backend/internal/utils"
)

func ValidateDrinkInput(drink models.Drink) error {
	if strings.TrimSpace(drink.Name) == "" {
		return errors.New("drink name is required")
	}
	if strings.TrimSpace(drink.Type) == "" {
		return errors.New("drink type is required")
	}
	if drink.DefaultServingSizeOz <= 0 {
		return errors.New("serving size must be greater than 0")
	}
	if drink.ABVPercent <= 0 || drink.ABVPercent > 100 {
		return errors.New("abvPercent must be between 0 and 100")
	}
	if drink.Calories < 0 {
		return errors.New("calories cannot be negative")
	}
	return nil
}

func EnrichLogStandardDrinks(log *models.DrinkLog) {
	log.StandardDrinks = utils.CalculateStandardDrinks(log.ServingSizeOz, log.ABVPercent)
}

func ValidateLogInput(log models.DrinkLog) error {
	if strings.TrimSpace(log.DrinkName) == "" {
		return errors.New("drinkName is required")
	}
	if strings.TrimSpace(log.DrinkType) == "" {
		return errors.New("drinkType is required")
	}
	if log.ServingSizeOz <= 0 {
		return errors.New("servingSizeOz must be greater than 0")
	}
	if log.ABVPercent <= 0 || log.ABVPercent > 100 {
		return errors.New("abvPercent must be between 0 and 100")
	}
	if log.Calories < 0 {
		return errors.New("calories cannot be negative")
	}
	if log.Rating < 0 || log.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}
	return nil
}
