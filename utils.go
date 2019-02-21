package randomock

import (
	"math"
	"math/rand"
)

// RoundTo Rounds to the nearest `digit` number of decimal places.
func RoundTo(value float64, digits int) float64 {
	factor := math.Pow(10, float64(digits))
	return math.Round(factor*value) / factor
}

// RandBetweenFloat64 generates a random number in the range of [a, b).
func RandBetweenFloat64(a, b float64) float64 {
	return (b-a)*rand.Float64() + a
}
