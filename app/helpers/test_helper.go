package helpers

import "math"

func HasError(e error) bool {
	return e != nil
}

func CompareWithTreshold(a, b, t float64) bool {
	return math.Abs(a-b) <= t
}
