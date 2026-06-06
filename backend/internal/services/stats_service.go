package services

import "drakemaye/backend/internal/models"

func BuildStatsSummary(rangeLabel string, logs []models.DrinkLog) models.StatsSummary {
	summary := models.StatsSummary{RangeLabel: rangeLabel}
	if len(logs) == 0 {
		return summary
	}

	typeCounts := map[string]int{}
	ratingTotals := map[string]int{}
	ratingCounts := map[string]int{}
	abvTotal := 0.0

	for _, log := range logs {
		summary.TotalDrinks++
		summary.TotalCalories += log.Calories
		summary.TotalStandardDrinks += log.StandardDrinks
		abvTotal += log.ABVPercent
		typeCounts[log.DrinkType]++
		ratingTotals[log.DrinkName] += log.Rating
		ratingCounts[log.DrinkName]++
	}

	bestType := ""
	bestTypeCount := 0
	for drinkType, count := range typeCounts {
		if count > bestTypeCount {
			bestType = drinkType
			bestTypeCount = count
		}
	}

	bestDrink := ""
	bestAvg := -1.0
	for drinkName, total := range ratingTotals {
		avg := float64(total) / float64(ratingCounts[drinkName])
		if avg > bestAvg {
			bestAvg = avg
			bestDrink = drinkName
		}
	}

	summary.AverageABV = abvTotal / float64(len(logs))
	summary.MostCommonDrinkType = bestType
	summary.FavoriteHighestRated = bestDrink
	return summary
}
