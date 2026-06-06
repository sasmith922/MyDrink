package utils

import "testing"

func TestCalculateStandardDrinks(t *testing.T) {
	got := CalculateStandardDrinks(12, 5)
	want := 3.00
	if got != want {
		t.Fatalf("expected %.2f, got %.2f", want, got)
	}
}
