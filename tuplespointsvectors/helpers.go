package tuplespointsvectors

import "math"

func Equals(a, b float64) bool {
	diff := math.Abs(a - b)
	return diff < 0.00001
}
