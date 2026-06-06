package utils

import "math"

func CalculateStandardDrinks(servingSizeOz, abvPercent float64) float64 {
	// Approximate formula requested for MVP:
	// standardDrinks = (servingSizeOz * abvPercent * 0.6) / 12
	result := (servingSizeOz * abvPercent * 0.6) / 12
	return math.Round(result*100) / 100
}
